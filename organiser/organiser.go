package organiser

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/spobly/rego/helper"
)

// Organiser stores information about the directory to be managed and user permissions
type Organiser struct {
	Path      string
	UseGlobal bool
}

// New returns a new Organiser with the provided path and permission or an error
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

// Run re-organises the data directory provided according to user permission
func (o *Organiser) Run() error {

	info, err := os.Stat(o.Path)
	if err != nil {
		e, _ := err.(*os.PathError)
		return e.Err
	}

	if info.IsDir() {

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

				if !o.UseGlobal {

					destPath, statPath := helper.CreatePaths(o.Path, fileCategory, file.Name())

					_, err = os.Stat(statPath)
					if err != nil {
						if os.IsNotExist(err) {
							err = os.Mkdir(statPath, fs.ModePerm)
							if err != nil {
								e, _ := err.(*os.PathError)
								return e.Err
							}
						}
					}

					err = helper.MoveFile(sourcePath, destPath)
					if err != nil {
						e, _ := err.(*os.PathError)
						return e.Err
					}

				} else {

					hd, err := homedir.Dir()
					if err != nil {
						return err
					}

					destPath, statPath := helper.CreatePaths(hd, fileCategory, file.Name())

					_, err = os.Stat(statPath)
					if err != nil {
						if os.IsNotExist(err) {
							err = os.Mkdir(statPath, fs.ModePerm)
							if err != nil {
								e, _ := err.(*os.PathError)
								return e.Err
							}
						}
						e, _ := err.(*os.PathError)
						return e.Err
					}

					err = helper.MoveFile(sourcePath, destPath)
					if err != nil {
						e, _ := err.(*os.PathError)
						return e.Err
					}

				}
			}
		}

	}
	return nil
}
