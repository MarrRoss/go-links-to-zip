package config

import (
	"fmt"
	"github.com/MarrRoss/go-links-to-zip/config/api"
)

type Factory struct {
	API api.Config
}

func New() *Factory {
	cfg := Load()
	fmt.Printf("Host: %v, Port: %v", cfg.API.Host, cfg.API.Port)
	return cfg
}
