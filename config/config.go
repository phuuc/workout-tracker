package config

import (
	"fmt"
	"os"

	"github.com/finnpn/workout-tracker/pkg/log"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Environment string  `yaml:"environment"`
	Server      *Server `yaml:"server"`
	Mysql       *Mysql  `yaml:"mysql"`
}

type Server struct {
	ApiHost string `yaml:"api_host"`
	ApiPort int    `yaml:"api_port"`
}

type Mysql struct {
	DriverName                 string `yaml:"driver_name"`
	MaxOpenConns               int    `yaml:"max_open_conns"`
	MaxIdleConns               int    `yaml:"max_idle_conns"`
	ConnMaxLifeTimeMiliseconds int64  `yaml:"conn_max_life_time_miliseconds"`
	//MigrationConnURL           string `yaml:"migration_conn_url"`
	IsDevMode bool   `yaml:"is_dev_mode"`
	UserName  string `yaml:"user_name"`
	Passwd    string `yaml:"passwd"`
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	Name      string `yaml:"name"`
}

func Load(filePath string) (*Config, error) {
	if len(filePath) == 0 {
		filePath = os.Getenv("CONFIG_FILE")
	}
	log.Info("loading config file with filepath =%s", filePath)

	configBytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Error("could not load config file with err =%v", err)
		return nil, err
	}

	configBytes = []byte(os.ExpandEnv(string(configBytes)))

	cfg := &Config{}

	err = yaml.Unmarshal(configBytes, cfg)
	if err != nil {
		log.Error("Failed to parse config file with err=%v", err)
		return nil, err
	}
	log.Info("done parsing config ===========")
	log.Info("environment %s", cfg.Environment)
	log.Info("mysql =%v", *cfg.Mysql)
	log.Info("server =%v", *cfg.Server)

	return cfg, nil
}

func (c *Config) Addr(host string, port int) string {
	return fmt.Sprintf("%s:%d", host, port)
}
