package types

type DataType int

type BaseType interface{}

type Arithmetic interface {
	Add(other Arithmetic) (Arithmetic, error)
	Sub(other Arithmetic) (Arithmetic, error)
	Mul(other Arithmetic) (Arithmetic, error)
	Div(other Arithmetic) (Arithmetic, error)
}

type Equality interface {
	Equal(other Equality) (bool, error)
}

type Comparison interface {
	LessThan(other Equality) (bool, error)
}

type Logical interface {
	Or(other Logical) (Logical, error)
	And(other Logical) (Logical, error)
}

type Representation interface {
	Repr() string
}

type PropertyAccess interface {
	GetProperty(name string) (BaseType, error)
	// SetProperty(name string, value BaseType) error
}
