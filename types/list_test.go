package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAndBetweenTwoDifferentLists(t *testing.T) {
  l1, _ := NewList([]BaseType{NewInt(1), NewInt(2)})
  l2, _ := NewList([]BaseType{NewBool(true), NewBool(false)})

  r, err := l1.And(l2)
  assert.Error(t, err)
  assert.Nil(t, r)
}

func TestAndBetweenTwoIntLists(t *testing.T) {
  l1, _ := NewList([]BaseType{NewInt(1), NewInt(2)})
  l2, _ := NewList([]BaseType{NewInt(2), NewInt(3)})

  r, err := l1.And(l2)
  assert.NoError(t, err)

  res := r.(*ListType)
  assert.Len(t, res.Value, 1)

  assert.Equal(t, 2, res.Value[0].(*IntType).Value)
}

func TestOrBetweenTwoIntLists(t *testing.T) {
  l1, _ := NewList([]BaseType{NewInt(1), NewInt(2)})
  l2, _ := NewList([]BaseType{NewInt(2), NewInt(3)})

  r, err := l1.Or(l2)
  assert.NoError(t, err)

  res := r.(*ListType)
  assert.Len(t, res.Value, 3)

  assert.Equal(t, 1, res.Value[0].(*IntType).Value)
  assert.Equal(t, 2, res.Value[1].(*IntType).Value)
  assert.Equal(t, 3, res.Value[2].(*IntType).Value)
}

func TestAndBetweenTwoUserDefinedLists(t *testing.T) {
  l1, _ := NewList([]BaseType{NewUserDefined("Person", map[string]BaseType{"age": NewInt(10), "isEmployeed": NewBool(true)})})
  l2, _ := NewList([]BaseType{NewUserDefined("Person", map[string]BaseType{"age": NewInt(12), "isEmployeed": NewBool(true)})})

  r, err := l1.And(l2)
  assert.NoError(t, err)

  res := r.(*ListType)
  assert.Len(t, res.Value, 0)
}

func TestOrBetweenTwoUserDefinedLists(t *testing.T) {
  l1, _ := NewList([]BaseType{NewUserDefined("Person", map[string]BaseType{"age": NewInt(10), "name": NewString("a")})})
  l2, _ := NewList([]BaseType{NewUserDefined("Person", map[string]BaseType{"age": NewInt(12), "name": NewString("b")})})

  r, err := l1.Or(l2)
  assert.NoError(t, err)

  res := r.(*ListType)
  assert.Len(t, res.Value, 2)

  assert.Equal(t, 10, res.Value[0].(*UserDefinedType).Properties["age"].(*IntType).Value)
  assert.Equal(t, "a", res.Value[0].(*UserDefinedType).Properties["name"].(*StringType).Value)

  assert.Equal(t, 12, res.Value[1].(*UserDefinedType).Properties["age"].(*IntType).Value)
  assert.Equal(t, "b", res.Value[1].(*UserDefinedType).Properties["name"].(*StringType).Value)
}
