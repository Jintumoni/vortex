package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntTypeAddTwoInt(t *testing.T) {
	x := NewInt(1)
	y := NewInt(2)

	z, err := x.Add(y)

	assert.NoError(t, err)
	assert.Equal(t, 3, z.(*IntType).Value)
}

func TestIntTypeSubTwoInt(t *testing.T) {
	x := NewInt(1)
	y := NewInt(2)

	z, err := x.Sub(y)

	assert.NoError(t, err)
	assert.Equal(t, -1, z.(*IntType).Value)
}

func TestIntTypeDivTwoInt(t *testing.T) {
	x := NewInt(1)
	y := NewInt(2)

	z, err := x.Div(y)

	assert.NoError(t, err)
	assert.Equal(t, 0, z.(*IntType).Value)
}

func TestIntTypeMulTwoInt(t *testing.T) {
	x := NewInt(1)
	y := NewInt(2)

	z, err := x.Mul(y)

	assert.NoError(t, err)
	assert.Equal(t, 2, z.(*IntType).Value)
}

func TestIntTypeAndIntAndBool(t *testing.T) {
	x := NewInt(1)
	y := NewBool(true)

	z, err := x.And(y)

	assert.NoError(t, err)
	assert.Equal(t, true, z.(*BoolType).Value)

	x = NewInt(0)
	y = NewBool(false)

	z, err = x.And(y)

	assert.NoError(t, err)
	assert.Equal(t, 0, z.(*IntType).Value)
}

func TestIntTypeOrIntAndBool(t *testing.T) {
	x := NewInt(1)
	y := NewBool(true)

	z, err := x.Or(y)

	assert.NoError(t, err)
	assert.Equal(t, 1, z.(*IntType).Value)

	x = NewInt(0)
	y = NewBool(false)

	z, err = x.Or(y)

	assert.NoError(t, err)
	assert.Equal(t, false, z.(*BoolType).Value)
}

func TestIntTypeRepr(t *testing.T) {
	x := NewInt(-1034343)
	z := x.Repr()

	assert.Equal(t, "-1034343", z)
}
