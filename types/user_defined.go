package types

import (
	"github.com/Jintumoni/vortex/errors"
)

type UserDefinedType struct {
	TypeName   string
	Properties map[string]BaseType
}

func (u *UserDefinedType) GetProperty(name string) (BaseType, error) {
	p, ok := u.Properties[name]
	if !ok {
		return nil, errors.UnknownField
	}

	return p, nil
}

func (u *UserDefinedType) Equal(other Equality) (bool, error) {
	o, ok := other.(*UserDefinedType)
	if !ok || u.TypeName != o.TypeName {
		return false, errors.InvalidOperation
	}
	for k, v := range u.Properties {
		otherProp, err := o.GetProperty(k)
		if err != nil {
			return false, err
		}

		switch v.(type) {
		case IntType:
			if b, err := v.(*IntType).Equal(otherProp.(*IntType)); err != nil && !b {
				return false, err
			}
		case BoolType:
			if b, err := v.(*BoolType).Equal(otherProp.(*BoolType)); err != nil && !b {
				return false, err
			}
		case UserDefinedType:
			if b, err := v.(*UserDefinedType).LessThan(otherProp.(*UserDefinedType)); err != nil && !b {
				return false, err
			}
		default:
			return false, errors.UnknownType
		}
	}

	return true, nil
}

func (u *UserDefinedType) LessThan(other Equality) (bool, error) {
	o, ok := other.(*UserDefinedType)
	if !ok || u.TypeName != o.TypeName {
		return false, errors.InvalidOperation
	}
	for k, v := range u.Properties {
		otherProp, err := o.GetProperty(k)
		if err != nil {
			return false, err
		}
		switch v.(type) {
		case IntType:
			if b, err := v.(*IntType).LessThan(otherProp.(*IntType)); err != nil && !b {
				return false, err
			}
		case UserDefinedType:
			if b, err := v.(*UserDefinedType).LessThan(otherProp.(*UserDefinedType)); err != nil && !b {
				return false, err
			}
		default:
			return false, errors.InvalidOperation
		}
	}

	return true, nil
}

func (u *UserDefinedType) Less(other *UserDefinedType) bool {
	r, err := u.LessThan(other)
	if err != nil {
		return false
	}
	return r
}
