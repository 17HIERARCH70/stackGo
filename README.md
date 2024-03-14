# Stack Package

A simple thread-safe stack implementation in Go.

## Features

- Supports thread-safe operations.
- Allows pushing and popping elements of any type onto the stack.

## Installation

To use this package in your Go module, run:

```sh
go get github.com/17HIERARCH70/stackGo
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/yourusername/stack"
)

func main() {
	// Create a new stack
	s := stack.NewStack()

	// Push elements onto the stack
	s.Push(1)
	s.Push("hello")
	s.Push(3.14)

	// Pop elements from the stack
	val, err := s.Pop()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Popped value:", val)
	}
}
```

## Contributing

Contributions are welcome! If you find any bugs or want to add new features, feel free to open an issue or submit a pull request.
