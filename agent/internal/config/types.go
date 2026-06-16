package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	App      AppConfig        `yaml:"app"`
	Server   []ServerConfig   `yaml:"server"`
	Database []DatabaseConfig `yaml:"database"`
	Services ServicesConfig   `yaml:"services,omitempty"`
	Logging  LoggingConfig    `yaml:"logging,omitempty"`
}

var (
	instance *Config
)

func InitialLoad() *Config {
	// 1. Try environment variable
	configPath := os.Getenv("CONFIG_PATH")

	// 2. Fallback to executable directory if not set or doesn't exist
	if configPath == "" || func() bool { _, err := os.Stat(configPath); return os.IsNotExist(err) }() {
		exePath, err := os.Executable()
		if err != nil {
			log.Fatal("Could not find executable path")
		}
		// Assign to the existing variable, don't use :=
		configPath = filepath.Join(filepath.Dir(exePath), "config.yaml")
	}

	var cfg Config
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("can not read config file: %s", err.Error())
	}

	log.Println("Config Path Setting: ", configPath)
	return &cfg
}

type AppConfig struct {
	Name        string          `yaml:"name" env:"APP_NAME" env-default:"go-gateway"`
	Version     string          `yaml:"version" env:"APP_VERSION" env-default:"1.0.0"`
	Environment string          `yaml:"environment" env:"APP_ENV" env-default:"development"`
	Title       string          `yaml:"title"`
	Tooltip     string          `yaml:"tooltip"`
	Tray        []TrayOptions   `yaml:"tray"`
	Dimension   DimensionConfig `yaml:"dimension"`
}

type TrayOptions struct {
	FuncId   int    `yaml:"id,omitempty"`
	Sperator bool   `yaml:"sperator,omitempty"`
	Title    string `yaml:"title,omitempty"`
	Tooltip  string `yaml:"tooltip,omitempty"`
}

type DimensionConfig struct {
	Width  int `yaml:"width"`
	Height int `yaml:"height"`
}

type ServerConfig struct {
	Tag    string `yaml:"tag"`
	Type   string `yaml:"type"`
	IP     string `yaml:"ip"`
	FQDN   string `yaml:"fqdn"`
	Domain string `yaml:"domain"`
	Port   string `yaml:"port"`
}

type DatabaseConfig struct {
	Tag      string `yaml:"tag"`
	Engine   string `yaml:"engine"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

type ServicesConfig struct {
}

type LoggingConfig struct {
}
