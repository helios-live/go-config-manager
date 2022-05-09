package config

import (
	"net/url"
	"path/filepath"
	"sync"

	"github.com/go-kit/log"

	"go.ideatocode.tech/config/pkg/interfaces"
	"go.ideatocode.tech/config/pkg/marshal"
	"go.ideatocode.tech/config/pkg/repository"
)

// New returns a pointer to a filled new instance of Configuration
func New(path string, options ...func(*DefaultManager)) *DefaultManager {

	cfg := &DefaultManager{
		Mutex:                &sync.Mutex{},
		FilePath:             path,
		AutoDetectMarshaller: true,
	}

	rd := &RepositoryDefinition{}
	err := rd.UnmarshalText(path)
	if err == nil {
		switch rd.url.Scheme {
		case "file":
			cfg.Repo = repository.File{}
			cfg.FilePath = rd.url.Path
		case "http", "https":

			// strip user:pass from the url
			stripped := rd.url.String()
			su, _ := url.Parse(stripped)
			su.User = nil
			stripped = su.String()

			pwd, _ := rd.url.User.Password()
			cfg.Repo = repository.HTTP{
				URL:   stripped,
				Token: rd.url.User.Username() + pwd,
			}
			cfg.FilePath = rd.url.Path
		}
	}

	for _, option := range options {
		option(cfg)
	}
	if cfg.AutoDetectMarshaller {
		path := rd.url.Path
		ext := filepath.Ext(path)

		switch ext {
		case ".json", ".jsonc":
			cfg.Marsh = marshal.JSONC{}
		case ".yaml":
			cfg.Marsh = marshal.YAML{}
		}
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
