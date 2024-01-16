package organiser

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/spobly/rego/helper"
)

type Organiser struct {
	Path      string
	UseGlobal bool
}

// New is the constructor function for the organiser struct
func New(path string, useGlobal bool) (*Organiser, error) {
	if path == "." {
		p, err := os.Getwd()
		if err != nil {
			return nil, err
		}
		path = p
	} else {
		path = strings.TrimPrefix(path, "/")

		hd, err := homedir.Dir()
		if err != nil {
			return nil, err
		}

		if !strings.HasPrefix(path, hd) {
			path = filepath.Join(hd, path)
		}
	}

	o := &Organiser{
		Path:      path,
		UseGlobal: useGlobal,
	}

	return o, nil
}

// Run is the main worker function for organiser
func (o *Organiser) Run() error {
	info, err := os.Stat(o.Path)
	if err != nil {
		e, _ := err.(*os.PathError)
		return e.Err
	}

	if info.IsDir() {
		// Check the file extensions
		files, err := os.ReadDir(o.Path)
		if err != nil {
			return err
		}

		for _, file := range files {
			if !file.IsDir() {
				fileCategory, err := helper.GetFileCategory(filepath.Ext(file.Name()))
				if err != nil {
					return err
				}

				sourcePath := filepath.Join(o.Path, file.Name())
				destPath := ""

				if !o.UseGlobal {
					destPath = filepath.Join(o.Path, fileCategory, file.Name())
					_, err = os.Stat(filepath.Join(o.Path, fileCategory))
					if err != nil {
						if os.IsNotExist(err) {
							err = os.Mkdir(filepath.Join(o.Path, fileCategory), fs.ModePerm)
							if err != nil {
								e, _ := err.(*os.PathError)
								return e.Err
							}
						}
					}
				} else {
					hd, err := homedir.Dir()
					if err != nil {
						return err
					}
					destPath = filepath.Join(hd, fileCategory, file.Name())
					_, err = os.Stat(filepath.Join(hd, fileCategory))
					if err != nil {
						if os.IsNotExist(err) {
							err = os.Mkdir(filepath.Join(hd, fileCategory), fs.ModePerm)
							if err != nil {
								e, _ := err.(*os.PathError)
								return e.Err
							}
						}

						e, _ := err.(*os.PathError)
						return e.Err
					}
				}

				err = helper.MoveFile(sourcePath, destPath)
				if err != nil {
					e, _ := err.(*os.PathError)
					return e.Err
				}
			}
		}

	}
	return nil
}
