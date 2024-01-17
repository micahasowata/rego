package organiser_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/mitchellh/go-homedir"
	"github.com/spobly/rego/internal/organiser"
	"github.com/stretchr/testify/require"
)

func TestNewWithDot(t *testing.T) {
	// arrange
	p := &organiser.Organiser{}
	wd, err := os.Getwd()
	require.NoError(t, err)

	// act
	o, err := organiser.New(".", false)

	// assert
	require.NoError(t, err)
	require.IsType(t, o, p)
	require.Equal(t, o.Path, wd)
	require.Equal(t, o.UseGlobal, false)
}

func TestNewWithAPath(t *testing.T) {
	p := "Documents"
	hd, err := homedir.Dir()
	require.NoError(t, err)

	o, err := organiser.New(p, false)

	require.NoError(t, err)
	require.Contains(t, o.Path, hd)
	require.Contains(t, o.Path, p)
	require.Equal(t, o.Path, filepath.Join(hd, p))
}
