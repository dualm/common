package common

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func BytesToInt32(raw []byte) ([]string, error) {
	buf := NewBuffer(raw)

	int32s := make([]int32, 0)
	buf.ReadLittle(int32s)

	if err := buf.Error(); err != nil {
		return nil, fmt.Errorf("error in encoding bytes to int32, Error: %w", err)
	}

	s := make([]string, len(int32s))
	for i := range int32s {
		s[i] = strconv.Itoa(int(int32s[i]))
	}

	return s, nil
}

func BytesToInt16(raw []byte) ([]string, error) {
	buf := NewBuffer(raw)

	ints := make([]int16, 0)
	buf.ReadLittle(ints)

	s := make([]string, len(ints))
	for i := range ints {
		s[i] = strconv.FormatInt(int64(ints[i]), 10)
	}

	return s, nil
}

func BytesToFloat32(raw []byte, count int) ([]string, error) {
	buffer := NewBuffer(raw)

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
