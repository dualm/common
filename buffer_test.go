package common

import (
	"bytes"
	"reflect"
	"testing"
)

var bs = []byte{0x01, 0x00}

func TestNew(t *testing.T) {

	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want *Buffer
	}{
		{
			name: "1",
			args: args{
				data: bs,
			},
			want: &Buffer{
				buffer: bytes.NewBuffer(bs),
				err:    nil,
			},
		},
		{
			name: "2",
			args: args{
				data: nil,
			},
			want: &Buffer{
				buffer: bytes.NewBuffer(nil),
				err:    nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBuffer(tt.args.data); !reflect.DeepEqual(got.Bytes(), tt.want.Bytes()) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkNewFromPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b := NewBuffer(nil)
		b.WriteLittle(uint8(0))
		b.Put()
	}
}

func BenchmarkNewEmptyFromPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b := NewEmptyBuffer()
		b.WriteLittle(uint8(0))
		b.Put()
	}
}

func TestBuffer_Len(t *testing.T) {
	tests := []struct {
		name string
		b    Buffer
		want int
	}{
		{
			name: "1",
			b: Buffer{
				buffer: bytes.NewBuffer([]byte{0x00, 0x01}),
				err:    nil,
			},
			want: 2,
		},
		{
			name: "2",
			b: Buffer{
				buffer: &bytes.Buffer{},
				err:    nil,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.Len(); got != tt.want {
				t.Errorf("Buffer.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuffer_Bytes(t *testing.T) {
	tests := []struct {
		name string
		b    Buffer
		want []byte
	}{
		{
			name: "1",
			b: Buffer{
				buffer: bytes.NewBuffer([]byte{0x00, 0x01}),
				err:    nil,
			},
			want: []byte{0x00, 0x01},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.Bytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Buffer.Bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuffer_Error(t *testing.T) {
	tests := []struct {
		name    string
		b       *Buffer
		wantErr bool
	}{
		{
			name: "1",
			b: &Buffer{
				buffer: &bytes.Buffer{},
				err:    nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.b.Error(); (err != nil) != tt.wantErr {
				t.Errorf("Buffer.Error() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBuffer_Reset(t *testing.T) {
	tests := []struct {
		name string
		b    *Buffer
	}{
		{
			name: "1",
			b: &Buffer{
				buffer: bytes.NewBuffer([]byte{0x00, 0x01}),
				err:    nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.b.Reset()
			if len(tt.b.buffer.Bytes()) != 0 {
				t.Errorf("Buffer.buffer = %v, want = %v", tt.b.buffer, []byte{})
			}
		})
	}
}
