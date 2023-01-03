package jslice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Append(t *testing.T) {
	IntSlice := New[int]()
	assert.NotEmpty(t, IntSlice)
	assert.Len(t, IntSlice.Data, 0)

	IntSlice.Append(1)
	assert.Len(t, IntSlice.Data, 1)
	IntSlice.Append(1,2,3)
	assert.Len(t, IntSlice.Data, 4)
	IntSlice.Append()
	assert.Len(t, IntSlice.Data, 4)
}

func Test_Prepend(t *testing.T) {
	IntSlice := Of(1,2,3)
	assert.Len(t, IntSlice.Data, 3)

	IntSlice.Prepend(0)
	assert.Len(t, IntSlice.Data, 4)
	IntSlice.Prepend(-1)
	assert.Len(t, IntSlice.Data, 5)
	assert.Equal(t, -1, IntSlice.Data[0])
}

func Test_At(t *testing.T) {
	IntSlice := Of(1,2,3)
	assert.NotEmpty(t, IntSlice)
	assert.Len(t, IntSlice.Data, 3)

	assert.Equal(t, 1, IntSlice.At(0))
	assert.Equal(t, 2, IntSlice.At(1))

	// [5] == 5 % 3 == [2]
	assert.Equal(t, 3, IntSlice.At(5))

	// [-1] == [2], [-2] == [1], [-3] == [0], [-4] == [2], ...
	assert.Equal(t, 3, IntSlice.At(-1))
	assert.Equal(t, 2, IntSlice.At(-5))
}

func Test_Concat(t *testing.T) {
	IntSlice := Of(1,2,3)
	assert.NotEmpty(t, IntSlice)
	assert.Len(t, IntSlice.Data, 3)

	intArr := []int{4,5}
	newIntSlice := IntSlice.Concat(intArr)
	assert.Len(t, newIntSlice.Data, 5)
}

func Test_Every(t *testing.T) {
	IntSlice := Of(1,2,3)
	assert.NotEmpty(t, IntSlice)
	assert.Len(t, IntSlice.Data, 3)

	assert.True(t, IntSlice.Every(func(el int) bool { return el < 5}))
	assert.False(t, IntSlice.Every(func(el int) bool { return el > 1}))
}

func Test_Filter(t *testing.T) {
	IntSlice := Of(1,2,3)
	assert.NotEmpty(t, IntSlice)
	assert.Len(t, IntSlice.Data, 3)

	newIntSlice := IntSlice.Filter(func(el int) bool { return el > 1})
	assert.Len(t, newIntSlice.Data, 2)
}

func Test_Find(t *testing.T) {
	IntSlice := Of(1,2,3)
	assert.NotEmpty(t, IntSlice)
	assert.Len(t, IntSlice.Data, 3)

	found := IntSlice.Find(func(el int) bool { return el > 1})
	assert.NotNil(t, found)
	assert.Equal(t, 2, *found)

	notFound := IntSlice.Find(func(el int) bool { return el < 0})
	assert.Nil(t, notFound)
}

func Test_FindLast(t *testing.T) {
	IntSlice := Of(1,2,3)
	assert.NotEmpty(t, IntSlice)
	assert.Len(t, IntSlice.Data, 3)

	found := IntSlice.FindLast(func(el int) bool { return el > 1})
	assert.NotNil(t, found)
	assert.Equal(t, 3, *found)

	notFound := IntSlice.FindLast(func(el int) bool { return el < 0})
	assert.Nil(t, notFound)
}

func Test_FindIndex(t *testing.T) {
	IntSlice := Of(1,2,3)
	assert.NotEmpty(t, IntSlice)
	assert.Len(t, IntSlice.Data, 3)

	assert.Equal(t, 1, IntSlice.FindIndex(func(el int) bool { return el > 1}))

	// not found
	assert.Equal(t, -1, IntSlice.FindIndex(func(el int) bool { return el < 0}))
}

func Test_FindLastIndex(t *testing.T) {
	IntSlice := Of(1,2,3)
	assert.NotEmpty(t, IntSlice)
	assert.Len(t, IntSlice.Data, 3)
	
	assert.Equal(t, 2, IntSlice.FindLastIndex(func(el int) bool { return el > 1}))

	// not found
	assert.Equal(t, -1, IntSlice.FindLastIndex(func(el int) bool { return el < 0}))
}

