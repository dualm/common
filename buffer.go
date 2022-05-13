package common

import (
	"bytes"
	"encoding/binary"
	"io"

	"sync"
)

type Buffer struct {
	buffer *bytes.Buffer
	err    error
}

func (b *Buffer) WriteLittle(target interface{}) {
	b.err = binary.Write(b.buffer, binary.LittleEndian, target)
}

func (b *Buffer) WriteBig(target interface{}) {
	b.err = binary.Write(b.buffer, binary.LittleEndian, target)
}

func (b *Buffer) ReadBig(target interface{}) {
	b.err = binary.Read(b.buffer, binary.BigEndian, target)
}

func (b *Buffer) ReadLittle(target interface{}) {
	b.err = binary.Read(b.buffer, binary.LittleEndian, target)
}

func (b *Buffer) Reset() {
	b.buffer.Reset()
}

func (b *Buffer) Error() error {
	if b.err == io.EOF {
		return nil
	}

	return b.err
}

func (b Buffer) Bytes() []byte {
	return b.buffer.Bytes()
}

func (b Buffer) Len() int {
	return b.buffer.Len()
}

func (b *Buffer) Put() {
	eipPool.Put(b)
}

// NewBuffer，从buffer池中取出并重置一个Buffer实例，将data写入后返回. data可以为nil.
func NewBuffer(data []byte) *Buffer {
	b := NewEmptyBuffer()

	b.buffer.Write(data)

	return b
}

func NewEmptyBuffer() *Buffer {
	b := eipPool.Get().(*Buffer)

	b.Reset()

	return b
}

var eipPool = sync.Pool{
	New: func() interface{} {
		return &Buffer{
			buffer: bytes.NewBuffer(nil),
			err:    nil,
		}
	},
}
