package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserDefinedEquality(t *testing.T) {
	person1 := NewUserDefined("Person", map[string]BaseType{"age": NewInt(10)})
	person2 := NewUserDefined("Person", map[string]BaseType{"age": NewInt(10)})

	res, err := person1.Equal(person2)
	assert.NoError(t, err)
	assert.True(t, res)

	person1 = NewUserDefined("Person", map[string]BaseType{"age": NewInt(10)})
	person2 = NewUserDefined("Person", map[string]BaseType{"age": NewInt(20)})

	res, err = person1.Equal(person2)
	assert.NoError(t, err)
	assert.False(t, res)

	person1 = NewUserDefined("Person", map[string]BaseType{"age": NewInt(10), "isEmployeed": NewBool(true)})
	person2 = NewUserDefined("Person", map[string]BaseType{"age": NewInt(20)})

	res, err = person1.Equal(person2)
	assert.False(t, res)

	person1 = NewUserDefined("Person", map[string]BaseType{"age": NewInt(20)})
	person2 = NewUserDefined("Person", map[string]BaseType{"age": NewInt(10), "isEmployeed": NewBool(true)})

	res, err = person1.Equal(person2)
	assert.False(t, res)
}

func TestUserDefinedComparison(t *testing.T) {
	person1 := NewUserDefined("Person", map[string]BaseType{"age": NewInt(10)})
	person2 := NewUserDefined("Person", map[string]BaseType{"age": NewInt(20)})

	res, err := person1.LessThan(person2)
	assert.NoError(t, err)
	assert.True(t, res)

	person1 = NewUserDefined("Person", map[string]BaseType{"age": NewInt(10)})
	person2 = NewUserDefined("Person", map[string]BaseType{"age": NewInt(8)})

	res, err = person1.LessThan(person2)
	assert.NoError(t, err)
	assert.False(t, res)

	person1 = NewUserDefined("Person", map[string]BaseType{"age": NewInt(10), "isEmployeed": NewBool(true)})
	person2 = NewUserDefined("Person", map[string]BaseType{"age": NewInt(20)})

	res, err = person1.LessThan(person2)
	assert.True(t, res)

	person1 = NewUserDefined("Person", map[string]BaseType{"age": NewInt(20)})
	person2 = NewUserDefined("Person", map[string]BaseType{"age": NewInt(10), "isEmployeed": NewBool(true)})

	res, err = person1.LessThan(person2)
	assert.False(t, res)

  person1 = NewUserDefined("Person", map[string]BaseType{"age": NewInt(20), "cloths": NewInt(4)})
  person2 = NewUserDefined("Person", map[string]BaseType{"age": NewInt(10), "cloths": NewInt(10)})

	res, err = person1.LessThan(person2)
	assert.False(t, res)

  person1 = NewUserDefined("Person", map[string]BaseType{"age": NewInt(20), "cloths": NewInt(40)})
  person2 = NewUserDefined("Person", map[string]BaseType{"age": NewInt(40), "cloths": NewInt(80)})

	res, err = person1.LessThan(person2)
	assert.True(t, res)
}