func Test_ForEach(t *testing.T) {
	IntSlice := Of(1,2,3)
	assert.NotEmpty(t, IntSlice)
	assert.Len(t, IntSlice.Data, 3)

	IntSlice.ForEach(func(el int) { el++ })
	assert.Equal(t, 1, IntSlice.Data[0])
	assert.Equal(t, 2, IntSlice.Data[1])
	assert.Equal(t, 3, IntSlice.Data[2])
}

func Test_Includes(t *testing.T) {
	IntSlice := Of(1,2,3)
	assert.NotEmpty(t, IntSlice)
	assert.Len(t, IntSlice.Data, 3)

	assert.True(t, IntSlice.Includes(1))
	assert.False(t, IntSlice.Includes(4))
}

func Test_Index(t *testing.T) {
	IntSlice := Of(1,2,3)
	assert.NotEmpty(t, IntSlice)
	assert.Len(t, IntSlice.Data, 3)

	assert.Equal(t, 0, IntSlice.IndexOf(1))
	assert.Equal(t, -1, IntSlice.IndexOf(5))
}

func Test_Join(t *testing.T) {
	IntSlice := Of(1,2,3)
	assert.NotEmpty(t, IntSlice)
	assert.Len(t, IntSlice.Data, 3)

	assert.Equal(t, "1,2,3", IntSlice.Join(","))
}

func Test_Keys(t *testing.T) {
	IntSlice := Of(1,2,3)
	assert.NotEmpty(t, IntSlice)
	assert.Len(t, IntSlice.Data, 3)

	keys := IntSlice.Keys()
	assert.Len(t, keys, 3)
	assert.Equal(t, 0, keys[0])
	assert.Equal(t, 1, keys[1])
	assert.Equal(t, 2, keys[2])
}

func Test_LastIndex(t *testing.T) {
	IntSlice := Of(1,2,3)
	assert.NotEmpty(t, IntSlice)
	assert.Len(t, IntSlice.Data, 3)

	assert.Equal(t, 2, IntSlice.LastIndexOf(3))
	assert.Equal(t, -1, IntSlice.LastIndexOf(5))
}

func Test_Map(t *testing.T) {
	IntSlice := Of(1,2,3)
	assert.NotEmpty(t, IntSlice)
	assert.Len(t, IntSlice.Data, 3)

	newIntSlice := IntSlice.Map(func(el int) int { return el * 2 })
	assert.Len(t, newIntSlice.Data, 3)
	assert.Equal(t, 2, newIntSlice.Data[0])
	assert.Equal(t, 4, newIntSlice.Data[1])
	assert.Equal(t, 6, newIntSlice.Data[2])
}

func Test_Pop(t *testing.T) {
	IntSlice := Of(1,2,3)
	assert.NotEmpty(t, IntSlice)
	assert.Len(t, IntSlice.Data, 3)

	popped := IntSlice.Pop()

	assert.NotNil(t, popped)
	assert.Equal(t, 3, *popped)
	assert.Len(t, IntSlice.Data, 2)

	IntSlice.Pop()
	IntSlice.Pop()
	assert.Len(t, IntSlice.Data, 0)

	popped = IntSlice.Pop()
	assert.Nil(t, popped)
}

func Test_Push(t *testing.T) {
	IntSlice := Of(1,2,3)
	assert.NotEmpty(t, IntSlice)
	assert.Len(t, IntSlice.Data, 3)

	IntSlice.Push(4)
	assert.Len(t, IntSlice.Data, 4)
	assert.Equal(t, 4, IntSlice.Data[3])
}

func Test_Reverse(t *testing.T) {
	IntSlice := Of(1,2,3)
	assert.NotEmpty(t, IntSlice)
	assert.Len(t, IntSlice.Data, 3)

	newIntSlice := IntSlice.Reverse()
	assert.Len(t, newIntSlice.Data, 3)
	assert.Equal(t, 3, newIntSlice.Data[0])
	assert.Equal(t, 2, newIntSlice.Data[1])
	assert.Equal(t, 1, newIntSlice.Data[2])
}

