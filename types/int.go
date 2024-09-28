package types

import "fmt"

type IntType struct {
	Value int
}

func NewInt(value int) *IntType {
	return &IntType{Value: value}
}

func (i *IntType) Add(other Arithmetic) (Arithmetic, error) {
	o, ok := other.(*IntType)
	if !ok {
	}

	return &IntType{Value: i.Value + o.Value}, nil
}

func (i *IntType) Sub(other Arithmetic) (Arithmetic, error) {
	o, ok := other.(*IntType)
	if !ok {
	}

	return &IntType{Value: i.Value - o.Value}, nil
}

func (i *IntType) Mul(other Arithmetic) (Arithmetic, error) {
	o, ok := other.(*IntType)
	if !ok {
	}

	return &IntType{Value: i.Value * o.Value}, nil
}

func (i *IntType) Div(other Arithmetic) (Arithmetic, error) {
	o, ok := other.(*IntType)
	if !ok {
	}

	if o.Value == 0 {

	}

	return &IntType{Value: i.Value / o.Value}, nil
}

func (i *IntType) Repr() string {
	return fmt.Sprintf("%d", i.Value)
}

func (i *IntType) And(other Logical) (Logical, error) {
	if i.Value == 0 {
		return i, nil
	}
	return other, nil
}

func (i *IntType) Or(other Logical) (Logical, error) {
	if i.Value == 0 {
		return other, nil
	}
	return i, nil
}

func (i *IntType) LessThan(other Equality) (bool, error) {
	o, ok := other.(*IntType)
	if !ok {

	}

	return i.Value < o.Value, nil
}

func (i *IntType) Equal(other Equality) (bool, error) {
	o, ok := other.(*IntType)
	if !ok {
	}

	return i.Value == o.Value, nil
}
