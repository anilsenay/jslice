package examples

import (
	"testing"

	"github.com/anilsenay/jslice"
	"github.com/stretchr/testify/assert"
)

func Test_New(t *testing.T) {
  newIntSlice := jslice.Of(1, 2, 3)
  assert.NotEmpty(t, newIntSlice)
  assert.Len(t, newIntSlice.Data, 3)

  newIntSlice.Append(4, 5, 6)
  assert.Len(t, newIntSlice.Data, 6)
}

type Person struct {
  Name string
  Age  int
  Courses jslice.Slice[string]
}
func Test_Person(t *testing.T) {
  newPerson := Person{
    Name: "John",
    Age:  30,
    Courses: *jslice.Of("Math", "English", "Science"),
  }

  assert.NotEmpty(t, newPerson)
  assert.Equal(t, 3, newPerson.Courses.Length())

  newPerson.Courses.Append("History")
  assert.Equal(t, 4, newPerson.Courses.Length())

  strArr := newPerson.Courses.Values()
  assert.Equal(t, 4, len(strArr))
}