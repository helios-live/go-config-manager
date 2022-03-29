package main

import (
	"flag"
	"os"

	"github.com/fatih/color"
	"github.com/go-kit/log"
	"go.ideatocode.tech/config"
	"go.ideatocode.tech/config/pkg/flags"
	"go.ideatocode.tech/config/pkg/marshal"
	"go.ideatocode.tech/config/pkg/repository"
)

type mainConfig struct {
	Token   string
	Counter int
}

func main() {

	flag.Parse()
	logger := log.NewLogfmtLogger(os.Stderr)

	// json File
	cfg := config.New("data/config.jsonc")
	mc := mainConfig{}

	err := config.Load(cfg, &mc)

	if err != nil {
		panic(color.RedString(err.Error()))
	}
	logger.Log("type", "default", "Token", mc.Token, "Counter", mc.Counter)

	// YAML file
	yc := config.New("data/config.yaml", func(d *config.DefaultManager) {
		d.Repo = repository.File{}
		d.Marsh = marshal.YAML{}
	})

	y := mainConfig{}

	err = config.Load(yc, &y)
	if err != nil {
		panic(color.RedString(err.Error()))
	}
	logger.Log("type", "yaml", "Token", y.Token, "Counter", y.Counter)

	// increase the counter
	y.Counter++
	config.Save(yc, &y)

	// direct access, json wrapped with flags
	fc := flags.Wrap(config.New("data/config.jsonc"))

	wrap := mainConfig{}

	err = fc.Load(&wrap)

	if err != nil {
		panic(color.RedString(err.Error()))
	}
	logger.Log("type", "direct access, flags wrapped", "Token", wrap.Token, "Counter", wrap.Counter)

	// JSON over http
	jc := config.New("data/config.jsonc", func(d *config.DefaultManager) {
		d.Repo = repository.HTTP{
			Token: "ZemExincRT6FgfvQWflCB8t1MTC8xOl4y1SwyAjmx7nl7WpdRzv0mZrgTr7nm0GJ",
			URL:   "https://peertonet.test/api/static-proxy/config",
		}
	})

	h := mainConfig{}

	err = config.Load(jc, &h)
	if err != nil {
		panic(color.RedString(err.Error()))
	}
	logger.Log("type", "http", "Token", h.Token, "Counter", h.Counter)

}
