package config

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"sync"

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
	once     sync.Once
)

func InitialLoad() *Config {
	once.Do(func() {
		// 1. Determine the path
		configPath := getConfigPath()

		// 2. Verify existence
		if _, err := os.Stat(configPath); os.IsNotExist(err) {
			log.Fatalf("Critical: Config file not found at: %s", configPath)
		}

		// 3. Read config
		var cfg Config
		err := cleanenv.ReadConfig(configPath, &cfg)
		if err != nil {
			log.Fatalf("Critical: Failed to parse config file: %v", err)
		}

		instance = &cfg
	})
	return instance
}

func getConfigPath() string {
	// Priority 1: Environment Variable
	if envPath := os.Getenv("CONFIG_PATH"); envPath != "" {
		return envPath
	}

	// Priority 2: Command line flag
	flagPath := flag.String("config", "", "path to the configuration file")
	flag.Parse()
	if *flagPath != "" {
		return *flagPath
	}

	// Priority 3: Current Working Directory (Default)
	cwd, _ := os.Getwd()
	localPath := filepath.Join(cwd, "config.yaml")
	if _, err := os.Stat(localPath); err == nil {
		return localPath
	}

	// Priority 4: Executable Directory (Fallback for production)
	exePath, _ := os.Executable()
	exeDir := filepath.Dir(exePath)
	return filepath.Join(exeDir, "config.yaml")
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
	FuncId  int    `yaml:"id"`
	Title   string `yaml:"title"`
	Tooltip string `yaml:"tooltip"`
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
