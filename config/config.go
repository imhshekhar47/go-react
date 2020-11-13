package config

import (
	"log"
	"os"
	"path"

	"github.com/ilyakaznacheev/cleanenv"
)

// ApplicationConfiguration  application cionfig definition
type ApplicationConfiguration struct {
	Server struct {
		Host string `yaml:"host" env:"SERVER_HOST" env-default:"localhost"`
		Port string `yaml:"port" env:"SERVER_PORT" env-default:"8080"`
	} `yaml:"server"`

	Application struct {
		Name      string `yaml:"name" env:"APP_NAME" env-default:"sample-app"`
		Version   string `yaml:"version" env:"APP_VERSION" env-default:"SNAPSHOT"`
		Developer string `yaml:"developer" env:"APP_DEVELOPER" env-default:"imhs47"`
	} `yaml:"application"`
}

// AppConfig singleton instance of application configuration
var AppConfig *ApplicationConfiguration

// GetConfig returns application configuration
func GetConfig() *ApplicationConfiguration {

	if AppConfig == nil {
		AppConfig = &ApplicationConfiguration{}

		isDefaultConfig := true
		// check for configuration files
		// dev: ./config.yml
		if _, err := os.Stat("config.yml"); err == nil || os.IsExist(err) {
			if err := cleanenv.ReadConfig("config.yml", AppConfig); err != nil {
				log.Println("[WARN] Could not load configuration from file ./config.yml")
			} else {
				isDefaultConfig = false
			}
		}

		// prod: ${APP_HOME}/config.yml
		appHome, exist := os.LookupEnv("APP_HOME")

		if exist {
			cfgPath := path.Join(appHome, "config.yml")
			if err := cleanenv.ReadConfig(cfgPath, AppConfig); err != nil {
				log.Println("[WARN] Could not load configuration from file ./config.yml")
			} else {
				isDefaultConfig = false
			}
		}

		if isDefaultConfig {
			log.Println("[WARN] No configuration file provided. Using with default config")
		}

	}

	return AppConfig
}
