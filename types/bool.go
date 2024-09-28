package types

import "fmt"

type BoolType struct {
	Value bool
}

func NewBool(value bool) *BoolType {
	return &BoolType{Value: value}
}

func (i *BoolType) And(other Logical) (Logical, error) {
	if !i.Value {
		return i, nil
	}
	return other, nil
}

func (i *BoolType) Or(other Logical) (Logical, error) {
	if i.Value {
		return i, nil
	}
	return other, nil
}

func (i *BoolType) Repr() string {
	return fmt.Sprintf("%t", i.Value)
}

func (i *BoolType) Equal(other Equality) (bool, error) {
	o, ok := other.(*BoolType)
	if !ok {
	}

	return i.Value == o.Value, nil
}
