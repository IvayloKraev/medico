package utils

import (
	"github.com/caarlos0/env/v11"
	"time"
)

// DATABASE CONFIGURATION

type DatabaseConfig struct {
	Host     string            `env:"DB_HOST,required,notEmpty"`
	Port     uint16            `env:"DB_PORT,required,notEmpty"`
	DBName   string            `env:"DB_NAME,required,notEmpty"`
	Username string            `env:"DB_USERNAME,required,notEmpty"`
	Password string            `env:"DB_PASSWORD,required,notEmpty"`
	Params   map[string]string `env:"DB_PARAMS" envSeparator:"&" envKeyValSeparator:"="`
}

var internalDatabaseConfig *DatabaseConfig

func LoadDatabaseConfig() {
	internalDatabaseConfig = &DatabaseConfig{}
	err := env.Parse(internalDatabaseConfig)
	if err != nil {
		panic(err)
	}
}

func GetDatabaseConfig() *DatabaseConfig {
	if internalDatabaseConfig == nil {
		LoadDatabaseConfig()
	}

	return internalDatabaseConfig
}

// CSRF CONFIGURATION

type CSRFConfig struct {
	Host           string        `env:"CSRF_HOST,required,notEmpty"`
	Port           int           `env:"CSRF_PORT,required,notEmpty"`
	Reset          bool          `env:"CSRF_RESET,required,notEmpty"`
	Username       string        `env:"CSRF_USERNAME,required,notEmpty"`
	Database       int           `env:"CSRF_DATABASE,required,notEmpty"`
	CookieName     string        `env:"CSRF_COOKIE_NAME,required,notEmpty"`
	SingleUseToken bool          `env:"CSRF_SINGLE_USE_TOKEN,required,notEmpty"`
	Expiration     time.Duration `env:"CSRF_EXPIRATION,required,notEmpty"`
}

var internalCSRFConfig *CSRFConfig

func LoadCSRFConfig() {
	internalCSRFConfig = &CSRFConfig{}
	err := env.Parse(internalCSRFConfig)
	if err != nil {
		panic(err)
	}
}

func GetCSRFConfig() *CSRFConfig {
	if internalCSRFConfig == nil {
		LoadCSRFConfig()
	}
	return internalCSRFConfig
}

// AUTHENTICATION SESSION CONFIGURATION

type AuthSessionConfig struct {
	Host       string        `env:"AUTH_SESSION_HOST,required,notEmpty"`
	Port       int           `env:"AUTH_SESSION_PORT,required,notEmpty"`
	Username   string        `env:"AUTH_SESSION_USERNAME,required,notEmpty"`
	Reset      bool          `env:"AUTH_SESSION_RESET,required,notEmpty"`
	Database   int           `env:"AUTH_SESSION_DATABASE,required,notEmpty"`
	CookieName string        `env:"AUTH_SESSION_COOKIE_NAME,required,notEmpty"`
	Expiration time.Duration `env:"AUTH_SESSION_EXPIRATION,required,notEmpty"`
}

var internalAuthSessionConfig *AuthSessionConfig

func LoadAuthSessionConfig() {
	internalAuthSessionConfig = &AuthSessionConfig{}
	err := env.Parse(internalAuthSessionConfig)
	if err != nil {
		panic(err)
	}
}

func GetAuthSessionConfig() *AuthSessionConfig {
	if internalAuthSessionConfig == nil {
		LoadAuthSessionConfig()
	}
	return internalAuthSessionConfig
}

// HASHING CONFIGURATION

type HashingConfig struct {
	Cost int `env:"HASHING_COST" envDefault:"16"`
}

var internalHashingCost *HashingConfig

func LoadHashingCost() {
	internalHashingCost = &HashingConfig{}
	err := env.Parse(internalHashingCost)
	if err != nil {
		panic(err)
	}
}

func GetHashingConfig() *HashingConfig {
	if internalHashingCost == nil {
		LoadHashingCost()
	}

	return internalHashingCost
}
