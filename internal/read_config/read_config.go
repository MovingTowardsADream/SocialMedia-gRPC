package read_config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Config struct {
	Env  string     `yaml:"env" env-default:"local"`
	GRPC GRPCConfig `yaml:"grpc"`
	DB   DBConfig   `yaml:"db"`
}

type GRPCConfig struct {
	Port     int           `yaml:"port"`
	TimeOut  time.Duration `yaml:"timeout"`
	TokenTLL time.Duration `yaml:"token_tll"`
}

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	UserName string `yaml:"username"`
	Name     string `yaml:"name"`
	Password string `yaml:"password"`
	SSLMode  string `yaml:"ssl_mode"`
}

func OpenConfig() *Config {
	path := getConfigPath()
	if path == "" {
		panic("config path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file does not exist: " + path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}
	return &cfg
}

func getConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "./configs/config.yaml", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
