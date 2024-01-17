package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spobly/rego/internal/organiser"
)

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

	useGlobal, err := cmd.Flags().GetBool("use-global")
	if err != nil {
		log.Println("use global permission couldn't be parsed")
		return
	}

	o, err := organiser.New(path, useGlobal)
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
