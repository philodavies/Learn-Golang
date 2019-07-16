package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Error when reading config: %v\n", err))
	}

	port := viper.GetInt("port")
	hostname := viper.GetString("hostname")
	auth := viper.GetStringMapString("auth")
	colo := viper.GetInt("colo")

	fmt.Printf("Reading config for port = %d\n", port)
	fmt.Printf("Reading config for hostname = %s\n", hostname)
	fmt.Printf("Reading config for auth = %v\n", auth)

	for k, v := range auth {
		fmt.Printf("k/v pair in auth! key: %v value: %v\n", k, v)
	}
}
