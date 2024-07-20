package cmd

import (
	"errors"
	"github.com/TauAdam/archivator/lib/compress"
	"github.com/TauAdam/archivator/lib/compress/vlc"
	"github.com/spf13/cobra"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var packCmd = &cobra.Command{
	Use:   "pack",
	Short: "Pack files",
	Run:   pack,
}

func init() {
	rootCmd.AddCommand(packCmd)
	packCmd.Flags().StringP("algorithm", "a", "", "compression algorithm")
	if err := packCmd.MarkFlagRequired("algorithm"); err != nil {
		panic(err)
	}
}

const (
	packedFileExtension = ".vlc"
)

func pack(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		handleError(errors.New("no file provided"))
	}
	var encoder compress.Encoder
	algo := cmd.Flag("algorithm").Value.String()
	switch algo {
	case "vlc":
		encoder = vlc.New()
	default:
		cmd.PrintErr("unsupported algorithm")
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
	packed := encoder.Encode(string(data))

	err = os.WriteFile(packFileName(pathToFile), packed, 0666)
	if err != nil {
		handleError(err)
	}
}

func packFileName(pathToFile string) string {
	base := filepath.Base(pathToFile)
	extension := filepath.Ext(pathToFile)
	return strings.TrimSuffix(base, extension) + packedFileExtension
}
