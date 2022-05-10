package common

import (
	"reflect"
	"testing"
)

func TestTrimByteToString(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				b: []byte{0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x00, 0x00, 0x00, 0x00, 0x52, 0x30, 0x32, 0x39, 0x20},
			},
			want: "R029",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrimByteToString(tt.args.b); got != tt.want {
				t.Errorf("TrimByteToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytesToAscii(t *testing.T) {
	type args struct {
		raw       []byte
		count     int
		charCount int
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				raw: []byte{
					0x52,
					0x65,
					0x63,
					0x69,
					0x70,
					0x65,
					0x53,
					0x65,
					0x74,
					0x31,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x4d,
					0x6f,
					0x64,
					0x65,
					0x6c,
					0x5f,
					0x31,
					0x37,
					0x75,
					0x6d,
					0x54,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
					0x0,
				},
				count:     10,
				charCount: 40,
			},
			want:    []string{"RecipeSet1", "Model_17umT", "", "", "", "", "", "", "", ""},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BytesToAscii(tt.args.raw, tt.args.count, tt.args.charCount)
			if (err != nil) != tt.wantErr {
				t.Errorf("BytesToAscii() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BytesToAscii() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytesToInt32(t *testing.T) {
	type args struct {
		raw   []byte
		count int
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				raw:   []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
				count: 2,
			},
			want:    []string{"0", "0"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BytesToInt32(tt.args.raw, tt.args.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("BytesToInt32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BytesToInt32() = %v, want %v", got, tt.want)
			} else {
				t.Logf("result: %v", got)
			}
		})
	}
}

var byteToInt32 = []byte{0x00, 0x01, 0x02, 0x03, 0x00, 0x01, 0x02, 0x03}

func BenchmarkByteToInt322(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BytesToInt32(byteToInt32, 1)
	}
}

func TestBytesToInt16(t *testing.T) {
	type args struct {
		raw   []byte
		count int
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				raw:   []byte{0x00, 0x01},
				count: 1,
			},
			want:    []string{"256"},
			wantErr: false,
		},
		{
			name: "2",
			args: args{
				raw:   []byte{0x00, 0x01, 0x01, 0x00},
				count: 2,
			},
			want:    []string{"256", "1"},
			wantErr: false,
		},
		{
			name: "3",
			args: args{
				raw:   []byte{0x00, 0x01, 0x01},
				count: 2,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "4",
			args: args{
				raw:   []byte{0x00, 0x01, 0x01, 0x00},
				count: 1,
			},
			want:    []string{"256"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BytesToInt16(tt.args.raw, tt.args.count)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BytesToInt16() = %v, want %v", got, tt.want)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("BytesToInt16() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func TestBytesToFloat32(t *testing.T) {
	type args struct {
		raw   []byte
		count int
		prec  int
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				raw:   []byte{0x14, 0xAE, 0xD9, 0x41},
				count: 1,
				prec:  3,
			},
			want:    []string{"27.210"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BytesToFloat32(tt.args.raw, tt.args.count, tt.args.prec)
			if (err != nil) != tt.wantErr {
				t.Errorf("BytesToFloat32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BytesToFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToByteByLength(t *testing.T) {
	type args struct {
		s      string
		length int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "1",
			args: args{
				s:      "Hello",
				length: 20,
			},
			want: []byte("Hello               "),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToByteByLength(tt.args.s, tt.args.length); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringToByteByLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrim(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				s: " Hello, World=.",
			},
			want: "Hello, World",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Trim(tt.args.s); got != tt.want {
				t.Errorf("Trim() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrimMap(t *testing.T) {
	type args struct {
		s map[string]string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
	{
		name: "1",
		args: args{
			s: map[string]string{
				"1": "Hello World.",
			},
		},
		want: map[string]string{
			"1": "Hello World",
		},
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrimMap(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TrimMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
