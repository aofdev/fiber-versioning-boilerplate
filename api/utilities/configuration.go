package utilities

import (
	"fmt"

	"github.com/spf13/viper"
)

func ViperEnvVariable(key string) string {

	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Error while reading config file: %s", err))
	}

	value, ok := viper.Get(key).(string)
	if !ok {
		panic("Invalid type assertion")
	}

	return value
}
