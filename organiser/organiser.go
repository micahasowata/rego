package organiser

import (
	"errors"
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
	// Get information about the provided path
	info, err := os.Stat(o.Path)
	if err != nil {
		e, _ := err.(*os.PathError)
		return e.Err
	}

	// Confirm that the path is directory
	if info.IsDir() {

		// Read the directory to get all the files that are in it. Just files, no sub-directories
		files, err := os.ReadDir(o.Path)
		if err != nil {
			return err
		}
		// Iterate through the files to carry out acions on individual files
		for _, file := range files {

			// Ensure that it is files only. No directories.
			if !file.IsDir() {
				// Get the file category based on its extension
				fileCategory, err := helper.GetFileCategory(filepath.Ext(file.Name()))
				if err != nil {
					return err
				}

				// Construct a source path which is the current location of the file
				sourcePath := filepath.Join(o.Path, file.Name())

				// Ensure that the next steps are within the confines of the user's permission
				if !o.UseGlobal {
					// Generate a destination path (directories plus file name) and the stat path (directories only. no file name.)
					destPath, statPath := helper.CreatePaths(o.Path, fileCategory, file.Name())

					// Check the stat path
					_, err = os.Stat(statPath)
					if err != nil {
						if os.IsNotExist(err) {
							// Create a new directory at the stat path if one does not exist
							err = os.Mkdir(statPath, fs.ModePerm)
							if err != nil {
								e, _ := err.(*os.PathError)
								return e.Err
							}
						}
					}

					// Move files from the source path to the destination path
					err = helper.MoveFile(sourcePath, destPath)
					if err != nil {
						e, _ := err.(*os.PathError)
						return e.Err
					}

				} else { // If permitted to move files to the home directory

					// Get the home directory
					hd, err := homedir.Dir()
					if err != nil {
						return err
					}

					// Construct the destination and stat paths
					destPath, statPath := helper.CreatePaths(hd, fileCategory, file.Name())

					// Verify the stat path
					_, err = os.Stat(statPath)
					if err != nil {
						if os.IsNotExist(err) {
							// Create a directory in that path if it doesn't exist
							err = os.Mkdir(statPath, fs.ModePerm)
							if err != nil {
								e, _ := err.(*os.PathError)
								return e.Err
							}
						}
					}

					// Move files from the source path to the destination path
					err = helper.MoveFile(sourcePath, destPath)
					if err != nil {
						e, _ := err.(*os.PathError)
						return e.Err
					}

				}
			}
		}

	} else {
		// return an error if the provided path leads to a file
		return errors.New(o.Path + " is not a directory")
	}
	return nil
}
