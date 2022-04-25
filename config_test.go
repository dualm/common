package common

import (
	"fmt"
	"strings"
	"testing"

	"github.com/spf13/viper"
)

var conf *viper.Viper

func TestGetString(t *testing.T) {
	var err error
	conf, err = InitConfig("config")
	if err != nil {
		panic(err)
	}
	type args struct {
		conf  *viper.Viper
		key   string
		nodes []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				conf:  conf,
				key:   "DbName",
				nodes: []string{"DB"},
			},
			want: "svdb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetString(tt.args.conf, tt.args.key, tt.args.nodes...); got != tt.want {
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
		}, {
			name: "2",
			args: args{
				keys: []string{"a", ""},
			},
			want: "a",
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
