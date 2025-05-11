package config

type ServerConfig struct {
	Environment Env `json:"environment"`
}

type BotConfig struct {
	SlackAppToken string `json:"slack_app_token"`
	SlackBotToken string `json:"slack_bot_token"`
}

type MockServerConfig struct {
	BaseURL string `json:"base_url"`
	Port    string `json:"port"`
}

type Config struct {
	ServerConfig     ServerConfig     `json:"server_config"`
	BotConfig        BotConfig        `json:"bot_config"`
	MockServerConfig MockServerConfig `json:"mock_server_config"`
}

type Env string

const (
	EnvLocalhost Env = "localhost"
	EnvDev       Env = "dev"
	EnvProd      Env = "prod"
)

func (e Env) ToString() string {
	return string(e)
}
