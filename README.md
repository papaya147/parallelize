# Parallelize

Parallelize your Go code with ease. This package abstracts away the direct use of channels and sync wait groups.

#### Example

```go
package main

import (
	"context"
	"fmt"

	"github.com/papaya147/parallelize"
)

func add(_ context.Context, args []int) (int, error) {
	var sum int = 0
	for _, v := range args {
		sum += v
	}
	return sum, nil
}

func multiply(_ context.Context, args []int) (int, error) {
	var product int = 1
	for _, v := range args {
		product *= v
	}
	return product, nil
}

func main() {
	g := parallelize.NewGroup()

	// add methods to the group
	res1 := parallelize.AddWithOutputWithArgs(g, add, context.Background(), []int{1, 2, 3, 4})
	res2 := parallelize.AddWithOutputWithArgs(g, multiply, context.Background(), []int{1, 2, 3, 4})

	// execute the group using go routines
	g.Execute()

	// extract outputs and use them!
	sum, err := res1.Read()
	if err != nil {
		panic(err)
	}

	prod, err := res2.Read()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Sum is %d\n", sum)      // Sum is 10
	fmt.Printf("Product is %d\n", prod) // Product is 24
}
```

### Supported Methods

- Method with output and with args
- Method with output and without args
- Method without output and with args
- Method without output and without args
