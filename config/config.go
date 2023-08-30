package config

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

const (
	// EnvLocal is the local environment
	EnvLocal = "LOCAL"
	// EnvContainer is the container environment
	EnvLocalContainer = "LOCAL_CONTAINER"
	// EnvDev is the development environment
	EnvDev = "DEV"
	// EnvProd is the production environment
	EnvProd = "PROD"
)

var AppConfig *App

type App struct {
	Server   Server   `mapstructure:"server"`
	Database Database `mapstructure:"database"`
}

type Server struct {
	Port string `mapstructure:"port"`
	Name string `mapstructure:"name"`
	Env  string `mapstructure:"env"`
}

type Database struct {
	Postgres `mapstructure:"postgres"`
}

type Postgres struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"user"`
	Password string
	Database string `mapstructure:"dbname"`
}

func (p Postgres) GetDSN() string {
	var builder strings.Builder
	builder.WriteString("host=")
	builder.WriteString(p.Host)
	builder.WriteString(" port=")
	builder.WriteString(p.Port)
	builder.WriteString(" user=")
	builder.WriteString(p.Username)
	builder.WriteString(" password=")
	builder.WriteString(p.Password)
	builder.WriteString(" dbname=")
	builder.WriteString(p.Database)
	builder.WriteString(" sslmode=disable")

	dsn := builder.String()

	return dsn
}

func LoadConfig(path string) (*App, error) {
	var filename string
	appEnv := os.Getenv("APPENV")
	switch appEnv {
	case EnvDev:
		filename = "config-dev.yml"
	case EnvProd:
		filename = "config-prod.yml"
	case EnvLocalContainer:
		filename = "config-container.yml"
	case "integration-test":
		filename = "config-integration-test.yml"
	default:
		filename = "config-local.yml"
	}

	fullPath := path + "/" + filename
	log.Println("Loading config file: ", fullPath)
	return ViperLoadConfig(fullPath)
}

func ViperLoadConfig(fullPath string) (*App, error) {
	viper.SetConfigFile(fullPath)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config App
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	config.Database.Postgres.Password = os.Getenv("POSTGRES_PASSWORD")

	return &config, nil
}
