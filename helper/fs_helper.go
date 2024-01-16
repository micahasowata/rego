package helper

import (
	"io"
	"os"
)

func MoveFile(sourcePath, destPath string) error {
	inputFile, err := os.Open(sourcePath)
	if err != nil {
		e, _ := err.(*os.PathError)
		return e.Err
	}
	defer inputFile.Close()

	outputFile, err := os.Create(destPath)
	if err != nil {
		e, _ := err.(*os.PathError)
		return e.Err
	}
	defer outputFile.Close()

	_, err = io.Copy(outputFile, inputFile)
	if err != nil {
		return err
	}

	err = os.Remove(sourcePath)
	if err != nil {
		e, _ := err.(*os.PathError)
		return e.Err
	}
	return nil
}
