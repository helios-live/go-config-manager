package config_test

import (
	"fmt"
	"net"
	"net/http"
	"os"
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
func TestCanEraseHTTP(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	assert.NoError(t, err)
	defer listener.Close()

	http.HandleFunc("/examples/options/data/config.jsonc", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})

	addr := listener.Addr().(*net.TCPAddr)
	go func() {
		http.Serve(listener, nil)
	}()

	c1 := config.New(fmt.Sprintf("http://%s/examples/options/data/config.jsonc", addr.String()))

	err = config.Erase(c1)

	assert.NoError(t, err)
}
func TestCanEraseFile(t *testing.T) {
	fpath := "./examples/options/data/tmp.jsonc"

	d1 := []byte("{\"some\":\"data\"}")
	err := os.WriteFile(fpath, d1, 0644)
	assert.NoError(t, err)

	x := make(map[string]string)
	c1 := config.New("file:" + fpath)
	config.Load(c1, &x)

	assert.Equal(t, "data", x["some"])

	err = config.Erase(c1)

	assert.NoError(t, err)

	_, err = os.Stat(fpath)
	assert.ErrorIs(t, err, os.ErrNotExist)
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
