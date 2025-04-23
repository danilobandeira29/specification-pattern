# Specification Pattern

```bash
go mod tidy

go run main.go
```

## Example Usage

```go
package main

import (
        "fmt"
        "github.com/danilobandeira29/specification-pattern/repository/v2"
        "github.com/danilobandeira29/specification-pattern/repository/v2/specification"
)

func main() {
        conn, err := v2.SQLiteConn()
	    if err != nil {
		        fmt.Printf("error conn: %v", err)
		        return
	    }
	    repo := v2.New(conn)
	    // this way works to
	    // spec := specification.New().
	    //	And(specification.NameContains("Adidas")).
	    //	And(specification.Active()).
	    //	And(specification.PriceBelow(111))
	    spec := specification.NameContains("Adidas").
	    	And(specification.Active()).
	    	And(specification.PriceBelow(111))
	    products, err := repo.FindBy(spec)
	    if err != nil {
	    	fmt.Printf("error querying: %v", err)
	    	return
	    }
	    for _, p := range products {
	    	fmt.Println(p)
	    }
}
```

Links: https://elemarjr.com/clube-de-estudos/artigos/specification-pattern-o-que-e-para-que-serve-e-quando-adotar/
