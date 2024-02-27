package config

import (
	"flag"
	"fmt"
)

var c = &Config{}

var (
	defaultPort   = "8080"
	defaultDBPath = "./fullcycle.db"
)

type Config struct {
	Port   string
	DBPath string
}

// GetConfig returns the app configuration
func GetConfig() *Config {
	return c
}

// ParseConfig parses the app configuration flags into the config object
func ParseConfig() {
	msg := fmt.Sprintf("The app's server HTTP port (default: %s)", defaultPort)
	flag.StringVar(&c.Port, "port", defaultPort, msg)

	msg = fmt.Sprintf("The file path that the sqlite db should use to save data (default: %s)", defaultDBPath)
	flag.StringVar(&c.DBPath, "db-path", defaultDBPath, msg)
	flag.Parse()
}
