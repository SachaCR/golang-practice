package config

import (
	"fmt"
	"golang-practice/pkg/env"

	"github.com/spf13/viper"
)

func LoadConfiguration(environment env.AppEnvironment) {

	// Server Default Configuration
	viper.SetDefault("server.port", "8080")
	viper.SetDefault("server.host", "localhost")

	viper.BindEnv("server.port", "API_PORT")
	viper.BindEnv("server.host", "API_HOST")

	// Database Default Configuration
	viper.SetDefault("postgres.host", "localhost")
	viper.SetDefault("postgres.port", "5432")
	viper.SetDefault("postgres.password", "local")
	viper.SetDefault("postgres.user", "local")
	viper.SetDefault("postgres.database", "testDB")
	viper.BindEnv("postres.password", "PG_PASSWORD")
	viper.BindEnv("postres.password", "PG_USER")

	viper.SetConfigName(environment.String())
	viper.SetConfigType("json")

	viper.AddConfigPath("./config")

	var err = viper.ReadInConfig()

	if err != nil {
		_, ok := err.(viper.ConfigFileNotFoundError)

		if ok {
			var errorMessage = "CONFIG FILE NOT FOUND"

			if environment == env.Production {
				panic(fmt.Errorf(errorMessage))
			}

			fmt.Println(errorMessage + ": " + "Default values loaded")
		} else {
			panic(fmt.Errorf("CONFIG FILE ERROR: %w", err))
		}
	}

}
