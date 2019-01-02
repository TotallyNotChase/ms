package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/shkh/lastfm-go/lastfm"
)

const (
	timeout = time.Duration(5 * time.Second)
)

type downData struct {
	url  string
	name string
	dir  string
}

// GetArtwork concurrently downloads the cover for the given album on the given
// directory. The given results channel sends a success
// mesage if an image was downloaded correctly, and any errors are sent to the
// errs channel. A WaitGroup is used to sync all goroutines.
func GetArtwork(api *lastfm.Api, album string, dir string, results chan string, errs chan error, wg *sync.WaitGroup) {
	al, err := api.Album.Search(lastfm.P{
		"album": album,
	})
	if err != nil {
		errs <- fmt.Errorf("Error while looking up %s: %s", album, err)
		return
	}

	err = createDirNotExist(dir)
	if err != nil {
		errs <- fmt.Errorf("Error while checking directory %s: %s", dir, err)
		return
	}

	// Generate all necessary workers
	datach := make(chan downData)
	for i := 0; i < runtime.NumCPU()+1; i++ {
		go downImage(datach, results, errs, wg)
		wg.Add(1)
	}

	for _, match := range al.AlbumMatches {
		for _, image := range match.Images {
			// Only work with big images
			if image.Size == "large" && image.Url != "" {
				filename := strings.Replace(match.Name, "/", " ", -1)
				// Don't download if it's already cached!
				if _, err := os.Stat(dir + "/" + filename + ".png"); !os.IsNotExist(err) {
					break
				}

				data := downData{
					url:  image.Url,
					name: filename,
					dir:  dir,
				}

				datach <- data
			}
		}
	}

	close(datach)
}

func downImage(datach chan downData, results chan string, errs chan error, wg *sync.WaitGroup) {
	defer wg.Done()

	// Custom client with timeout
	client := http.Client{
		Timeout: timeout,
	}

	for data := range datach {
		response, err := client.Get(data.url)
		defer func() {
			err := response.Body.Close()
			if err != nil {
				errs <- fmt.Errorf("Error with closing the response: %s", err)
			}
		}()

		if response == nil {
			errs <- fmt.Errorf("Error: Null response for %s", data.url)
			return
		}
		if err != nil {
			errs <- fmt.Errorf("Error fetching URL %s: %s", data.url, err)
			return
		}

		fileName := data.dir + "/" + data.name + ".png"

		file, err := os.Create(fileName)
		if err != nil {
			errs <- fmt.Errorf("Error while creating file %s: %s", fileName, err)
			return
		}

		_, err = io.Copy(file, response.Body)
		if err != nil {
			errs <- fmt.Errorf("Error while writing image to disk: %s", err)
			return
		}

		results <- "Sucessfully downloaded " + data.name
	}
}

func createDirNotExist(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.MkdirAll(dir, 0755)
	}

	return nil
}
