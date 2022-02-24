package apiconfig

import (
	"flag"
	"log"
	"sync"
)

func init() {
	flag.Parse()
}

// ConfigFile is the file where the settings are stored
var ConfigFile = "./config.json"

// ConfigurationInterface .
type ConfigurationInterface interface {

	// ConfigurationInterface needs to implement Locker
	sync.Locker

	// ConfigurationInterface needs to have a root AuthToken
	AuthToken() string
	// setObj(interface{})
	// GetPath() string

	// Sync is called to sync changes
	Sync()

	// SyncFunc gets and sets the SyncFuncDef
	// SyncFunc(SyncFuncDef) sets it
	// SyncFunc(nil) loads it
	// the function that is called when you call Sync on the configuration
	SyncFunc(SyncFuncDef) SyncFuncDef

	// GetParent
	GetParent() *Configuration
}

// Configuration is the base Configuration object
type Configuration struct {
	*sync.Mutex `json:"-"`
	Token       string `json:"AuthToken,omitempty"`

	// Group represents a logical grouping
	// for JSON it'll be used as the folder
	// for MySql it'll be used as the database
	Group string

	// Item represents this whole configuration object
	// for JSON it'll be the file name
	// for MySql it'll be the table name
	Item string

	syncFunc SyncFuncDef
}

// NewConfig returns a pointer to a filled new instance of Configuration
func NewConfig(file string) *Configuration {
	if len(file) == 0 {
		file = ConfigFile
	}
	return &Configuration{
		Mutex: &sync.Mutex{},
		// configFile: file,
	}
}

// SyncFunc return the root authToken
func (c *Configuration) SyncFunc(sf SyncFuncDef) SyncFuncDef {
	if sf != nil {
		c.syncFunc = sf
	}

	return c.syncFunc
}

// AuthToken return the root authToken
func (c *Configuration) AuthToken() string {
	return c.Token
}

// func (c *Configuration) setObj(obj interface{}) {
// 	c = obj
// }

// GetPath returns the path to the configFile
// func (c *Configuration) GetPath() string {
// 	return c.configFile
// }

// LoadConfig loads the config and return the fiilled object
func LoadConfig(Config ConfigurationInterface) {
	// spew.Dump(Config)
	// spew.Dump(plugins)
	// spew.Dump(*Driver)
	sf := plugins[*Driver](Config)
	Config.SyncFunc(sf)
}

// LoadConfig loads the config and return the fiilled object
// func (c *Configuration) LoadConfig(Config ConfigurationInterface) ConfigurationInterface {

// 	jsonFile, err := os.Open(c.configFile)
// 	// if we os.Open returns an error then handle it
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	byteValue, _ := ioutil.ReadAll(jsonFile)
// 	jsonFile.Close()

// 	byteValue = jsonc.ToJSON(byteValue)

// 	err = json.Unmarshal(byteValue, &Config)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	Config.setObj(Config)
// 	return Config
// }

// Sync Writes the config to disk
func Sync(Config ConfigurationInterface) {
	sf := Config.SyncFunc(nil)
	err := sf(Config)
	if err != nil {
		log.Println("GCM sync error")
	}
}

// Sync Writes the config to disk
// Presumably after you've changed it but it does not do any checks
func (c *Configuration) Sync() {
	Sync(c)
}

// GetParent Writes the config to disk
// Presumably after you've changed it but it does not do any checks
func (c *Configuration) GetParent() *Configuration {
	return c
}

// c.Lock()
// b, err := json.MarshalIndent(c, "", "\t")
// c.Unlock()
// if err != nil {
// 	log.Panicf("Json Marshal Error: %s", err)
// }

// 	err = ioutil.WriteFile(c.configFile, b, 0644)

// 	if err != nil {
// 		log.Panicf("Failed to write e: %s, p: %s", err, c.GetPath())
// 	}
// }
