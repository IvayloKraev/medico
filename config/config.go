package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

type DBMS string

const (
	MySQL   DBMS = "mysql"
	MariaDB DBMS = "mariadb"
)

const (
	databaseConfigPath    = "./config/database.config.yml"
	csrfStorageConfigPath = "./config/csrf.config.yml"
	authSessionConfigPath = "./config/authSession.config.yml"
)

type AuthSessionConfig struct {
	Host       string        `yaml:"host"`
	Port       int           `yaml:"port"`
	Username   string        `yaml:"username"`
	Reset      bool          `yaml:"reset"`
	Database   int           `yaml:"database"`
	CookieName string        `yaml:"cookie_name"`
	Expiration time.Duration `yaml:"expiration"`
}

func loadConfig(configPath string, out interface{}) {
	configFile, err := os.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(configFile, out)
	if err != nil {
		panic(err)
	}
}

func LoadAuthSessionConfig() *AuthSessionConfig {
	authSessionConfig := &AuthSessionConfig{}
	loadConfig(authSessionConfigPath, &authSessionConfig)
	return authSessionConfig
}
