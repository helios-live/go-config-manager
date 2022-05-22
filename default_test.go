package config_test

import (
	"testing"

	"github.com/magiconair/properties/assert"
	"go.ideatocode.tech/config"
)

func TestConfigPath(t *testing.T) {
	c1 := config.New("file:./example.txt")

	assert.Equal(t, "file:./example.txt", c1.Path())

	c2 := config.New("http://example.com/example.txt")

	assert.Equal(t, "http://example.com/example.txt", c2.Path())

	c3 := config.New("https://example.com/example.txt")

	assert.Equal(t, "https://example.com/example.txt", c3.Path())

	c4 := config.New("https://a:b@example.com/example.txt?1=2")

	assert.Equal(t, "https://a:b@example.com/example.txt?1=2", c4.Path())
}
