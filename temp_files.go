package common

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ugorji/go/codec"
)

// MakePath, 创建文件夹路径，如果不存在就创建一个，如果存在一个就复用，如果存在多个就全部清理并生成新的
//	dir: 文件夹上级路径，如果为空则使用os.TempDir()
//	pattern: 最后一个“*”会使用一串随机字符替代
func MakePath(dir, pattern string) (string, error) {
	if len(dir) == 0 {
		dir = os.TempDir()
	}

	matches, err := filepath.Glob(filepath.Join(dir, pattern))
	if err != nil {
		return "", err
	}

	if len(matches) == 1 {
		return matches[0], nil
	}

	for _, match := range matches {
		if err := os.RemoveAll(match); err != nil {
			return "", err
		}
	}

	return os.MkdirTemp(dir, pattern)
}

// MakeFile: 生成临时文件、缓存文件、日志文件等，返回该文件供读写。在生成文件前，会扫描dir，如果存在一个符合pattern的文件文件就复用
// 否则清除路径下的所有符合pattern的文件清除。如果没有任何符合patter的文件，则创建一个新的
//	dir: 文件路径
//	pattern: 文件名样式，最后一个"*"会使用随机字符串进行替代
// 	readOnly: 是否是只读模式
func GetFile(dir, pattern string, readOnly bool) (*os.File, error) {
	cacheGlobPattern := filepath.Join(dir, pattern)

	matches, err := filepath.Glob(cacheGlobPattern)
	if err != nil {
		return nil, fmt.Errorf("getFile error, pattern: %s, %w", cacheGlobPattern, err)
	}

	if len(matches) == 1 {
		if readOnly {
			return os.OpenFile(matches[0], os.O_RDONLY, 0644)
		}

		return os.OpenFile(matches[0], os.O_RDONLY|os.O_TRUNC, 0644)
	}

	for _, match := range matches {
		if err := os.RemoveAll(match); err != nil {
			return nil, fmt.Errorf("remove files error when GetFile, %w", err)
		}
	}

	return os.CreateTemp(dir, pattern)
}

func DecodeFile(v interface{}, f *os.File, h codec.Handle) (interface{}, error) {
	dec := codec.NewDecoder(f, h)

	if err := dec.Decode(v); err != nil {
		return nil, err
	}

	return v, nil
}

func EncodeFile(f *os.File, v interface{}, encode codec.Handle) error {
	enc := codec.NewEncoder(f, encode)

	if err := enc.Encode(v); err != nil {
		return err
	}

	if err := f.Sync(); err != nil {
		return err
	}

	return nil
}
