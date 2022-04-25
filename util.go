package common

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/dualm/ethernet-ip/bufferEip"
	"go.uber.org/zap"
)

func BytesToInt32(raw []byte, count int) ([]string, error) {
	buffer := bufferEip.New(raw)

	var n int32
	result := make([]string, count)

	for i := 0; i < count; i++ {
		buffer.ReadLittle(&n)
		result[i] = strconv.FormatInt(int64(n), 10)
	}

	if err := buffer.Error(); err != nil {
		return nil, fmt.Errorf("error in encoding bytes to int32, Error: %w", err)
	}

	return result, nil
}

func BytesToInt16(raw []byte, count int) ([]string, error) {
	buffer := bufferEip.New(raw)

	var n int32
	result := make([]string, count)

	for i := 0; i < count; i++ {
		buffer.ReadLittle(&n)
		result[i] = strconv.FormatInt(int64(int16(n)), 10)
	}

	if err := buffer.Error(); err != nil {
		return nil, fmt.Errorf("error in encoding bytes to int16, Error: %w", err)
	}

	return result, nil
}

func BytesToFloat32(raw []byte, count int) ([]string, error) {
	buffer := bufferEip.New(raw)

	var n float32
	result := make([]string, count)

	for i := 0; i < count; i++ {
		buffer.ReadLittle(&n)

		result[i] = strconv.FormatFloat(float64(n), 'f', 4, 32)
	}

	if err := buffer.Error(); err != nil {
		return nil, fmt.Errorf("error in encoding bytes to float32, Error: %w", err)
	}

	return result, nil
}

func BytesToAscii(raw []byte, count int, charCount int) ([]string, error) {
	buffer := bufferEip.New(raw)

	s := make([]byte, charCount)
	result := make([]string, count)

	for i := 0; i < count; i++ {
		buffer.ReadLittle(&s)

		result[i] = TrimByteToString(s)
	}

	if err := buffer.Error(); err != nil {
		return nil, fmt.Errorf("error in encoding bytes to ascii, Error: %w", err)
	}

	return result, nil
}

func TrimByteToString(b []byte) string {
	return string(bytes.TrimFunc(b, func(r rune) bool {
		return !unicode.IsNumber(r) && !unicode.IsLetter(r)
	}))
}

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

func Trim(s string) string {
	return strings.TrimFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r)
	})
}

func TrimMap(s map[string]string) map[string]string {
	re := make(map[string]string, len(s))
	for k, v := range s {
		re[k] = Trim(v)
	}

	return re
}

func PrintlnError(err error) {
	if err != nil {
		zap.S().Error(err)
	}
}
