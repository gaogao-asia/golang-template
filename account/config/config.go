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
	EnvContainer = "CONTAINER"
	// EnvDev is the development environment
	EnvDev = "DEV"
	// EnvProd is the production environment
	EnvProd = "PROD"
)

var AppConfig *App

type App struct {
	Server   Server   `mapstructure:"server"`
	Database Database `mapstructure:"database"`
	Monitor  Monitor  `mapstructure:"monitor"`
}

type Server struct {
	Port string `mapstructure:"port"`
	Name string `mapstructure:"name"`
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

type Monitor struct {
	OpenTelemetry `mapstructure:"opentelemetry"`
	Tempo         `mapstructure:"tempo"`
}

type OpenTelemetry struct {
	Enable bool `mapstructure:"enable"`
}
type Tempo struct {
	Endpoint string `mapstructure:"endpoint"`
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
	case EnvContainer:
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
	log.Println("config.Database.Postgres.Password : ", config.Database.Postgres.Password)

	return &config, nil
}
