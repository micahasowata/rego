package helper_test

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/spobly/rego/helper"
	"github.com/stretchr/testify/require"
)

var (
	source   = "source"
	file     string
	filename = "sample.txt"
	stat     = "stat"
)

func setUp() {
	// create a source directory
	err := os.Mkdir(source, fs.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	// create a sample file in that directory
	file = filepath.Join(source, filename)

	err = os.WriteFile(file, []byte{1, 2}, fs.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	// create a stat directory
	err = os.Mkdir("stat", fs.ModePerm)
	if err != nil {
		log.Fatal("stat", err)
	}
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
	// set up
	setUp()
	defer tearDown()

	// act
	err := helper.MoveFile(file, filepath.Join(stat, filename))
	require.NoError(t, err)

	// assert
	require.NoError(t, err)
	require.DirExists(t, source)
	require.DirExists(t, stat)
	require.NoFileExists(t, file)
	require.FileExists(t, filepath.Join(stat, filename))
}
