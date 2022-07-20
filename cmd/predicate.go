package demo


const (
	opEQ = "="
)


// 谓词
type Predicate struct {
	left Expression
	op string
	right Expression
}

func (p Predicate) expr() {}


type Expression interface {
	expr()
}
