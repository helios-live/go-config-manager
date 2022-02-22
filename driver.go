package apiconfig

import (
	"flag"
	"log"

	"github.com/fatih/color"
)

type plugin func(Config ConfigurationInterface) syncFunc
type syncFunc func(Config ConfigurationInterface) error

// Driver is the current driver used to store the configs
var Driver = flag.String("gcm-driver", "jsonc", "The apiconfig source to use")

var plugins map[string]plugin
var pluginsInitialized = false

func addPlugin(name string, call plugin) {
	initPlugins()
	plugins[name] = call
	log.Println(color.YellowString("Loading"), color.GreenString(name), "apiconfig plugin")
}
func initPlugins() bool {
	if pluginsInitialized {
		return false
	}
	plugins = make(map[string]plugin)
	pluginsInitialized = true
	return true
}
