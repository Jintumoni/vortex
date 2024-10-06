package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEqualityBetweenStrings(t *testing.T) {
  s1 := NewString("hello")
  s2 := NewString("he")

  r, err := s1.Equal(s2)
  assert.NoError(t, err)
  assert.False(t, r)
}

func TestComparisonBetweenStrings(t *testing.T) {
  s1 := NewString("hello")
  s2 := NewString("he")

  r, err := s1.LessThan(s2)
  assert.NoError(t, err)
  assert.False(t, r)
}
