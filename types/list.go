package types

import (
	"reflect"
	"sort"

	"github.com/Jintumoni/vortex/errors"
)

type ListType struct {
	Value    []BaseType
	elemType reflect.Type
}

func NewList(value []BaseType, elemType reflect.Type) (*ListType, error) {
	// check if all the elements of the list are of same type
	for _, e := range value {
		if reflect.TypeOf(e) != elemType {
			return nil, errors.UnhomogeneousType
		}
	}
	return &ListType{
		Value:    value,
		elemType: elemType,
	}, nil
}

func (l *ListType) And(other Logical) (Logical, error) {
	// If the other is a list type, then intersection
	// of two vectors would be returned
	o, ok := other.(*ListType)
	if !ok {
		return nil, errors.InvalidOperation
	}

	// validate if both the list have same type of elements
	if l.elemType != o.elemType {
		return nil, errors.InvalidOperation
	}

	res, _ := NewList(nil, l.elemType)

	// intersect the two lists
	sort.Slice(l.Value, func(i, j int) bool {
		switch l.Value[i].(type) {
		case *IntType:
			less, err := l.Value[i].(*IntType).LessThan(l.Value[j].(*IntType))
			if err != nil {
				return false
			}
			return less
		case *BoolType:
			return false
		case *UserDefinedType:
			less, err := l.Value[i].(*UserDefinedType).LessThan(l.Value[j].(*UserDefinedType))
			if err != nil {
				return false
			}
			return less
		default:
			return false
		}
	})

	return res, nil
}

func (l *ListType) Or(other Logical) (Logical, error) {
	// If the other is a list type, then intersection
	// of two vectors would be returned
	o, ok := other.(*ListType)
	if !ok {
		return nil, errors.InvalidOperation
	}

	// validate if both the list have same type of elements
	if l.elemType != o.elemType {
		return nil, errors.InvalidOperation
	}

	res, _ := NewList(nil, l.elemType)

  // unite the two lists

  return res, nil
}
