package scanexpr

const (
	CompId    = iota
	CompFunc  = iota
	CompArr   = iota
	CompOp    = iota
	CompMerge = iota
)

// Component
type Comp interface {
	Type() int
}

type CompIdImpl struct {
	Name string
}
type CompFuncImpl struct {
	Name string
	Args []Comp
}
type CompArrImpl struct {
	Name  string
	Index Comp
}
type CompOpImpl struct {
	Op string
}

// merge **using operators**
type CompMergeImpl struct {
	Comps []Comp
}

func (c CompIdImpl) Type() int {
	return CompId
}

func (c CompFuncImpl) Type() int {
	return CompFunc
}

func (c CompArrImpl) Type() int {
	return CompArr
}

func (c CompOpImpl) Type() int {
	return CompOp
}

func (c CompMergeImpl) Type() int {
	return CompMerge
}

func Id(name string) Comp {
	return CompIdImpl{
		Name: name,
	}
}

func Func(name string, args ...Comp) Comp {
	return CompFuncImpl{
		Name: name,
		Args: args,
	}
}

func Arr(name string, index Comp) Comp {
	return CompArrImpl{
		Name:  name,
		Index: index,
	}
}

func Op(op string) Comp {
	return CompOpImpl{
		Op: op,
	}
}

func Merge(comps ...Comp) Comp {
	return CompMergeImpl{
		Comps: comps,
	}
}
