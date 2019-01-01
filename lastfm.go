package main

import (
	"gitlab.com/Sacules/jsonfile"
)

// APIKey hold the public and secret keys for accessing Last.FM's API.
type APIKey struct {
	Key    string
	Secret string
}

// Load the keys from a JSON file.
func (a *APIKey) Load() error {
	return jsonfile.LoadFile(a, "apikey.json")
}

// Save the keys to a JSON file.
func (a *APIKey) Save() error {
	return jsonfile.SaveFile(a, "apikey.json")
}
