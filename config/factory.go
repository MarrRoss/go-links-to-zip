package config

import (
	"github.com/MarrRoss/go-links-to-zip/config/api"
)

type Factory struct {
	API api.Config
}

func New() *Factory {
	cfg := Load()
	return cfg
}
