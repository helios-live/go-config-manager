package config

import "net/url"

type RepositoryDefinition struct {
	url *url.URL
}

// String returns a string representation of the Repository Definition
func (rd *RepositoryDefinition) String() string {
	return rd.url.String()
}

// MarshalJSON turns a RepositoryDefinition into a json string
func (rd *RepositoryDefinition) MarshalJSON() ([]byte, error) {

	return []byte("\"" + rd.url.String() + "\""), nil
}

// UnmarshalJSON parses a json string into a RepositoryDefinition
func (rd *RepositoryDefinition) UnmarshalJSON(data []byte) error {
	return rd.UnmarshalText(string(data[1 : len(data)-1]))
}

// UnmarshalText parses a string into a RepositoryDefinition
func (rd *RepositoryDefinition) UnmarshalText(text string) error {
	uri, err := url.Parse(text)
	if err != nil {
		return err
	}

	if uri.Opaque != "" {
		uri.Path = uri.Opaque
		uri.Opaque = ""
	}
	if uri.Scheme == "" {
		uri.Scheme = "file"
	}

	rd.url = uri
	return nil
}
