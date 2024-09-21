package nodes

type EdgeType int

const (
	OneWayEdge EdgeType = iota + 1
	TwoWayEdge
)

func (e EdgeType) String() string {
	switch e {
	case OneWayEdge:
		return "OneWay"
	case TwoWayEdge:
		return "TwoWay"
	default:
		return ""
	}
}

func GetAllEdgeTypes() []EdgeType {
	return []EdgeType{
		OneWayEdge,
		TwoWayEdge,
	}
}

type FuncType int

const (
	SumFunc FuncType = iota + 1
	MaxFunc
	MinFunc
	StartWithFunc
)

func (e FuncType) String() string {
	switch e {
	case SumFunc:
		return "Sum"
	case MaxFunc:
		return "Max"
	case MinFunc:
		return "Min"
	case StartWithFunc:
		return "StartsWith"
	default:
		return ""
	}
}

func GetAllFuncTypes() []FuncType {
	return []FuncType{
		SumFunc,
		MaxFunc,
		MinFunc,
		StartWithFunc,
	}
}
