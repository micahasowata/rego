package helper

import "path/filepath"

func CreatePaths(base, fileCategory, fileName string) (string, string) {
	return filepath.Join(base, fileCategory, fileName), filepath.Join(base, fileCategory)
}
