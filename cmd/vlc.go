package cmd

import (
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
	pathToFile := args[0]

	f, err := os.Open(pathToFile)
	if err != nil {
		handleError(err)
	}
	data, err := io.ReadAll(f)
	if err != nil {
		handleError(err)
	}
	// TODO: Implement variable length encoding
	packed := string(data)

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
