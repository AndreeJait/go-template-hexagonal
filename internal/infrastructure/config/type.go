package config

import (
	"github.com/AndreeJait/go-template-hexagonal/internal/constant"
	"github.com/AndreeJait/go-utility/emailw"
	"os"
)

type Config struct {
	Service     Service     `yaml:"service"`
	DB          DB          `yaml:"db"`
	InternalApi InternalApi `yaml:"internal_api"`
	Jwt         Jwt         `yaml:"jwt"`

	Email emailw.EmailConfig `yaml:"email"`

	Setting Setting `yaml:"setting"`
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

type Setting struct {
	ActivationTokenInMinute int64 `yaml:"activation_token_in_minute"`
}

type Service struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Name string `yaml:"name"`

	EmailSender      string `yaml:"email_sender"`
	RedirectFrontend string `yaml:"redirect_frontend"`
}

type DB struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Name     string `yaml:"name"`
	Password string `yaml:"password"`
}

type InternalApi struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Jwt struct {
	Secret                string `yaml:"secret"`
	TokenExpiredInMinutes int    `yaml:"token_expired_in_minutes"`
}
