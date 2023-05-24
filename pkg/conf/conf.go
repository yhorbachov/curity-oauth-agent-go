package conf

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port              int
	EndpointsPefix    string
	TrustedWebOrigins []string
}

func NewDefaultConfig() Config {
	v := viper.New()
	v.AutomaticEnv()
	v.SetEnvPrefix("COA")

	v.SetDefault("PORT", 8080)
	v.SetDefault("ENDPOINTS_PREFIX", "/oauth-agent")

	return Config{
		Port:              v.GetInt("PORT"),
		EndpointsPefix:    v.GetString("ENDPOINTS_PREFIX"),
		TrustedWebOrigins: v.GetStringSlice("TRUSTED_WEB_ORIGINS"),
	}
}
