package demo

import (
	"fmt"
	"strings"
)

type Query struct {
	SQL string
	Arg []interface{}
}


type Selector struct {
	tbl string
	where 	[]Predicate
	sb strings.Builder
	args []int
	db *DB
}


func (s *Selector) From(tbl string) *Selector {
	s.tbl = tbl
	return s
}


func (s *Selector) Where(ps ...Predicate) *Selector {
	s.where = ps
	return s
}


func (s *Selector) Having(ps ...Predicate) *Selector {
	
	// TODO
	panic("")
}



func (s *Selector) Build() (*Query, error) {
	
	s.sb.WriteString("SELECT * FROM `")
	s.sb.WriteString(s.tbl)
	s.sb.WriteString("`")

	
	// 构建 WHERE
	if len(s.where) > 0 {
		s.sb.WriteString(" WHERE")
		p := s.where[0]
		if err := s.buildExpression(p); err != nil {
			return nil, err
		}
	}

	s.sb.WriteString(";")

	return &Query{
		SQL: s.sb.String(),
	}, nil
}


func (s *Selector) buildExpression(e Expression) error {
	if e == nil {
		return nil
	}
	
	// 反射 实现指针类型 与 原生类型区别
	switch exp := e.(type) {
	case Column:
		s.sb.WriteString(fmt.Sprintf(" %s ", exp.Name))

	case Value:
		s.sb.WriteString(" ? ")
		s.args = append(s.args, exp.Val)
	case Predicate:
		if err := s.buildExpression(exp.left); err != nil {
			return err
		}
		s.sb.WriteString(exp.op)
		if err := s.buildExpression(exp.right); err != nil {
			return err
		}

	default:
		return fmt.Errorf("toy-orm：不支持表达式 %v", exp)
	}

	return nil
} 