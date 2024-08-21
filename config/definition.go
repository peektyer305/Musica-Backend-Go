package config

import (
	"bytes"
	_ "embed"
	"os"
	"strings"

	"github.com/spf13/viper"
)

//go:embed config.yaml
var configYaml []byte

var Conf *config

type config struct {
	App struct {
		Name string `mapstructure:"name"`
		Env string `mapstructure:"env"`
		URL string `mapstructure:"url"`
		IsMaintenanceMode bool `mapstructure:"is_maintenance_mode"`
		SkipMaintenanceModeToken string `mapstructure:"skip_maintenance_mode_token"`
		AllowOrigins []string `mapstructure:"allow_origins"`
		Port string `mapstructure:"port"`
	} `mapstructure:"app"`
	Db struct {
		User string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host string `mapstructure:"host"`
		Port string `mapstructure:"port"`
		Name string `mapstructure:"name"`
	} `mapstructure:"db"`
	}

	func (c config) IsLocal() bool {
		return c.App.Env == "local"
	}

	func (c config) IsDevelopment() bool {
		return c.App.Env == "development"
	}

	func (c config) IsProduction() bool {
		return c.App.Env == "production"
	}

	func (c config) IsTest() bool {
		return c.App.Env == "test"
	}

	func (c config) GetPort () string {
		if os.Getenv("PORT") != "" {
			return os.Getenv("PORT")
		} else if c.App.Port != "" {
			return c.App.Port
		} else {
			return "8000"
		}
	}

	func IsRunningInCloudRun() bool {
		return os.Getenv("IS_RUNNING_IN_CLOUD_RUN") == "true"
	}

	func init () {
		viper.SetConfigType("yaml")
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		
		if err := viper.ReadConfig(bytes.NewBuffer(configYaml)); err != nil {
			panic(err)
		}

		if err := viper.Unmarshal(&Conf); err != nil {
			panic(err)
		}
	}

