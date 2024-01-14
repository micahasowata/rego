package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Organiser struct {
	Path      string
	UseGlobal bool
}

// NewOrganiser is the constructor function for the organiser struct
func NewOrganiser(path string) *Organiser {
	if !(strings.HasPrefix(path, "/") && strings.HasSuffix(path, "/")) {
		path = "/" + path + "/"
	}

	if path == "." {
		p, err := os.Getwd()
		if err != nil {
			log.Fatal(err.Error())
		}
		path = p
	}

	return &Organiser{
		Path: path,
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
		fmt.Println("Hey it's a directory")
	}
	// Check the read/write permissions on each file in the directory
	// Check the file extensions
	// Based on UseGlobal move files to either root based paths or create directories in the CWD
	// ("jpg, jpeg, png" -> "Images", ".txt, .pdf" -> "Documents")
}
