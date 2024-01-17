package helper_test

import (
	"fmt"
	"testing"

	"github.com/spobly/rego/internal/helper"
	"github.com/stretchr/testify/assert"
)

func TestCreatePaths(t *testing.T) {
	base := "/x/y"
	fc := "Data"
	filename := "quotes.csv"

	dp, sp := helper.CreatePaths(base, fc, filename)

	assert.Equal(t, dp, fmt.Sprintf("%s/%s/%s", base, fc, filename))
	assert.Equal(t, sp, fmt.Sprintf("%s/%s", base, fc))
}
