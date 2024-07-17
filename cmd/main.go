package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Short: "Plain archivator",
}

func Execute() error {
	if err:= rootCmd.Execute(){
	 fmt.Errorf("Error: %v", err)
	 os.Exit(1)
	}
}

func main(){
	// pack
	Execute()
}