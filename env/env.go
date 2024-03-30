
package env

import (
	"json"
	"log"
	"os"
)

type EnvVariables struct {
	Env      string      `json:"env"`
	AppName  string      `json:"app_name"`
	Database EnvDatabase `json:"database"`
}

type EnvDatabase struct {
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type EnvService struct {
	Env  EnvVariables
	Path string
}

func NewEnvService(path string) *EnvService {
	return &EnvService{
		Path: path,
	}
}

func (e *EnvService) Load() {
	configuration := EnvVariables{}
	configFile, err := os.Open(e.Path)
	if err != nil {
		log.Fatal("opening config file", err.Error())
	}

	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&configuration); err != nil {
		log.Fatal("parsing config file", err.Error())
	}
	e.Env = configuration
}

func (e *EnvService) GetVariables() *EnvVariables {
	return e.Env
}
