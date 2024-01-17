package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rego",
	Short: "Organize your files into easy to access folders",
	Long: `rego makes it fast and easy to organize all your files in to folders 
	that can be, if need be, quickly searched and makes your file system a lot neater.`,
	Version:       "0.1.0",
	SilenceErrors: true,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
