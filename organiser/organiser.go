package organiser

import (
	"io/fs"
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

// New is the constructor function for the organiser struct
func New(path string, useGlobal bool) *Organiser {
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
			path = filepath.Join(hd, path)
		}
	}

	return &Organiser{
		Path:      path,
		UseGlobal: useGlobal,
	}
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
			if !file.IsDir() {
				fileCategory, err := ft.GetFileCategory(filepath.Ext(file.Name()))
				if err != nil {
					log.Fatal(err)
				}

				sourcePath := ""
				destPath := ""

				if !o.UseGlobal {
					sourcePath = filepath.Join(o.Path, file.Name())
					destPath = filepath.Join(o.Path, fileCategory, file.Name())
					_, err = os.Stat(filepath.Join(o.Path, fileCategory))
					if err != nil {
						if os.IsNotExist(err) {
							err = os.Mkdir(filepath.Join(o.Path, fileCategory), fs.ModePerm)
							if err != nil {
								log.Fatal(err)
							}
						}
					}
				} else {
					hd, err := homedir.Dir()
					if err != nil {
						log.Fatal(err)
					}
					sourcePath = filepath.Join(o.Path, file.Name())
					destPath = filepath.Join(hd, fileCategory, file.Name())
					_, err = os.Stat(filepath.Join(hd, fileCategory))
					if err != nil {
						if os.IsNotExist(err) {
							err = os.Mkdir(filepath.Join(hd, fileCategory), fs.ModePerm)
							if err != nil {
								log.Fatal(err)
							}
						}
					}
				}

				moveFile(sourcePath, destPath)
			}
		}

	}

	// Based on UseGlobal move files to either root based paths or create directories in the CWD
	// ("jpg, jpeg, png" -> "Images", ".txt, .pdf" -> "Documents")

}
