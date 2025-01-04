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
