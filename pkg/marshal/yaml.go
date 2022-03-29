package marshal

import (
	"gopkg.in/yaml.v2"
)

// YAML .
type YAML struct{}

// Marshal .
func (j YAML) Marshal(v interface{}) (data []byte, err error) {
	buf, err := yaml.Marshal(v)
	return buf, err
}

// Unmarshal .
func (j YAML) Unmarshal(data []byte, v interface{}) error {
	return yaml.Unmarshal(data, v)

}
