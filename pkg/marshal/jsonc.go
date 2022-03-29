package marshal

import (
	"encoding/json"

	"github.com/muhammadmuzzammil1998/jsonc"
)

// JSONC .
type JSONC struct{}

// Marshal .
func (j JSONC) Marshal(v interface{}) (data []byte, err error) {
	return json.MarshalIndent(v, "", "\t")
}

// Unmarshal .
func (j JSONC) Unmarshal(data []byte, v interface{}) error {
	data = jsonc.ToJSON(data)
	return json.Unmarshal(data, v)

}
