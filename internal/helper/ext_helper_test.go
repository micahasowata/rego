package helper_test

import (
	"testing"

	"github.com/spobly/rego/internal/helper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidExtension(t *testing.T) {
	ext := ".png"
	want := "Images"

	got, err := helper.GetFileCategory(ext)

	require.NoError(t, err)
	assert.Equal(t, want, got, "file category should be Images")
}

func TestExtensionMissingDot(t *testing.T) {
	ext := "dmg"
	want := "Executables"

	got, err := helper.GetFileCategory(ext)

	require.NoError(t, err)
	assert.Equal(t, want, got, "add . and match file type")
}

func TestMissingExtension(t *testing.T) {
	ext := ".img"
	want := "not found"

	got, err := helper.GetFileCategory(ext)

	require.Error(t, err)
	require.Contains(t, err.Error(), want)
	assert.Empty(t, got)
}
