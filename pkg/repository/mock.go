package repository

import (
	"errors"

	"go.ideatocode.tech/config/pkg/interfaces"
)

// Mock .
type Mock struct {
	ReturnError bool
}

// Load .
func (r Mock) Load(cfg interfaces.Manager, data interface{}) error {
	if r.ReturnError {
		return errors.New("let's not make a mockery out of this")
	}
	return nil
}

// Save .
func (r Mock) Save(cfg interfaces.Manager, data interface{}) error {
	if r.ReturnError {
		return errors.New("let's not make a mockery out of this")
	}

	return nil
}

// Erase .
func (r Mock) Erase(cfg interfaces.Manager) error {
	if r.ReturnError {
		return errors.New("let's not make a mockery out of this")
	}
	return nil
}
