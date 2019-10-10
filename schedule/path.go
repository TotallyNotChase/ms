package schedule

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/adrg/xdg"
)

const (
	app      = "ms"
	savefile = "current.json"
)

var (
	dataDir  string
	dataPath string
)

// SetDataPath tries to fetch the current OS' data directory instead
// of hardcoding it
func SetDataPath() error {
	var (
		dir string
		err error
	)

	switch runtime.GOOS {
	case "linux":
		dir = xdg.DataHome

	default: // windows, osx, etc
		dir, err = os.UserConfigDir()
		if err != nil {
			return fmt.Errorf("couldn't find the config dir: %v", err)
		}
	}

	dataDir = filepath.Join(dir, app)
	dataPath = filepath.Join(dataDir, savefile)

	return nil
}
