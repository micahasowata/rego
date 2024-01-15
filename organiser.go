package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/spobly/ft"
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

	return &Organiser{
		Path:      path,
		UseGlobal: useGlobal,
	}
}

func moveFile(source, destination string) (err error) {
	src, err := os.Open(source)
	if err != nil {
		return err
	}
	defer src.Close()
	fi, err := src.Stat()
	if err != nil {
		return err
	}
	flag := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	perm := fi.Mode() & os.ModePerm
	dst, err := os.OpenFile(destination, flag, perm)
	if err != nil {
		return err
	}
	defer dst.Close()
	_, err = io.Copy(dst, src)
	if err != nil {
		dst.Close()
		os.Remove(destination)
		return err
	}
	err = dst.Close()
	if err != nil {
		return err
	}
	err = src.Close()
	if err != nil {
		return err
	}
	err = os.Remove(source)
	if err != nil {
		return err
	}
	return nil
}

// Run is the main worker function for organiser
func (o *Organiser) Run() {
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

		for _, file := range files {
			if !file.IsDir() && !o.UseGlobal {
				fmt.Println(file.Name())
				fileCategory, err := ft.GetFileCategory(filepath.Ext(file.Name()))
				if err != nil {
					log.Fatal(err)
				}

				sourcePath := filepath.Join(o.Path, file.Name())
				destPath := filepath.Join(o.Path, fileCategory, file.Name())
				moveFile(sourcePath, destPath)
			}
		}

	}

	// Based on UseGlobal move files to either root based paths or create directories in the CWD
	// ("jpg, jpeg, png" -> "Images", ".txt, .pdf" -> "Documents")

}
