package common

import (
	"strings"

	"github.com/spf13/viper"
)

func GetString(conf *viper.Viper, key string, nodes ...string) string {
	nodes = append(nodes, key)
	return conf.GetString(makeKey(nodes))
}

func GetInt(conf *viper.Viper, key string, nodes ...string) int {
	nodes = append(nodes, key)
	return conf.GetInt(makeKey(nodes))
}

func GetStringSlice(conf *viper.Viper, key string, nodes ...string) []string {
	nodes = append(nodes, key)
	return conf.GetStringSlice(makeKey(nodes))
}

func GetBool(conf *viper.Viper, key string, nodes ...string) bool {
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
