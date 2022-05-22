package repository

import (
	"fmt"
	"io/ioutil"
	"os"

	"go.ideatocode.tech/config/pkg/interfaces"
)

// File .
type File struct{}

// Load .
func (r File) Load(cfg interfaces.Manager, data interface{}) error {

	file, err := os.Open(cfg.Path())
	if err != nil {
		return err
	}
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)

	err = cfg.Marshaler().Unmarshal(byteValue, data)

	if err != nil {
		return fmt.Errorf("File Unmarshal Error: %s", err)
	}

	return err
}

// Save .
func (r File) Save(cfg interfaces.Manager, data interface{}) error {
	cfg.Logger().Log("File: Flushing changes to disk")
	cfg.Lock()
	defer cfg.Unlock()
	b, err := cfg.Marshaler().Marshal(data)
	if err != nil {
		return fmt.Errorf("Json Marshal Error: %s", err)
	}
	err = ioutil.WriteFile(cfg.Path(), b, 0644)

	if err != nil {
		return fmt.Errorf("Failed to write e: %s, p: %s", err, cfg.Path())
	}

	return nil
}

// Erase .
func (r File) Erase(cfg interfaces.Manager) error {
	return os.Remove(cfg.Path())
}
