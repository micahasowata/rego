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
	return &Organiser{
		Path: path,
	}
}

// Run is the main worker function for organiser
func (o *Organiser) Run() {
	// Confirm if the path is a valid directory

	if !(strings.HasPrefix(o.Path, "/") && strings.HasSuffix(o.Path, "/")) {
		o.Path = "/" + o.Path + "/"
	}

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
