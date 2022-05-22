package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.ideatocode.tech/config"
)

func TestCanLoadFile(t *testing.T) {
	c1 := config.New("file:./examples/options/data/config.jsonc")

	var x map[string]interface{}

	err := config.Load(c1, &x)

	assert.NoError(t, err)
}
func TestConfigPath(t *testing.T) {
	c1 := config.New("file:./example.txt")

	assert.Equal(t, "file:./example.txt", c1.FullURL())

	c2 := config.New("http://example.com/example.txt")

	assert.Equal(t, "http://example.com/example.txt", c2.FullURL())

	c3 := config.New("https://example.com/example.txt")

	assert.Equal(t, "https://example.com/example.txt", c3.FullURL())

	c4 := config.New("https://a:b@example.com/example.txt?1=2")

	assert.Equal(t, "https://a:b@example.com/example.txt?1=2", c4.FullURL())
}
