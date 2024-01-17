package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spobly/rego/internal/organiser"
)

func init() {
	rootCmd.AddCommand(organizeCmd)

	organizeCmd.Flags().StringP("path", "p", "", "path to the folder to be organised e.g Documents/Test or if you want to organise the current folder just put .")
	organizeCmd.Flags().BoolP("global", "g", false, "global determines if the new folders are created within the current folder or on the main home path")
}

var organizeCmd = &cobra.Command{
	Use:   "o",
	Short: "Organise the files into folders",
	Long: `Organise the files into folders based on their file type i.e extension e.g .jpg.
	It would not re-arrange sub-directories and it might return an error if it encounters an unknown 
	file type.
	
	Usage:
	
	rego o -p -g`,

	Run: organise,
}

func organise(cmd *cobra.Command, args []string) {
	path, err := cmd.Flags().GetString("path")
	if err != nil {
		log.Println("directory path couldn't be parsed")
		return
	}

	global, err := cmd.Flags().GetBool("global")
	if err != nil {
		log.Println("global permission couldn't be parsed")
		return
	}

	o, err := organiser.New(path, global)
	if err != nil {
		log.Println("organiser could not set up properly")
		return
	}

	err = o.Run()
	if err != nil {
		log.Printf("there was an error while re-arranging %s\n", o.Path)
		return
	}
}
