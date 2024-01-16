package organiser_test

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/mitchellh/go-homedir"
	"github.com/spobly/rego/organiser"
	"github.com/stretchr/testify/require"
)

func setUp(source, stat string) {
	// create a source directory
	err := os.Mkdir(source, fs.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	// create a sample file in that directory
	file := filepath.Join(source, "sample.mp3")

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

func tearDown(source, stat string) {
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

func TestOrganiserRunWithLocalPermission(t *testing.T) {
	setUp("source", "stat")
	defer tearDown("source", "stat")

	o, err := organiser.New(".", false)
	require.NoError(t, err)

	o.Path = filepath.Join(o.Path, "source")
	require.DirExists(t, o.Path)

	err = o.Run()
	require.NoError(t, err)
}

func TestOrganiserRunWithGlobalPermission(t *testing.T) {
	setUp("source", "stat")
	defer tearDown("source", "stat")

	o, err := organiser.New(".", true)
	require.NoError(t, err)

	hd, err := homedir.Dir()
	require.NoError(t, err)

	o.Path = filepath.Join(o.Path, "source")
	require.DirExists(t, o.Path)

	err = o.Run()
	require.NoError(t, err)

	require.DirExists(t, filepath.Join(hd, "Audio"))

	err = os.RemoveAll(filepath.Join(hd, "Audio"))
	require.NoError(t, err)
}

func TestOrganiserRunWithFilePath(t *testing.T) {
	setUp("source", "stat")
	defer tearDown("source", "stat")

	o, err := organiser.New(".", false)
	require.NoError(t, err)

	o.Path = filepath.Join(o.Path, "source", "sample.mp3")
	err = o.Run()

	require.Error(t, err)
	require.Contains(t, err.Error(), "not a directory")
}

func TestInvalidFilePath(t *testing.T) {
	o, err := organiser.New(".", false)
	require.NoError(t, err)

	o.Path = filepath.Join(o.Path, "source")

	err = o.Run()
	require.Error(t, err)
	require.Contains(t, err.Error(), "no such file or directory")
}
