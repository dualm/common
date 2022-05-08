package common

import (
	"fmt"
	"strings"
	"testing"

	"github.com/spf13/viper"
)

var conf *viper.Viper

func TestGetString(t *testing.T) {
	e7conf := func() *viper.Viper {
		conf, err := InitConfig("e7")
		if err != nil {
			panic(err)
		}

		return conf
	}()
	type args struct {
		conf  *viper.Viper
		nodes string
		key   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				conf: func() *viper.Viper {
					conf, err := InitConfig("config")
					if err != nil {
						panic(err)
					}

					return conf
				}(),
				key:   "DbName",
				nodes: "DB",
			},
			want: "svdb",
		},
		{
			name: "2",
			args: args{
				conf:  e7conf,
				key:   "Name",
				nodes: "StatusCompo",
			},
			want: e7conf.GetString("StatusCompo.Name"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetString(tt.args.conf, tt.args.nodes, tt.args.key); got != tt.want {
				t.Errorf("GetString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func InitConfig(configId string) (*viper.Viper, error) {
	conf := viper.New()
	conf.SetConfigType("toml")
	conf.SetConfigName(strings.ToLower(configId))
	conf.AddConfigPath("./test")
	err := conf.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("initialize config error, ConfigId: %s, Error:%w", configId, err)
	}

	return conf, nil
}

func Test_makeKey(t *testing.T) {
	type args struct {
		keys []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				keys: []string{"a", "b", ""},
			},
			want: "a.b",
		},
		{
			name: "2",
			args: args{
				keys: []string{"a", ""},
			},
			want: "a",
		},
		{
			name: "3",
			args: args{
				keys: []string{"Name"},
			},
			want: "Name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeKey(tt.args.keys); got != tt.want {
				t.Errorf("makeKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
