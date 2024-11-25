
# Positronic Variables in Go

This module implements positronic variables that simulate variables moving backwards in time. It uses time loops and quantum superpositions to achieve convergence of variable states.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
  - [Example](#example)
- [Execution Instructions](#execution-instructions)
- [Expected Output](#expected-output)
- [Notes](#notes)
- [License](#license)

## Installation

To use this module, you need to have Go installed on your system. If you haven't installed Go yet, please download it from [the official website](https://golang.org/dl/) and follow the installation instructions.

You also need to get the `QuantumSuperPosition-Go` module, which provides quantum superposition functionality.

### Steps:

1. **Set Up Your Go Workspace**

   Ensure your `GOPATH` and `GOROOT` are set up correctly. You can check your Go environment by running:

   ```bash
   go env
   ```

2. **Get Dependencies**

   Install the `QuantumSuperPosition-Go` module (assuming it's hosted at `github.com/hutchpd/QuantumSuperPosition-Go`):

   ```bash
   go get github.com/hutchpd/QuantumSuperPosition-Go/pkg/quantum
   ```

## Usage

You can use the positronic variables module in your Go projects to simulate variables that can move backwards in time and converge to superpositions.

### Example

Below is an example of how to use the module.

**Project Structure**

```
your_project/
├── go.mod
├── main.go
└── positronic/
    └── positronic.go
```

**main.go**

```go
package main

import (
	"github.com/hutchpd/positronic-variables/pkg/positronic"
)

func Program(antival *positronic.PositronicVariable, entropy int) {
	// Use the Output method instead of fmt.Printf
	antival.Output("The antival is %v\n", antival)

	// Perform arithmetic operations with the positronic variable
	val := (antival.CurrentState().(int) + 1) % 3

	// Output the computed value
	antival.Output("The value is %v\n", val)

	// Update the positronic variable with the new value, passing entropy
	antival.Assign(val, entropy)
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

## Execution Instructions

1. **Ensure Go Is Installed**

   Make sure you have Go installed by running:

   ```bash
   go version
   ```

2. **Set Up the Project Structure**

   Create the following directory structure:

   ```
   your_project/
   ├── go.mod
   ├── main.go
   ├── positronic/
   │   └── positronic.go
   └── quantum/
       └── quantum.go
   ```

3. **Initialize the Go Module**

   In the root directory of your project (`your_project/`), run:

   ```bash
   go mod init your_project
   ```

   Then, run:

   ```bash
   go mod tidy
   ```

   This will download any necessary dependencies.

5. **Build and Run the Program**

   Navigate to the root directory of your project and execute the program by running:

   ```bash
   go run main.go
   ```

## Expected Output

When you run `main.go`, you should see the following output:

```
The antival is any([0 1 2])
The value is any([0 1 2])
```

## Notes

- **Adjusting Convergence Criteria:**

  You can adjust the convergence criteria in the `checkConvergence` method by changing the number of timelines compared.

- **Quantum Superposition Module:**

  Ensure that the `quantum` package correctly handles superpositions and implements the `String` method for proper output formatting.

- **Thread Safety:**

  The module uses a mutex (`sync.Mutex`) to ensure thread safety, allowing for potential concurrent use.

- **TODO:**

1. Redefine fprintf rather than use the ugly output function
2. avoid having to track the entropy in main.go, should be only ever tracked in positronic.go
3. remove the need for the assignment function, should instead override the assignment operator =

## License

This project is open-source as in the licence file.