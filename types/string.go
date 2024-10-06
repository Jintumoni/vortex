package types

import "fmt"

type StringType struct {
	Value string
}

func NewString(value string) *StringType {
	return &StringType{Value: value}
}

func (i *StringType) Repr() string {
	return fmt.Sprintf("%s", i.Value)
}

func (i *StringType) LessThan(other Comparison) (bool, error) {
	o, ok := other.(*StringType)
	if !ok {

	}

	return i.Value < o.Value, nil
}

func (i *StringType) Equal(other Equality) (bool, error) {
	o, ok := other.(*StringType)
	if !ok {
	}

	return i.Value == o.Value, nil
}
