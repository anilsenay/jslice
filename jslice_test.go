package jslice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_New(t *testing.T) {
	newIntSlice := New[int]()
	assert.NotEmpty(t, newIntSlice)
	assert.Len(t, newIntSlice.Data, 0)
}

func Test_Of(t *testing.T) {
	newIntSlice := Of(1, 2, 3)
	assert.NotEmpty(t, newIntSlice)
	assert.Len(t, newIntSlice.Data, 3)
}

func Test_From(t *testing.T) {
	oldSlice := []int{1,2,3}
	
	newIntSlice := From(oldSlice)
	assert.NotEmpty(t, newIntSlice)
	assert.Len(t, newIntSlice.Data, 3)
}

func Test_Length(t *testing.T) {
	newIntSlice := Of(1, 2, 3)
	assert.NotEmpty(t, newIntSlice)
	assert.Len(t, newIntSlice.Data, 3)
	assert.Equal(t, 3, newIntSlice.Length())
}

func Test_String(t *testing.T) {
	newIntSlice := Of(1, 2, 3)
	assert.NotEmpty(t, newIntSlice)
	assert.Len(t, newIntSlice.Data, 3)
	assert.Equal(t, "[1,2,3]", newIntSlice.String())
}