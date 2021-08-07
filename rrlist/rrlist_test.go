package rrlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRRList(t *testing.T) {
	rrl := New()
	assert.Equal(t, 0, rrl.length)
	rrl.Add("172.54.23.5:8000")
	assert.Equal(t, 1, rrl.length)
	addr, err := rrl.Head()
	assert.NoError(t, err)
	assert.Equal(t, "172.54.23.5:8000", addr)
	rrl.Add("172.54.33.5:80")
	assert.Equal(t, 2, rrl.length)
	addr, err = rrl.Reveal()
	assert.NoError(t, err)
	assert.Equal(t, "172.54.23.5:8000", addr)
	addr, err = rrl.Head()
	assert.NoError(t, err)
	assert.Equal(t, "172.54.33.5:80", addr)
}
