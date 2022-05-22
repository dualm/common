package common

import (
	"testing"

	"github.com/spf13/viper"
)

var v, _ = InitConfig("config", "./test")

func TestGetString(t *testing.T) {
	e7conf, err := InitConfig("e7", "./test")
	if err != nil {
		t.Error(err)
	}

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
				conf:  v,
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

func TestGetSliceNodeString(t *testing.T) {
	e7conf, err := InitConfig("e7", "./test")
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		conf  *viper.Viper
		nodes []string
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
				conf:  e7conf,
				nodes: []string{"StatusCompo"},
				key:   "Name",
			},
			want: e7conf.GetString("StatusCompo.Name"),
		},
		{
			name: "2",
			args: args{
				conf:  e7conf,
				nodes: []string{},
				key:   "Port",
			},
			want: e7conf.GetString("Port"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSliceNodeString(tt.args.conf, tt.args.nodes, tt.args.key); got != tt.want {
				t.Errorf("GetSliceNodeString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSliceNodeInt(t *testing.T) {
	e7conf, err := InitConfig("e7", "./test")
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		conf  *viper.Viper
		nodes []string
		key   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				conf:  e7conf,
				nodes: []string{""},
				key:   "Node",
			},
			want: e7conf.GetInt("Node"),
		},
		{
			name: "2",
			args: args{
				conf:  e7conf,
				nodes: []string{"StatusCompo"},
				key:   "Count",
			},
			want: e7conf.GetInt("StatusCompo.Count"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSliceNodeInt(tt.args.conf, tt.args.nodes, tt.args.key); got != tt.want {
				t.Errorf("GetSliceNodeInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetInt(t *testing.T) {
	e7conf, err := InitConfig("e7", "./test")
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		conf *viper.Viper
		node string
		key  string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				conf: e7conf,
				node: "",
				key:  "Port",
			},
			want: e7conf.GetInt("Port"),
		},
		{
			name: "2",
			args: args{
				conf: e7conf,
				node: "StatusCompo",
				key:  "Count",
			},
			want: e7conf.GetInt("StatusCompo.Count"),
		},
		{
			name: "3",
			args: args{
				conf: e7conf,
				node: "StatusCompo.Count",
				key:  "",
			},
			want: e7conf.GetInt("StatusCompo.Count"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetInt(tt.args.conf, tt.args.node, tt.args.key); got != tt.want {
				t.Errorf("GetInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetBool(t *testing.T) {
	e7conf, err := InitConfig("e7", "./test")
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		conf *viper.Viper
		node string
		key  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				conf: e7conf,
				node: "RecipeValidationEnable",
				key:  "",
			},
			want: e7conf.GetBool("RecipeValidationEnable"),
		},
		{
			name: "2",
			args: args{
				conf: e7conf,
				node: "",
				key:  "RecipeValidationEnable",
			},
			want: e7conf.GetBool("RecipeValidationEnable"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetBool(tt.args.conf, tt.args.node, tt.args.key); got != tt.want {
				t.Errorf("GetBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSliceNodeBool(t *testing.T) {
	type args struct {
		conf  *viper.Viper
		nodes []string
		key   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				conf:  v,
				nodes: []string{"LOG", "Production"},
				key:   "",
			},
			want: v.GetBool("LOG.Production"),
		}, {
			name: "2",
			args: args{
				conf:  v,
				nodes: []string{},
				key:   "LOG.Production",
			},
			want: v.GetBool("LOG.Production"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSliceNodeBool(tt.args.conf, tt.args.nodes, tt.args.key); got != tt.want {
				t.Errorf("GetSliceNodeBool() = %v, want %v", got, tt.want)
			}
		})
	}
}
