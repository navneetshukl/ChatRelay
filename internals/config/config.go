package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

func LoadConfig(configPath string) (conf *Config, err error) {
	env := os.Getenv("ENV")
	if env==""{
		env=EnvLocalhost.ToString()
	}
	log.Println("Environment is ",env)
	viper.AutomaticEnv()

	envConfigFileName:=fmt.Sprintf("config.%s",env)
	conf=&Config{}

	viper.SetConfigName(envConfigFileName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)

	err=viper.ReadInConfig()
	if err!=nil{
		return nil,fmt.Errorf("failed to read config file: %w",err)
	}

	err=viper.Unmarshal(&conf)
	if err!=nil{
		return nil,fmt.Errorf("unable to decode into config struct, %v",err)
	}

	if conf.ServerConfig.Environment==""{
		return nil,fmt.Errorf("unable to find environment in config file")
	}

	if conf.ServerConfig.Environment==EnvLocalhost{
		envData,err:=loadEnv(conf.ServerConfig.Environment)
		if err!=nil{
			log.Panic("Env file is not loaded ",err)
		}

		conf.BotConfig.SlackAppToken=envData.SlackAppToken
		conf.BotConfig.SlackBotToken=envData.SlackBotToken
		conf.MockServerConfig.BaseURL=envData.MockServerUrl
		conf.MockServerConfig.Port=envData.MockServerPort
		return conf,nil
		
	}

	// load config for other environment also
	return nil,err
}