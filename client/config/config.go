package config

import (
	"flag"
	"fmt"
)

var c = &Config{}

var (
	defaultServerURL = "http://localhost:8080"
	defaultFilePath  = "cotacao.txt"
)

type Config struct {
	ServerURL string
	FilePath  string
}

// GetConfig returns the app configuration
func GetConfig() *Config {
	return c
}

// ParseConfig parses the app configuration flags into the config object
func ParseConfig() {
	msg := fmt.Sprintf("The app's server URL (default: %s)", defaultServerURL)
	flag.StringVar(&c.ServerURL, "server-url", defaultServerURL, msg)

	msg = fmt.Sprintf("The file to write the dollar bid (default: %s)", defaultFilePath)
	flag.StringVar(&c.FilePath, "file-path", defaultFilePath, msg)

	flag.Parse()
}
