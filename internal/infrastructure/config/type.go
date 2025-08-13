package config

import (
	"github.com/AndreeJait/go-template-hexagonal/internal/constant"
	"os"
)

type Config struct {
	Service Service `yaml:"service"`
	DB      DB      `yaml:"db"`
}

func IsDevelopment() bool {
	var appMode = os.Getenv("APP_ENV")
	return constant.AppMode(appMode) == constant.DevelopmentMode ||
		constant.AppMode(appMode) == constant.StagingMode
}

func IsProduction() bool {
	var appMode = os.Getenv("APP_ENV")
	return constant.AppMode(appMode) == constant.ProductionMode
}

type Service struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Name string `yaml:"name"`
}

type DB struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Name     string `yaml:"name"`
	Password string `yaml:"password"`
}
