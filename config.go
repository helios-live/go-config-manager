package config

import (
	"sync"

	"github.com/go-kit/log"

	"go.ideatocode.tech/config/pkg/interfaces"
	"go.ideatocode.tech/config/pkg/marshal"
	"go.ideatocode.tech/config/pkg/repository"
)

// New returns a pointer to a filled new instance of Configuration
func New(path string, options ...func(*DefaultManager)) *DefaultManager {
	cfg := &DefaultManager{
		Mutex:    &sync.Mutex{},
		FilePath: path,
	}

	for _, option := range options {
		option(cfg)
	}

	// use the fileRepository by default
	if cfg.Repo == nil {
		cfg.Repo = repository.File{}
	}

	// use the JSONC Marhaller by default
	if cfg.Marsh == nil {
		cfg.Marsh = marshal.JSONC{}
	}

	// use the NopLogger by default
	if cfg.Log == nil {
		cfg.Log = log.NewNopLogger()
	}

	return cfg
}

// Save .
func Save(config interfaces.Manager, data interface{}) error {
	return config.Repository().Save(config, data)
}

// Load .
func Load(config interfaces.Manager, data interface{}) error {
	return config.Repository().Load(config, data)
}
