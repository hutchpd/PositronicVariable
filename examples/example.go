package main

import (
    "fmt"
    "github.com/hutchpd/positronic-variables/pkg/positronic"
)

func main() {
    // Example usage of the positronic variable
    antival := positronic.NewPositronicVariable(42)
    fmt.Printf("Initial antival: %v\n", antival.CurrentState())
}
