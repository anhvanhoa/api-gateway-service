package bootstrap

import (
	"strings"

	"github.com/anhvanhoa/service-core/bootstrap/config"
)

type ServiceConfig struct {
	Name   string `mapstructure:"name"`
	Route  string `mapstructure:"route"`
	Host   string `mapstructure:"host"`
	Port   string `mapstructure:"port"`
	Folder string `mapstructure:"folder"`
}

type Env struct {
	NodeEnv  string          `mapstructure:"node_env"`
	Port     int             `mapstructure:"port"`
	Services []ServiceConfig `mapstructure:"services"`
}

func NewEnv(env any) {
	setting := config.DefaultSettingsConfig()
	if setting.IsProduction() {
		// setting.SetPath("/config")
		setting.SetFile("api_gateway.config")
	} else {
		setting.SetFile("dev.config")
	}
	config.NewConfig(setting, env)
}

func (env *Env) IsProduction() bool {
	return strings.ToLower(env.NodeEnv) == "production"
}
