package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoolTypeAndIntAndBool(t *testing.T) {
	x := NewBool(true)
	y := NewInt(1)

	z, err := x.And(y)

	assert.NoError(t, err)
	assert.Equal(t, 1, z.(*IntType).Value)

	x = NewBool(false)
	y = NewInt(0)

	z, err = x.And(y)

	assert.NoError(t, err)
	assert.Equal(t, false, z.(*BoolType).Value)
}

func TestBoolTypeOrIntAndBool(t *testing.T) {
	x := NewBool(true)
	y := NewInt(1)

	z, err := x.Or(y)

	assert.NoError(t, err)
	assert.Equal(t, true, z.(*BoolType).Value)

	x = NewBool(false)
	y = NewInt(0)

	z, err = x.Or(y)

	assert.NoError(t, err)
	assert.Equal(t, 0, z.(*IntType).Value)
}

func TestBoolTypeAndBoolAndBool(t *testing.T) {
	x := NewBool(true)
	y := NewBool(false)

	z, err := x.And(y)

	assert.NoError(t, err)
	assert.Equal(t, false, z.(*BoolType).Value)

	x = NewBool(true)
	y = NewBool(true)

	z, err = x.And(y)

	assert.NoError(t, err)
	assert.Equal(t, true, z.(*BoolType).Value)
}

func TestBoolTypeAndBoolOrBool(t *testing.T) {
	x := NewBool(true)
	y := NewBool(false)

	z, err := x.Or(y)

	assert.NoError(t, err)
	assert.Equal(t, true, z.(*BoolType).Value)

	x = NewBool(true)
	y = NewBool(true)

	z, err = x.Or(y)

	assert.NoError(t, err)
	assert.Equal(t, true, z.(*BoolType).Value)
}

func TestBoolTypeRepr(t *testing.T) {
	x := NewBool(true)
	z := x.Repr()

	assert.Equal(t, "true", z)
}

func TestBoolLessThan(t *testing.T) {
  x := NewBool(true)
  y := NewBool(false)

  z, err := x.LessThan(y)
  assert.NoError(t, err)
  assert.Equal(t, false, z)

  x = NewBool(true)
  y = NewBool(true)

  z, err = x.LessThan(y)
  assert.NoError(t, err)
  assert.Equal(t, false, z)

  x = NewBool(false)
  y = NewBool(true)

  z, err = x.LessThan(y)
  assert.NoError(t, err)
  assert.Equal(t, true, z)

  x = NewBool(false)
  y = NewBool(false)

  z, err = x.LessThan(y)
  assert.NoError(t, err)
  assert.Equal(t, false, z)
}
