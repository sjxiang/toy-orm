package demo


// id
type Column struct {
	Name string
}

func (c Column) expr() {}


// 1
type Value struct {
	Val int
}

func (v Value) expr() {}




func C(name string) Column {
	return Column{Name: name}
}

// EQ 例如 C("id").Eq(12)
func (c Column) EQ(arg int) Predicate {
	return Predicate{
		left:  c,
		op:    opEQ,
		right: Value{Val: arg},
	}
}