func Test_Shift(t *testing.T) {
	IntSlice := Of(1,2,3)
	assert.NotEmpty(t, IntSlice)
	assert.Len(t, IntSlice.Data, 3)

	shifted := IntSlice.Shift()
	assert.Len(t, shifted.Data, 2)
	assert.Equal(t, 2, shifted.Data[0])
}

func Test_Slice(t *testing.T) {
	IntSlice := Of(1,2,3)
	assert.NotEmpty(t, IntSlice)
	assert.Len(t, IntSlice.Data, 3)

	newIntSlice := IntSlice.Slice(0, 2)
	assert.Len(t, newIntSlice.Data, 2)
	assert.Equal(t, 1, newIntSlice.Data[0])
	assert.Equal(t, 2, newIntSlice.Data[1])
}

func Test_Some(t *testing.T) {
	IntSlice := Of(1,2,3)
	assert.NotEmpty(t, IntSlice)
	assert.Len(t, IntSlice.Data, 3)

	assert.True(t, IntSlice.Some(func(el int) bool { return el > 1}))
	assert.False(t, IntSlice.Some(func(el int) bool { return el > 3}))
}

func Test_Sort(t *testing.T) {
	IntSlice := Of(3,2,1)
	assert.NotEmpty(t, IntSlice)
	assert.Len(t, IntSlice.Data, 3)

	newIntSlice := IntSlice.Sort(func(a, b int) bool { return b < a})
	assert.Len(t, newIntSlice.Data, 3)
	assert.Equal(t, 1, newIntSlice.Data[0])
	assert.Equal(t, 2, newIntSlice.Data[1])
	assert.Equal(t, 3, newIntSlice.Data[2])
}

func Test_Take(t *testing.T) {
	IntSlice := Of(1,2,3)
	assert.NotEmpty(t, IntSlice)
	assert.Len(t, IntSlice.Data, 3)

	newIntSlice := IntSlice.Take(2)
	assert.Len(t, newIntSlice.Data, 2)
	assert.Equal(t, 1, newIntSlice.Data[0])
	assert.Equal(t, 2, newIntSlice.Data[1])

	newIntSlice_2 := IntSlice.Take(100)
	assert.Len(t, newIntSlice_2.Data, 3)
}

func Test_ToString(t *testing.T) {
	IntSlice := Of(1,2,3)
	assert.NotEmpty(t, IntSlice)
	assert.Len(t, IntSlice.Data, 3)

	assert.Equal(t, "1,2,3", IntSlice.ToString())
}

func Test_Uniq(t *testing.T) {
	IntSlice := Of(1,2,3,3,2,1)
	assert.NotEmpty(t, IntSlice)
	assert.Len(t, IntSlice.Data, 6)

	newIntSlice := IntSlice.Unique()
	assert.Len(t, newIntSlice.Data, 3)
	assert.Equal(t, 1, newIntSlice.Data[0])
	assert.Equal(t, 2, newIntSlice.Data[1])
	assert.Equal(t, 3, newIntSlice.Data[2])
}

func Test_Unshift(t *testing.T) {
	IntSlice := Of(1,2,3)
	assert.NotEmpty(t, IntSlice)
	assert.Len(t, IntSlice.Data, 3)

	unshifted := IntSlice.Unshift(-2, -1, 0)
	assert.Len(t, unshifted.Data, 6)
	assert.Equal(t, -2, unshifted.Data[0])
}

func Test_Values(t *testing.T) {
	IntSlice := Of(1,2,3)
	assert.NotEmpty(t, IntSlice)
	assert.Len(t, IntSlice.Data, 3)

	assert.Equal(t, 1, IntSlice.Values()[0])
	assert.Equal(t, 2, IntSlice.Values()[1])
	assert.Equal(t, 3, IntSlice.Values()[2])
}

// extra tests

func Test_Chaining(t *testing.T) {
	IntSlice := Of(1,2,3)
	assert.NotEmpty(t, IntSlice)
	assert.Len(t, IntSlice.Data, 3)

	newIntSlice := IntSlice.
		Filter(func(el int) bool { return el > 1}). // [2,3]
		Map(func(el int) int { return el * 2}). // [4,6]
		Reverse(). // [6,4]
		Slice(0, 1) // [6]

	assert.Len(t, newIntSlice.Data, 1)
	assert.Equal(t, 6, newIntSlice.Data[0])
}