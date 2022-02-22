package apiconfig

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/muhammadmuzzammil1998/jsonc"
)

func init() {
	addPlugin("jsonc", loadJsonc)
}

func loadJsonc(Config ConfigurationInterface) SyncFuncDef {
	// log.Fatalln("Not yet implemented")

	jsonFile, err := os.Open(jsonGetPath(Config))
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatalln(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	jsonFile.Close()

	byteValue = jsonc.ToJSON(byteValue)
	err = json.Unmarshal(byteValue, &Config)
	if err != nil {
		log.Fatalln(err)
	}

	return syncJSONC
}

func syncJSONC(Config ConfigurationInterface) error {
	log.Println("syncJSONC: Flushing changes to disk")
	Config.Lock()
	b, err := json.MarshalIndent(Config, "", "\t")
	Config.Unlock()
	if err != nil {
		log.Panicf("Json Marshal Error: %s", err)
	}
	err = ioutil.WriteFile(jsonGetPath(Config), b, 0644)

	if err != nil {
		log.Panicf("Failed to write e: %s, p: %s", err, jsonGetPath(Config))
	}

	return nil
}

func jsonGetPath(Config ConfigurationInterface) string {
	p := Config.GetParent()

	return p.Group + "/" + p.Item + ".jsonc"
}
