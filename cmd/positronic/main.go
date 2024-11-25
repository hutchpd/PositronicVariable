package main

import (
	"github.com/hutchpd/positronic-variables/pkg/positronic"
)

func Program(antival *positronic.PositronicVariable) {
	// Use the Output method instead of fmt.Printf
	antival.Output("The antival is %v\n", antival)

	// Perform arithmetic operations with the positronic variable
	val := (antival.CurrentState().(int) + 1) % 3

	// Output the computed value
	antival.Output("The value is %v\n", val)

	// Update the positronic variable with the new value, passing entropy
	antival.Assign(val)
}

func main() {
	// Create a new positronic variable
	antival := positronic.NewPositronicVariable(-1)

	// Run the program with the positronic variable
	antival.RunProgram(Program)
}

// Output:
// The final antival is any(0, 1, 2)
// The value is any(The value is 0, The value is 1, The value is 2)