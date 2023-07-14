package main

import (
	"fmt"
	"golang-practice/pkg/config"
	"golang-practice/pkg/env"
	"golang-practice/pkg/server"
	"os"

	"github.com/spf13/viper"
)

func main() {

	var environment env.AppEnvironment = env.ParseEnvFromString(os.Getenv("GO_ENV"))

	config.LoadConfiguration(environment)

	var serverAddress = viper.GetString("server.host") + ":" + viper.GetString("server.port")

	var router = server.New(environment)

	fmt.Println("Server listen on: " + serverAddress)
	router.Run(serverAddress)
}
