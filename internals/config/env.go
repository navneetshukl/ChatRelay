package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type EnvStruct struct {
	SlackAppToken  string `mapstructure:"SLACK_APP_TOKEN"`
	SlackBotToken  string `mapstructure:"SLACK_BOT_TOKEN"`
	MockServerUrl  string `mapstructure:"MOCK_SERVER_BASE_URL"`
	MockServerPort string `mapstructure:"MOCK_SERVER_PORT"`
}

func loadEnv(env Env) (*EnvStruct, error) {
	envData := &EnvStruct{}

	log.Println("Comming ENv is ", env)

	viper.AutomaticEnv()
	//envConfigFileName := fmt.Sprintf(".env.%s", env)

	// log.Println("EnvConfigFileName is ",envConfigFileName)
	// viper.SetConfigFile(envConfigFileName)
	// viper.SetConfigType("env")
	// viper.AddConfigPath("./.secrets")

	envConfigFileName := fmt.Sprintf("./.secrets/.env.%s", env)
	log.Println("EnvConfigFileName is ", envConfigFileName)

	viper.SetConfigFile(envConfigFileName)
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("Config file not found.Using environment variables")
		} else {
			return nil, fmt.Errorf("failed to read config file: %w", err)
		}
	}
	err = viper.Unmarshal(envData)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}
	return envData, nil

}
