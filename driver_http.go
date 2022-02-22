package apiconfig

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// var dsn = flag.String("mysql-dsn", "", "The mysql Data Source Name. I.e. user:password@tcp(your-amazonaws-uri.com:3306)/dbname")
var httpSource = flag.String("gcm-http-ds", "", "The HTTP source to get and save the config from")
var httpToken = flag.String("gcm-http-auth", "", "The HTTP Authorization: Bearer Token Header")

func init() {
	addPlugin("http", loadHTTP)
}

func loadHTTP(Config ConfigurationInterface) syncFunc {

	resp, err := httpNewRequest("GET", fmt.Sprintf("Group=%s&Item=%s", Config.GetParent().Group, Config.GetParent().Item), nil)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(bytes))
	err = json.Unmarshal(bytes, &Config)
	if err != nil {
		log.Fatalln(err)
	}

	return syncHTTP
}

func syncHTTP(Config ConfigurationInterface) error {

	data, err := json.Marshal(Config)
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := httpNewRequest("PUT", fmt.Sprintf("Group=%s&Item=%s", Config.GetParent().Group, Config.GetParent().Item), bytes.NewReader(data))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(bytes))

	return nil
}

// credit: https://stackoverflow.com/questions/51452148/how-can-i-make-a-request-with-a-bearer-token-in-go
func httpNewRequest(method string, query string, body io.Reader) (*http.Response, error) {

	var bearer = "Bearer " + *httpToken
	var uri = *httpSource

	pq, err := url.ParseQuery(query)
	if err != nil {
		return nil, err
	}
	parsed, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	q := parsed.Query()
	for k, v := range pq {
		q.Add(k, v[0])
	}
	parsed.RawQuery = q.Encode()

	// Create a new request using http
	req, err := http.NewRequest(method, parsed.String(), body)
	if err != nil {
		return nil, err
	}

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)

	// Accept only JSON
	req.Header.Add("Accept", "application/json")

	if body != nil {
		// Content-Type is always json
		req.Header.Add("Content-Type", "application/json")
	}

	client := &http.Client{}
	r, err := client.Do(req)
	if r.StatusCode != 200 {
		return r, fmt.Errorf("Request %s not successful: %s", r.Request.URL, r.Status)
	}
	return r, nil
}
