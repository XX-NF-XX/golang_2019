package main

import (
	"flag"
)

// Config - basic app configuration
type Config struct {
	ServerPort string
	ServerIP   string
}

var config Config

const defaultServerIP string = "[::1]"
const defaultServerPort string = "3000"

func (c *Config) getAddress() string {
	return c.ServerIP + ":" + c.ServerPort
}

func loadConfig() {
	flag.StringVar(&config.ServerPort, "port", defaultServerPort, "Server port")
	flag.StringVar(&config.ServerIP, "host", defaultServerIP, "Remote server IP. If IP isn't set, start local server.")

	flag.Parse()
}
