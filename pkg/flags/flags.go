package flags

import (
	"flag"

	"go.ideatocode.tech/config"
	"go.ideatocode.tech/config/pkg/marshal"
	"go.ideatocode.tech/config/pkg/repository"
)

var repo = flag.String("repository", "", "Forcefully use this repository [file, http]")
var marsh = flag.String("marshaler", "", "Forcefully use this marshaler [jsonc, yaml]")
var httpSource = flag.String("http-ds", "", "Forcefully use HTTP source to get and save the config from")
var httpToken = flag.String("http-auth", "", "Forcefully use HTTP Authorization: Bearer Token Header")

// Wrap returns a new config.Default with the options changed by the flags
func Wrap(cfg *config.DefaultManager) *config.DefaultManager {
	switch *repo {
	case "file":
		cfg.Repo = repository.File{}

	case "http":
		cfg.Repo = repository.HTTP{
			Token: *httpToken,
			URL:   *httpSource,
		}
	}

	switch *marsh {
	case "jsonc":
		cfg.Marsh = marshal.JSONC{}

	case "yaml":
		cfg.Marsh = marshal.YAML{}
	}
	return cfg
}
