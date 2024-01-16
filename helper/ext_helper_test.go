package helper_test

import (
	"testing"

	"github.com/spobly/rego/helper"
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
