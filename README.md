# Parallelize

Parallelize Go methods using this package.

# Install

Install via `go get`.

```sh
# After: go mod init ...
go get github.com/papaya147/parallelize
```

# Example

Here is a very basic example to show the functionality of this package:

```go
package main

import (
	"context"
	"fmt"

	"github.com/papaya147/parallelize"
)

func someCoolParallelFunction(ctx context.Context, a int, b *int) error {
	*b = a + 1
	return nil
}

func main() {
	group := parallelize.NewSyncGroup()

	var b int
	parallelize.AddMethod(group, someCoolParallelFunction, parallelize.OutputArgMethodArgs[int, *int]{
		Context: context.TODO(),
		Arg1:    1,
		Arg2:    &b,
	})

	var c int
	parallelize.AddMethod(group, someCoolParallelFunction, parallelize.OutputArgMethodArgs[int, *int]{
		Context: context.TODO(),
		Arg1:    5,
		Arg2:    &c,
	})

	if err := group.Run(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("executed successfully!")
		fmt.Printf("b is %d\n", b)
		fmt.Printf("c is %d\n", c)
	}
}

```
