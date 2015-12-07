package packagecloud

import (
	"encoding/json"
	"errors"
	"fmt"
)

// ErrInvalidKey is an error for when the key provided in the config is
// invalid.
var ErrInvalidKey = errors.New("invalid key, must be non-zero in length")

// License is a struct which contains the unmarshaled data from a License request.
type License struct {
	License   string `json:"license"`
	Signature string `json:"signature"`
}

// Licenses is a function to retreive a packagecloud:enterprise license file
// and GPG signature. Signature can be verified with the packagecloud
// GPG key: https://packagecloud.io/gpg.key.
func (c *Client) Licenses(key string) (*License, error) {
	if len(key) == 0 {
		return nil, ErrInvalidKey
	}

	path := fmt.Sprintf("licenses/%s/licenses.json", key)

	b, err := c.request("GET", path)

	if err != nil {
		return nil, err
	}

	l := &License{}

	err = json.Unmarshal(b, l)

	return l, err
}
