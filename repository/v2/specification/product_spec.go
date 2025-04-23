package specification

import "strings"

func PriceBelow(max float64) Specification {
	return &specification{
		sql:  "price < ?",
		args: []interface{}{max},
	}
}

type NameContainsSpecification struct {
	Name string
}

func NameContains(n string) Specification {
	return &specification{
		sql:  "lower(name) like lower(?)",
		args: []interface{}{"%" + strings.ToLower(n) + "%"},
	}
}

func Active() Specification {
	return &specification{
		sql:  "is_active is true",
		args: nil,
	}
}
