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

func someCoolParallelFunction(ctx context.Context, input int, output *int) error {
	*output = input + 1
	return nil
}

func anotherFunction(ctx context.Context, input int) error {
	fmt.Println(input)
	return nil
}

func main() {
	group := parallelize.NewSyncGroup()

	var b int
	parallelize.AddOutputtingMethodWithArgs(group, someCoolParallelFunction, parallelize.OutputtingMethodWithArgsParams[int, *int]{
		Context: context.TODO(),
		Input:   1,
		Output:  &b,
	})

	var c int
	parallelize.AddOutputtingMethodWithArgs(group, someCoolParallelFunction, parallelize.OutputtingMethodWithArgsParams[int, *int]{
		Context: context.TODO(),
		Input:   5,
		Output:  &c,
	})

	parallelize.AddMethodWithArgs(group, anotherFunction, parallelize.MethodWithArgsParams[int]{
		Context: context.TODO(),
		Input:   10,
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
