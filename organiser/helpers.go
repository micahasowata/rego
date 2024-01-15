package organiser

import (
	"io"
	"log"
	"os"
)

func moveFile(sourcePath, destPath string) error {
	inputFile, err := os.Open(sourcePath)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	outputFile, err := os.Create(destPath)
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	_, err = io.Copy(outputFile, inputFile)
	if err != nil {
		log.Fatal(err)
	}

	err = os.Remove(sourcePath)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
