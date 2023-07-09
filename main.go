package main

import (
	"fmt"
	"golang-practice/pkg/config"
	"golang-practice/pkg/env"
	inmemory "golang-practice/pkg/repositories/inMemory"
	"golang-practice/pkg/server"
	"golang-practice/pkg/todo"
	"os"

	"github.com/spf13/viper"
)

func main() {

	var environment env.AppEnvironment = env.ParseEnvFromString(os.Getenv("GO_ENV"))

	config.LoadConfiguration(environment)

	var serverAddress = viper.GetString("server.host") + ":" + viper.GetString("server.port")

	var todoRepository = inmemory.New()

	todoRepository.Save(todo.New(todo.TodoToCreate{
		Id:          "1",
		Title:       "My Todo Title",
		Description: "My Todo Description",
	}))

	var dependencies = server.ServerDependencies{TodoRepository: todoRepository}
	var router = server.New(environment, dependencies)

	fmt.Println("Server listen on: " + serverAddress)
	router.Run(serverAddress)
}
