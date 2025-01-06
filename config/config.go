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
	databaseConfigPath = "./config/database.config.yml"
)

type DatabaseConfig struct {
	DBMS      DBMS
	Host      string
	Port      uint16
	DBName    string
	Username  string
	Password  string
	Migration bool
}

func LoadDatabaseConfig() *DatabaseConfig {
	databaseConfigFile, err := os.ReadFile(databaseConfigPath)
	if err != nil {
		panic(err)
	}

	databaseConfig := &DatabaseConfig{}
	if err := yaml.Unmarshal(databaseConfigFile, databaseConfig); err != nil {
		panic(err)
	}

	return databaseConfig
}
