package helper_test

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"testing"
)

var (
	source string
	file   string
	stat   string
)

func setUp() {
	// create a source directory
	source, err := os.MkdirTemp("", "source")
	if err != nil {
		log.Fatal(err)
	}

	// create a sample file in that directory
	file = filepath.Join(source, "sample.txt")
	err = os.WriteFile(file, []byte{1, 2}, fs.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	// create a stat directory
	stat, err = os.MkdirTemp("", "stat")
	log.Fatal(err)
}

func tearDown() {
	// delete source directory
	err := os.RemoveAll(source)
	if err != nil {
		log.Fatal(err)
	}

	// delete stat directory
	err = os.RemoveAll(stat)
	if err != nil {
		log.Fatal(err)
	}
}

func TestMoveFile(t *testing.T) {
	setUp()
	defer tearDown()
	// arrange
	// act
	// assert
}
