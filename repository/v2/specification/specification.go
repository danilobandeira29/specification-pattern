package specification

import "fmt"

type Specification interface {
	And(other Specification) Specification
	Or(other Specification) Specification
	Not() Specification
	ToSQL() (string, []interface{})
}

// New
// example:
//
//	     package main
//	     import (
//		        "fmt"
//		        "github.com/danilobandeira29/specification-pattern/repository/v2"
//		        "github.com/danilobandeira29/specification-pattern/repository/v2/specification"
//	     )
//	     func main() {
//	             spec := specification.New()
//	             spec.And(specification.PriceBelow(10.00))
//	             products, err := repository.FindBy(spec)
//	     }
func New() Specification {
	return &specification{
		sql:  "true",
		args: make([]interface{}, 0),
	}
}

type specification struct {
	sql  string
	args []interface{}
}

func (s *specification) ToSQL() (string, []interface{}) {
	return s.sql, s.args
}

func (s *specification) And(other Specification) Specification {
	otherSQL, otherArgs := other.ToSQL()
	return &specification{
		sql:  fmt.Sprintf("(%s and %s)", s.sql, otherSQL),
		args: append(s.args, otherArgs...),
	}
}

func (s *specification) Or(other Specification) Specification {
	otherSQL, otherArgs := other.ToSQL()
	return &specification{
		sql:  fmt.Sprintf("(%s OR %s)", s.sql, otherSQL),
		args: append(s.args, otherArgs...),
	}
}

func (s *specification) Not() Specification {
	return &specification{
		sql:  fmt.Sprintf("(NOT %s)", s.sql),
		args: s.args,
	}
}
