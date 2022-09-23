package conf

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type DBConfiguration struct {
	URL string `json:"url"`
}

type APIConfiguration struct {
	Host        string `envconfig:"HOST"`
	Port        int    `envconfig:"PORT" default:"8081"`
	ExternalURL string `json:"external_url" envconfig:"API_EXTERNAL_URL"`
}

type GlobalConfiguration struct {
	API APIConfiguration
	DB  DBConfiguration
}

func LoadGlobal(filename string) (*GlobalConfiguration, error) {
	if err := loadEnvironment(filename); err != nil {
		return nil, err
	}

	config := new(GlobalConfiguration)
	if err := envconfig.Process("app", config); err != nil {
		return nil, err
	}

	return config, nil
}

func loadEnvironment(filename string) error {
	var err error
	if filename != "" {
		err = godotenv.Overload(filename)
	} else {
		err = godotenv.Load()
		// handle if .env file does not exist, this is OK
		if os.IsNotExist(err) {
			return nil
		}
	}
	return err
}
