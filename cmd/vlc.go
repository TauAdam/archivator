package cmd

import (
	"errors"
	"github.com/TauAdam/archivator/lib/compress/vlc"
	"github.com/spf13/cobra"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var vlcCmd = &cobra.Command{
	Use:   "vlc",
	Short: "Variable length encoding",
	Run:   pack,
}

const (
	packedFileExtension = ".vlc"
)

func init() {
	packCmd.AddCommand(vlcCmd)
}

func pack(_ *cobra.Command, args []string) {
	if len(args) == 0 {
		handleError(errors.New("no file provided"))
	}

	pathToFile := args[0]

	f, err := os.Open(pathToFile)
	if err != nil {
		handleError(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			handleError(err)
		}
	}(f)
	data, err := io.ReadAll(f)
	if err != nil {
		handleError(err)
	}
	packed := vlc.Encode(string(data))

	err = os.WriteFile(packFileName(pathToFile), []byte(packed), 0666)
	if err != nil {
		handleError(err)
	}
}

func packFileName(pathToFile string) string {
	base := filepath.Base(pathToFile)
	extension := filepath.Ext(pathToFile)
	return strings.TrimSuffix(base, extension) + packedFileExtension
}
