package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mitchellh/go-homedir"
)

type Organiser struct {
	Path      string
	UseGlobal bool
}

// NewOrganiser is the constructor function for the organiser struct
func NewOrganiser(path string, useGlobal bool) *Organiser {
	if path == "." {
		p, err := os.Getwd()
		if err != nil {
			log.Fatal(err.Error())
		}
		path = p
	} else {
		path = strings.TrimPrefix(path, "/")

		hd, err := homedir.Dir()
		if err != nil {
			log.Fatal(err.Error())
		}

		if !strings.HasPrefix(path, hd) {
			path = fmt.Sprintf("%s/%s", hd, path)
		}
	}

	if !strings.HasSuffix(path, "/") {
		path = fmt.Sprintf("%s/", path)
	}

	return &Organiser{
		Path:      path,
		UseGlobal: useGlobal,
	}
}

// Run is the main worker function for organiser
func (o *Organiser) Run() {
	// Confirm if the path is a valid directory
	info, err := os.Stat(o.Path)
	if err != nil {
		log.Fatal(err)
	}

	if info.IsDir() {
		// Check the file extensions
		files, err := os.ReadDir(o.Path)
		if err != nil {
			log.Fatal(err)
		}

		fileExtensions := map[string][]string{
			"Audio":       {".mp3", ".wav", ".aac", ".flac", ".m4a"},
			"Video":       {".mp4", ".avi", ".mkv", ".mov", ".wmv"},
			"Images":      {".jpg", ".jpeg", ".png", ".gif", ".bmp", ".tiff", ".tif"},
			"Documents":   {".pdf", ".docx", ".xlsx", ".pptx", ".txt"},
			"Archives":    {".zip", ".rar", ".7z", ".tar", ".gz"},
			"Executables": {".exe", ".app", ".dmg", ".bin", ".msi"}}

		for _, file := range files {
			if !file.IsDir() && !o.UseGlobal {

			}
		}

	}

	// Based on UseGlobal move files to either root based paths or create directories in the CWD
	// ("jpg, jpeg, png" -> "Images", ".txt, .pdf" -> "Documents")

}
