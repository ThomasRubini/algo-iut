package scanexpr

const (
	CompVar   = iota
	CompFunc  = iota
	CompArr   = iota
	CompOp    = iota
	CompMerge = iota
)

// Component
type Comp interface {
	Type() int
}

type CompVarImpl struct {
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
type CompMergeImpl struct {
	Comps []Comp
}

func (c CompVarImpl) Type() int {
	return CompVar
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
