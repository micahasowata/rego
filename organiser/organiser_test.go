package organiser_test

import (
	"os"
	"testing"

	"github.com/spobly/rego/organiser"
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
