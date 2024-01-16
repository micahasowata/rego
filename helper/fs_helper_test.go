package helper_test

import (
	"io/fs"
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
	source, _ = os.MkdirTemp("", "source")
	// create a sample file in that directory
	file = filepath.Join(source, "sample.txt")
	_ = os.WriteFile(file, []byte{1, 2}, fs.ModePerm)

	// create a stat directory
	stat, _ = os.MkdirTemp("", "stat")
}

func tearDown() {
	// delete source directory
	// delete stat directory
}

func TestMoveFile(t *testing.T) {
	setUp()
	defer tearDown()
	// arrange
	// act
	// assert
}
