package config

import (
	"github.com/spf13/viper"
)

type LogLevel string

const (
	Debug    LogLevel = "debug"
	Info     LogLevel = "info"
	Warning  LogLevel = "warning"
	Critical LogLevel = "critical"
)

// Application holds the application configuration
type Application struct {
	Base     string
	Port     int
	LogLevel LogLevel
}

// app is the default application configuration
var app Application

// App returns the default application configuration
func App() *Application {
	return &app
}

// LoadApp loads application configuration
func LoadApp() {
	mu.Lock()
	defer mu.Unlock()

	app = Application{
		Base:     viper.GetString("base"),
		Port:     viper.GetInt("port"),
		LogLevel: LogLevel(viper.GetString("log_level")),
	}
}
