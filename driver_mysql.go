package apiconfig

import (
	"flag"
	"log"
)

var dsn = flag.String("mysql-dsn", "", "The mysql Data Source Name. I.e. user:password@tcp(your-amazonaws-uri.com:3306)/dbname")

func init() {
	addPlugin("mysql", loadMySQL)
}

func loadMySQL(Config ConfigurationInterface) syncFunc {
	log.Fatalln("Not yet implemented")

	return syncMySQL
}

func syncMySQL(Config ConfigurationInterface) error {
	return nil
}
