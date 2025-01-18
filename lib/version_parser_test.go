package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseVersionStr(t *testing.T) {
	verionStr := "v1.2.0"
	v, err := ParseVersionStr(verionStr)
	assert.Nil(t, err)
	assert.Equal(t, 1, v.major)
	assert.Equal(t, 2, v.minor)
	assert.Equal(t, 0, v.bugfix)
}

func TestVersion_IsNewer(t *testing.T) {
	v1 := &Version{1, 2, 3}
	v2 := &Version{1, 2, 4}
	assert.False(t, v1.IsNewer(v2))
	assert.True(t, v2.IsNewer(v1))
}

func TestVersion_IsNewer2(t *testing.T) {
	v1 := &Version{1, 2, 4}
	v2 := &Version{1, 3, 4}
	assert.False(t, v1.IsNewer(v2))
	assert.True(t, v2.IsNewer(v1))
}

func TestVersion_IsNewer3(t *testing.T) {
	v1 := &Version{1, 2, 4}
	v2 := &Version{2, 2, 4}
	assert.False(t, v1.IsNewer(v2))
	assert.True(t, v2.IsNewer(v1))
}
