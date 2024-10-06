package types

import (
	"reflect"
	"sort"
	"strings"

	"github.com/Jintumoni/vortex/errors"
)

type ListType struct {
	Value []BaseType
}

func NewList(value []BaseType) (*ListType, error) {
	// check if all the elements of the list are of same type
	for _, e := range value {
		if reflect.TypeOf(e) != reflect.TypeOf(value[0]) {
			return nil, errors.UnhomogeneousType
		}
	}
	return &ListType{Value: value}, nil
}

func (l *ListType) And(other Logical) (Logical, error) {
	// If the other is a list type, then intersection
	// of two vectors would be returned
	o, ok := other.(*ListType)
	if !ok {
		return nil, errors.InvalidOperation
	}

	// Validate if both the list have same type of elements
  // A list is expected to be homogeneous, so validating only the first element works
	if reflect.TypeOf(l.Value[0]) != reflect.TypeOf(o.Value[0]) {
		return nil, errors.InvalidOperation
	}

	return intersect(l, o)
}

func (l *ListType) Or(other Logical) (Logical, error) {
	// If the other is a list type, then intersection
	// of two vectors would be returned
	o, ok := other.(*ListType)
	if !ok {
		return nil, errors.InvalidOperation
	}

	// Validate if both the list have same type of elements
  // A list is expected to be homogeneous, so validating only the first element works
	if reflect.TypeOf(l.Value[0]) != reflect.TypeOf(o.Value[0]) {
		return nil, errors.InvalidOperation
	}

	return unite(l, o)
}

func (l *ListType) Repr() string {
  var s strings.Builder
  s.WriteString("[")
  for i, e := range l.Value {
    if i > 0 {
      s.WriteString(", ")
    }
    s.WriteString(e.Repr())
  }
  s.WriteString("]")

  return s.String()
}

func (l *ListType) validateType() error {
	for _, e := range l.Value {
		if reflect.TypeOf(e) != reflect.TypeOf(l.Value[0]) {
			return errors.UnhomogeneousType
		}
	}
  return nil
}

func intersect(a *ListType, b *ListType) (*ListType, error) {
	// intersect the two lists
	sort.Slice(a.Value, func(i, j int) bool {
		less, err := a.Value[i].LessThan(a.Value[j])
		if err != nil {
			return false
		}
		return less
	})

	sort.Slice(b.Value, func(i, j int) bool {
		less, err := b.Value[i].LessThan(b.Value[j])
		if err != nil {
			return false
		}
		return less
	})

	var res ListType
	p1, p2 := 0, 0

	for p1 < len(a.Value) && p2 < len(b.Value) {
		lt, err := a.Value[p1].LessThan(b.Value[p2])
		if err != nil {
			return nil, err
		}
		eq, err := a.Value[p1].Equal(b.Value[p2])
		if err != nil {
			return nil, err
		}

		if lt {
			p1++
		} else if eq {
			res.Value = append(res.Value, a.Value[p1])
			p1++
			p2++
		} else {
			p2++
		}
	}

	return &res, nil
}

func unite(a *ListType, b *ListType) (*ListType, error) {
	// intersect the two lists
	sort.Slice(a.Value, func(i, j int) bool {
		less, err := a.Value[i].LessThan(a.Value[j])
		if err != nil {
			return false
		}
		return less
	})

	sort.Slice(b.Value, func(i, j int) bool {
		less, err := b.Value[i].LessThan(b.Value[j])
		if err != nil {
			return false
		}
		return less
	})

	var res ListType
	p1, p2 := 0, 0

	for p1 < len(a.Value) && p2 < len(b.Value) {
		lt, err := a.Value[p1].LessThan(b.Value[p2])
		if err != nil {
			return nil, err
		}
		eq, err := a.Value[p1].Equal(b.Value[p2])
		if err != nil {
			return nil, err
		}

		if lt {
      res.Value = append(res.Value, a.Value[p1])
			p1++
		} else if eq {
      res.Value = append(res.Value, a.Value[p1])
			p1++
			p2++
		} else {
      res.Value = append(res.Value, b.Value[p2])
			p2++
		}
	}

	for p1 < len(a.Value) {
		res.Value = append(res.Value, a.Value[p1])
    p1++
	}
	for p2 < len(b.Value) {
		res.Value = append(res.Value, b.Value[p2])
    p2++
	}

	return &res, nil
}
