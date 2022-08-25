package errs

import (
	"errors"
	"io/fs"
	"os"
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

	assert.True(t, errs.Is(err1))
	assert.True(t, errs.Is(err2))

	_, err := os.Open("non-existing")
	errs.Add(err)

	var pathError *fs.PathError
	assert.True(t, errs.As(&pathError))
	assert.EqualValues(t, "non-existing", pathError.Path)
}
