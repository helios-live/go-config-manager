package interfaces

import (
	"sync"

	"github.com/go-kit/log"
)

// Manager .
// must be able to Save options,
// must be able to Load options,
// must be able to use multiple "drivers"
// must be thread safe
// must be able to use multiple logical/physical configs
type Manager interface {

	// ConfigurationInterface needs to implement Locker
	sync.Locker

	Repository() Repository

	Marshaler() Marshaler

	Path() string

	Logger() log.Logger
}

// Repository is responsible for interacting with the storage system
type Repository interface {

	// Load is called to load options
	Load(Manager, interface{}) error

	// Save is called to save options
	Save(Manager, interface{}) error
}

// Marshaler is responsible for serializing and deserializing data
type Marshaler interface {
	Marshal(v interface{}) (data []byte, err error)
	Unmarshal(data []byte, v interface{}) error
}
