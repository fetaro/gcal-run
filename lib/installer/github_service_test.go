package installer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGithubService_GetLatestVersion(t *testing.T) {
	g := NewGithubService()
	version, err := g.GetLatestVersion()
	assert.Nil(t, err)
	assert.NotNil(t, version)
}
