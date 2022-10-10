package json

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshal(t *testing.T) {
	type User struct {
		Username string `json:"username"`
		Age      int    `json:"age"`
	}
	user := &User{
		Username: "foo",
		Age:      2,
	}

	data, err := json.Marshal(user)
	assert.NoError(t, err)
	u, err := Unmarshal[User](data)
	assert.NoError(t, err)
	assert.EqualValues(t, user, u)

	u, err = Unmarshal[User]([]byte("{u}"))
	assert.Error(t, err)
	assert.Nil(t, u)
}

func TestUnmarshalString(t *testing.T) {
	type User struct {
		Username string `json:"username"`
		Age      int    `json:"age"`
	}
	user := &User{
		Username: "foo",
		Age:      2,
	}

	data, err := json.Marshal(user)
	assert.NoError(t, err)
	u, err := UnmarshalString[User](string(data))
	assert.NoError(t, err)
	assert.EqualValues(t, user, u)
}
