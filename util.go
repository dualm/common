package common

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// BytesToInt32, 从raw中提取count个32位整数，并转换为字符串切片.
func BytesToInt32(raw []byte, count int) ([]string, error) {
	_buf := NewBuffer(raw)

	var n int32

	s := make([]string, count)
	for i := 0; i < count; i++ {
		_buf.ReadLittle(&n)

		if err := _buf.Error(); err != nil {
			return nil, fmt.Errorf("error in encoding bytes to int32, Error: %w", err)
		}

		s[i] = strconv.Itoa(int(n))
	}

	return s, nil
}

// BytesToInt16, 从raw中提取count个16位整数，并转换为字符串切片.
func BytesToInt16(raw []byte, count int) ([]string, error) {
	_buf := NewBuffer(raw)

	var _n int16

	s := make([]string, count)

	for i := 0; i < count; i++ {
		_buf.ReadLittle(&_n)

		if err := _buf.Error(); err != nil {
			return nil, fmt.Errorf("error in encoding bytes to int16, Error: %w", err)
		}

		s[i] = strconv.Itoa(int(_n))
	}

	return s, nil
}

// BytesToFloat32, 从raw中提取count个32位浮点数，精度为prec，并转换为字符串切片.
func BytesToFloat32(raw []byte, count, prec int) ([]string, error) {
	buffer := NewBuffer(raw)

	var n float32
	result := make([]string, count)

	for i := 0; i < count; i++ {
		buffer.ReadLittle(&n)

		result[i] = strconv.FormatFloat(float64(n), 'f', prec, 32)
	}

	if err := buffer.Error(); err != nil {
		return nil, fmt.Errorf("error in encoding bytes to float32, Error: %w", err)
	}

	return result, nil
}

// BytesToAscii，从raw中提取包含charCount个字符的count个字符串.
func BytesToAscii(raw []byte, count int, charCount int) ([]string, error) {
	buf := NewBuffer(raw)

	s := make([]byte, charCount)
	result := make([]string, count)

	for i := 0; i < count; i++ {
		buf.ReadLittle(&s)

		result[i] = TrimByteToString(s)
	}

	if err := buf.Error(); err != nil {
		return nil, fmt.Errorf("error in encoding bytes to ascii, Error: %w", err)
	}

	return result, nil
}

// TrimByteToString，移除b中非数据和非字母的内容，并返回字符串
func TrimByteToString(b []byte) string {
	_b := make([]byte, 0, len(b))

	for i := range b {
		if unicode.IsLetter(rune(b[i])) || unicode.IsNumber(rune(b[i])) {
			_b = append(_b, b[i])
		}
	}

	return string(_b)
}

// StringToByteByLength，返回长度为length的字符串s的字节切片
func StringToByteByLength(s string, length int) []byte {
	buf := bytes.NewBufferString(s)
	for buf.Len() < length {
		buf.Write([]byte{0x20})
	}
	if length%2 == 1 {
		buf.Write([]byte{0x20})
	}

	return buf.Bytes()
}

// Trim, 移除s前端、后端非字母和数字的内容.
func Trim(s string) string {
	return strings.TrimFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r)
	})
}

// TrimMap，移除s中所有value里的非字母和数字的部分
func TrimMap(s map[string]string) map[string]string {
	re := make(map[string]string, len(s))
	for k, v := range s {
		re[k] = Trim(v)
	}

	return re
}

// TrimLittleEndianUint16ToString, 小端编码的word字符串解码
func TrimLittleEndianUint16ToString(b []byte) string {
	if len(b)&0x01 == 1 {
		b = append(b, 0x00)
	}

	_b := make([]byte, 0, len(b))

	for i := 1; i < len(b); i++ {
		_b = append(_b, b[i], b[i-1])
		i += 1
	}

	return TrimByteToString(_b)
}

// TrimLittleEndianUint16ToString, 大端编码的word字符串解码
func TrimBigEndianUint16ToString(b []byte) string {
	return TrimByteToString(b)
}
