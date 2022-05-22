package common

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func GetString(conf *viper.Viper, node, key string) string {
	nodes := makeNodes(node)

	nodes = append(nodes, key)

	return conf.GetString(makeKey(nodes))
}

func GetSliceNodeString(conf *viper.Viper, nodes []string, key string) string {
	nodes = append(nodes, key)

	return conf.GetString(makeKey(nodes))
}

func GetInt(conf *viper.Viper, node, key string) int {
	nodes := makeNodes(node)

	nodes = append(nodes, key)

	return conf.GetInt(makeKey(nodes))
}

func GetSliceNodeInt(conf *viper.Viper, nodes []string, key string) int {
	nodes = append(nodes, key)

	return conf.GetInt(makeKey(nodes))
}

func GetStringSlice(conf *viper.Viper, node, key string) []string {
	nodes := makeNodes(node)

	nodes = append(nodes, key)

	return conf.GetStringSlice(makeKey(nodes))
}

func GetSliceNodeStringSlice(conf *viper.Viper, nodes []string, key string) []string {
	nodes = append(nodes, key)

	return conf.GetStringSlice(makeKey(nodes))
}

func GetBool(conf *viper.Viper, node, key string) bool {
	nodes := makeNodes(node)

	nodes = append(nodes, key)

	return conf.GetBool(makeKey(nodes))
}

func GetSliceNodeBool(conf *viper.Viper, nodes []string, key string) bool {
	nodes = append(nodes, key)

	return conf.GetBool(makeKey(nodes))
}

func makeKey(keys []string) string {
	k := make([]string, 0, len(keys))

	for i := range keys {
		if len(keys[i]) != 0 {
			k = append(k, keys[i])
		}
	}

	return strings.Join(k, ".")
}

func makeNodes(node string) []string {
	nodes := make([]string, 0)

	if _n := strings.Split(node, "."); len(_n) >= 1 {
		nodes = append(nodes, _n...)
	} else {
		nodes = append(nodes, node)
	}

	return nodes
}

func InitConfig(configId, configPath string) (*viper.Viper, error) {
	conf := viper.New()

	conf.SetConfigType("toml")
	conf.SetConfigName(strings.ToLower(configId))
	conf.AddConfigPath(configPath)
	err := conf.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("initialize config error, ConfigId: %s, Error:%w", configId, err)
	}

	return conf, nil
}
