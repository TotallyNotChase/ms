package main

import (
	"fmt"
	"os"

	"github.com/shkh/lastfm-go/lastfm"
	"gitlab.com/Sacules/jsonfile"
)

// APIKey hold the public and secret keys for accessing Last.FM's API.
type APIKey struct {
	Key    string `json:"key"`
	Secret string `json:"secret"`
}

// InitAPI tries to read the API keys from disk and returns a pointer to
// Last.FM's API. If the keys are missing, it will exit and print a warning.
func InitAPI() *lastfm.Api {
	apiKey := &APIKey{}
	err := apiKey.Load()
	if err != nil {
		fmt.Println("API Key couldn't be found, creating a new apikey.json file.")
		fmt.Println("Please put your keys there before using this program.")
		apiKey.Save()
		os.Exit(1)
	}

	return lastfm.New(apiKey.Key, apiKey.Secret)
}

// Load the keys from a JSON file.
func (a *APIKey) Load() error {
	return jsonfile.LoadFile(a, "apikey.json")
}

// Save the keys to a JSON file.
func (a *APIKey) Save() error {
	return jsonfile.SaveFile(a, "apikey.json")
}
