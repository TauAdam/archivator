package cmd

import (
	"errors"
	"github.com/TauAdam/archivator/lib/compress"
	"github.com/TauAdam/archivator/lib/compress/vlc"
	"github.com/TauAdam/archivator/lib/compress/vlc/table/shannon_fano"
	"github.com/spf13/cobra"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var unpackCmd = &cobra.Command{
	Use:   "unpack",
	Short: "Unpack files",
	Run:   unpack,
}

func init() {
	rootCmd.AddCommand(unpackCmd)
	unpackCmd.Flags().StringP("algorithm", "a", "", "decompression algorithm")
	if err := unpackCmd.MarkFlagRequired("algorithm"); err != nil {
		panic(err)
	}
}

func unpack(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		handleError(errors.New("no file provided"))
	}

	var decoder compress.Decoder
	algo := cmd.Flag("algorithm").Value.String()
	switch algo {
	case "vlc":
		decoder = vlc.New(shannon_fano.Generator{})
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
	packed := decoder.Decode(data)

	err = os.WriteFile(unpackFileName(pathToFile), []byte(packed), 0666)
	if err != nil {
		handleError(err)
	}
}

func unpackFileName(pathToFile string) string {
	base := filepath.Base(pathToFile)
	extension := filepath.Ext(pathToFile)
	return strings.TrimSuffix(base, extension) + ".txt"
}
