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

	err1 := errors.New("1")
	err2 := errors.New("2")
	errs.Add(err1)
	errs.Add(err2)
	assert.Equal(t, "1\n2", errs.Error())

	errs.Add(err2)
	assert.Equal(t, "1\n2\n2", errs.Error())
	assert.True(t, errs.NotNil())
	assert.EqualValues(t, []error{err1, err2, err2}, errs.Errors())
}
