package config

import (
	"gopkg.in/yaml.v3"
	"os"
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
