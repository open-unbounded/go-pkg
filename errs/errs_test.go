package errs

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrs(t *testing.T) {
	errs := Errs{}
	errs.Add(nil)
	assert.False(t, errs.NotNil())
	assert.Equal(t, "", errs.Error())

	errs.Add(errors.New("1"))
	errs.Add(errors.New("2"))
	assert.Equal(t, "1\n2", errs.Error())

	errs.Add(errors.New("2"))
	assert.Equal(t, "1\n2\n2", errs.Error())

	assert.True(t, errs.NotNil())
}
