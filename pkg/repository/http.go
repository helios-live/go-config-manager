package repository

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"go.ideatocode.tech/config/pkg/interfaces"
)

// HTTP .
type HTTP struct {
	Token string
	URL   string
}

type httpResponse struct {
	Status string
}

// Load .
func (h HTTP) Load(cfg interfaces.Manager, data interface{}) error {

	resp, err := h.httpNewRequest("GET", h.query(cfg.Path()), nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = cfg.Marshaler().Unmarshal(bytes, &data)
	if err != nil {
		return err
	}

	return nil
}

// Erase .
func (h HTTP) Erase(cfg interfaces.Manager) error {
	resp, err := h.httpNewRequest("DELETE", h.query(cfg.Path()), nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {

		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		respStatus := &httpResponse{}
		err = json.Unmarshal(bytes, &respStatus)
		if err != nil {
			return err
		}

		return errors.New(respStatus.Status)
	}
	return nil
}

// Save .
func (h HTTP) Save(cfg interfaces.Manager, data interface{}) error {

	buf, err := cfg.Marshaler().Marshal(data)
	if err != nil {
		return err
	}

	resp, err := h.httpNewRequest("PUT", h.query(cfg.Path()), bytes.NewReader(buf))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {

		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		respStatus := &httpResponse{}
		err = json.Unmarshal(bytes, &respStatus)
		if err != nil {
			return err
		}

		return errors.New(respStatus.Status)
	}
	return nil
}

func (h HTTP) query(path string) string {
	return fmt.Sprintf("path=%s", path)
}

// credit: https://stackoverflow.com/questions/51452148/how-can-i-make-a-request-with-a-bearer-token-in-go
func (h HTTP) httpNewRequest(method string, query string, body io.Reader) (*http.Response, error) {

	var bearer = "Bearer " + h.Token

	pq, err := url.ParseQuery(query)
	if err != nil {
		return nil, err
	}
	parsed, err := url.Parse(h.URL)
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
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return r, fmt.Errorf("Request %s not successful: %s", r.Request.URL, r.Status)
	}
	return r, nil
}
