package config

import (
	"flag"
	"github.com/go-playground/validator/v10"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	*yamlConfig
	*envConfig
}

type yamlConfig struct {
	Port int `mapstructure:"port"`
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
		log.Panic("error to get .env")
	}

	configPath := fetchConfigPath()
	if configPath == "" {
		log.Panic("error to init config path")
	}

	_, err = os.Stat(configPath)
	if os.IsNotExist(err) {
		log.Panicf("config file does not exist: %s", configPath)
	}

	viper.AddConfigPath(filepath.Dir(configPath))
	viper.SetConfigType("yaml")
	viper.SetConfigName(filepath.Base(configPath))

	err = viper.ReadInConfig()
	if err != nil {
		log.Panicf("failed to read config file: %w", err)
	}

	var yamlCfg yamlConfig

	err = viper.Unmarshal(&yamlCfg)
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

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
