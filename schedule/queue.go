package schedule

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/adrg/xdg"
	"gitlab.com/Sacules/jsonfile"
)

const (
	app = "/ms"
)

// Queue is the overall scheduler, consisting of 4 blocks.
type Queue [4]*Block

// Add a block to the queue and get rid of any old ones.
func (q *Queue) Add(b *Block) {
	for i := len(q) - 1; i > 0; i-- {
		q[i] = q[i-1]
	}

	q[0] = b
}

// Load queue from disk.
func (q *Queue) Load() error {
	err := createDirNotExist(xdg.ConfigHome + app)
	if err != nil {
		return err
	}

	return jsonfile.Load(q, xdg.ConfigHome+app+"/current.json")
}

// Replace goes through the queue and replaces an album for a new one.
func (q *Queue) Replace(old, actual Album) {
	for _, block := range q {
		block.Replace(old, actual)
	}
}

// Save queue to disk.
func (q *Queue) Save() error {
	err := createDirNotExist(xdg.ConfigHome + app)
	if err != nil {
		return err
	}

	return jsonfile.Save(q, xdg.ConfigHome+app+"/current.json")
}

// ShowCurrent prints the current week of records in the queue.
func (q *Queue) ShowCurrent() {
	var (
		w       = tabwriter.NewWriter(os.Stdout, 0, 8, 2, '\t', 0)
		blankln = "\t\t\n"
		headers = []string{
			"Block",
			"Listened",
			"Rated",
		}
		underlined    = make([]string, len(headers))
		headerstabbed = strings.Join(headers, "\t")
	)
	defer w.Flush()

	// Calculate the right amount of underlines
	for i, h := range headers {
		underlined[i] = strings.Repeat("-", len(h))
	}
	underlinedtabbed := strings.Join(underlined, "\t")

	fmt.Fprintln(w, headerstabbed+"\n"+underlinedtabbed)
	for i, block := range q {
		if block != nil && i != 2 {
			fmt.Fprintf(w, "%s\t\t\n", block.Name)
			for _, album := range block.Albums {
				var (
					listened = "❌"
					rated    = " "
				)

				switch i {
				case 0:
					if album.FirstListen {
						listened = "✓"
					}

				case 1:
					if album.SecondListen {
						listened = "✓"
					}

				case 3:
					if album.ThirdListen {
						listened = "✓"
					}
				}

				if album.Rated {
					rated = "✓"
				}

				fmt.Fprintf(w, "• %s\t%s\t%s\n", album.Name, listened, rated)
			}
			fmt.Fprintf(w, blankln)
		}
	}
}
