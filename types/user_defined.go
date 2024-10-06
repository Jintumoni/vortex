package types

import (
	"encoding/json"
	"fmt"

	"github.com/Jintumoni/vortex/errors"
)

type UserDefinedType struct {
	TypeName   string
	Properties map[string]BaseType
}

func NewUserDefined(typeName string, properties map[string]BaseType) *UserDefinedType {
  return &UserDefinedType{typeName, properties}
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

    if b, err := v.Equal(otherProp.(Equality)); err != nil || !b {
      return false, err
    }
	}

	return true, nil
}

func (u *UserDefinedType) LessThan(other Comparison) (bool, error) {
	o, ok := other.(*UserDefinedType)
	if !ok || u.TypeName != o.TypeName {
		return false, errors.InvalidOperation
	}
	for k, v := range u.Properties {
		otherProp, err := o.GetProperty(k)
		if err != nil {
			return false, err
		}
    b, err := v.Equal(otherProp.(Equality))
    if err != nil {
      return false, nil
    }

    if b {
      continue
    }

    b, err = v.LessThan(otherProp.(Comparison))
    if err != nil {
      return false, nil
    }
    return b, nil
	}

	return false, nil
}

func (u *UserDefinedType) Less(other *UserDefinedType) bool {
	r, err := u.LessThan(other)
	if err != nil {
		return false
	}
	return r
}

func (u *UserDefinedType) Repr() string {
  bytes, _ := json.Marshal(u)
  return fmt.Sprintf("%v", string(bytes))
}
