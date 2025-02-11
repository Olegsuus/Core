package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/subosito/gotenv"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	*yamlConfig
	*envConfig
}

type yamlConfig struct {
	Server  serverConfig  `yaml:"server"`
	Log     logConfig     `yaml:"log"`
	Metrics metricsConfig `yaml:"metrics"`
}

type serverConfig struct {
	Port int `yaml:"port"`
}

type logConfig struct {
	LogFilePath string `yaml:"log_file_path"`
}

type metricsConfig struct {
	Port int `yaml:"port"`
}

type envConfig struct {
	Env string `env:"env"`
	DB  dbConfig
}

type dbConfig struct {
	Host     string `env:"POSTGRES_HOST"`
	Port     int    `env:"POSTGRES_PORT"`
	DBName   string `env:"POSTGRES_DB"`
	User     string `env:"POSTGRES_USER"`
	Password string `env:"POSTGRES_PASSWORD"`
}

func MustConfig() *Config {
	err := gotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using environment variables from OS")
	}

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Panic("error to init config path")
	}

	_, err = os.Stat(configPath)
	if os.IsNotExist(err) {
		log.Panicf("config file does not exist: %s", configPath)
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		log.Panicf("failed to read config file: %v", err)
	}

	var yamlCfg yamlConfig

	err = yaml.Unmarshal(data, &yamlCfg)
	if err != nil {
		log.Panicf("failed to unmarshal config: %w", err)
	}

	validate := validator.New()
	err = validate.Struct(yamlCfg)
	if err != nil {
		log.Panicf("failed to validate config: %w", err)
	}

	var envCfg envConfig
	err = cleanenv.ReadEnv(&envCfg)
	if err != nil {
		log.Panicf("failed to read evn config: %w", err)
	}

	return &Config{
		&yamlCfg,
		&envCfg,
	}
}
