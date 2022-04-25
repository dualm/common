package common

import (
	"strings"

	"github.com/spf13/viper"
)

func GetString(conf *viper.Viper, keys ...string) string {
	return conf.GetString(makeKey(keys))
}

func GetInt(conf *viper.Viper, keys ...string) int {
	return conf.GetInt(makeKey(keys))
}

func GetStringSlice(conf *viper.Viper, keys ...string) []string {
	return conf.GetStringSlice(makeKey(keys))
}

func GetBool(conf *viper.Viper, keys ...string) bool {
	return conf.GetBool(makeKey(keys))
}

func makeKey(keys []string) string {
	k := make([]string, 0, len(keys))

	for i := range keys {
		k = append(k, keys[i])
	}

	return strings.Join(k, ".")
}
