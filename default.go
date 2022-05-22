package config

import (
	"sync"

	"github.com/go-kit/log"
	"go.ideatocode.tech/config/pkg/interfaces"
)

// DefaultManager is the base DefaultManager object
type DefaultManager struct {
	*sync.Mutex

	// FilePath represents the path to the config file
	FilePath string `json:"-" yaml:"-"`

	// fullURL represents the full url passed to config.New
	fullURL string `json:"-" yaml:"-"`

	// Repo is responsible for interracting with the storage medium
	Repo interfaces.Repository `json:"-" yaml:"-"`

	// Marsh is responsible for serializing and deserializing data
	Marsh interfaces.Marshaler `json:"-" yaml:"-"`

	Log log.Logger `json:"-" yaml:"-"`

	// should we check the path extension and determine the appropriate
	// marshaller?
	// Default: true
	AutoDetectMarshaller bool `json:"-" yaml:"-"`
}

// Repository returns the repository
func (c DefaultManager) Repository() interfaces.Repository {
	return c.Repo
}

// Marshaler returns the repository
func (c DefaultManager) Marshaler() interfaces.Marshaler {
	return c.Marsh
}

// Logger returns the logger
func (c DefaultManager) Logger() log.Logger {
	return c.Log
}

// Path returns the path of the manager
func (c DefaultManager) Path() string {
	return c.FilePath
}

// Load .
func (c DefaultManager) Load(data interface{}) error {
	return c.Repository().Load(c, data)
}

// Save .
func (c DefaultManager) Save(data interface{}) error {
	return c.Repository().Save(c, data)
}

// FullURL returns the full url passed to config.New
func (c DefaultManager) FullURL() string {
	return c.fullURL
}
