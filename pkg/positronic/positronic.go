package positronic

import (
	"fmt"
	"sync"
	"github.com/hutchpd/QuantumSuperPosition-Go/pkg/quantum"
)

// OutputEntry represents a single timeline of a positronic variable
type OutputEntry struct {
    format string
    args   []interface{}
}

// PositronicVariable represents a variable with timelines and convergence tracking
type PositronicVariable struct {
	timeline    [][]interface{}
	convergence bool
	mu          sync.Mutex
	outputLogs [][]OutputEntry 
	iteration   int
}

// NewPositronicVariable initializes a new positronic variable
func NewPositronicVariable(initialValue interface{}) *PositronicVariable {
	return &PositronicVariable{
		timeline:   [][]interface{}{{initialValue}},
		outputLogs: [][]OutputEntry{},
	}
}

// Reinitialize resets the positronic variable to the initial value
func (pv *PositronicVariable) Reinitialize(value interface{}) {
	pv.mu.Lock()
	defer pv.mu.Unlock()

	// Reset the timeline to contain only the initial value
	pv.timeline = [][]interface{}{{value}}
	pv.convergence = false
	pv.outputLogs = [][]OutputEntry{}
}

// Assign assigns a new value to the positronic variable
func (pv *PositronicVariable) Output(format string, args ...interface{}) {
    pv.mu.Lock()
    defer pv.mu.Unlock()

    // Skip logging if this is the first iteration (iteration == 0)
    if pv.iteration == 0 {
        return
    }

    if len(pv.outputLogs) == 0 {
        pv.outputLogs = [][]OutputEntry{{}}
    }

    // Append the output to the current iteration's log
    pv.outputLogs[len(pv.outputLogs)-1] = append(
        pv.outputLogs[len(pv.outputLogs)-1],
        OutputEntry{format: format, args: args},
    )
}


// Assign assigns a new value to the positronic variable
func (pv *PositronicVariable) Assign(value interface{}) {
	pv.mu.Lock()
	defer pv.mu.Unlock()

	// Always append the new value to the timeline
	pv.timeline = append(pv.timeline, []interface{}{value})
}

// CurrentState returns the current state of the positronic variable
func (pv *PositronicVariable) CurrentState() interface{} {
	pv.mu.Lock()
	defer pv.mu.Unlock()

	if len(pv.timeline) == 0 {
		return nil
	}

	currentTimeline := pv.timeline[len(pv.timeline)-1]
	if len(currentTimeline) == 0 {
		return nil
	}

	return currentTimeline[0]
}

// RunProgram runs the program function over time loops until convergence
func (pv *PositronicVariable) RunProgram(program func(*PositronicVariable)) {
    entropy := 1         // Start with forward time
    maxIterations := 100 // Prevent infinite loops
    var cycleLen int     // To store the detected cycle length

    // Reinitialize before starting
    pv.Reinitialize(pv.timeline[0][0])

    for iterations := 0; iterations < maxIterations; iterations++ {
        pv.mu.Lock()
        pv.iteration = iterations // Set iteration count
        // Start a new output log entry for this iteration
        pv.outputLogs = append(pv.outputLogs, []OutputEntry{})
        pv.mu.Unlock()

        // Run the program, passing the current entropy
        program(pv)

        // Check for convergence after backward run
        if entropy < 0 {
            converged, cl := pv.checkConvergence()
            if converged {
                pv.convergence = true
                pv.createSuperpositions()
                cycleLen = cl
                break // Convergence achieved
            }
        }

        // Reverse the arrow of time
        entropy = -entropy
    }

    // After convergence, process the outputs
    pv.processOutputs(cycleLen)
}

// checkConvergence checks if the timelines have converged
func (pv *PositronicVariable) checkConvergence() (bool, int) {
	pv.mu.Lock()
	defer pv.mu.Unlock()

	n := len(pv.timeline)
	maxCycleLength := 10 // Define a reasonable maximum cycle length

	// Start checking from cycle length 1 up to maxCycleLength
	for cycleLen := 1; cycleLen <= maxCycleLength && cycleLen*2 <= n; cycleLen++ {
		match := true
		for i := 0; i < cycleLen; i++ {
			if pv.timeline[n-1-i][0] != pv.timeline[n-1-i-cycleLen][0] {
				match = false
				break
			}
		}
		if match {
			return true, cycleLen // Convergence detected, return cycle length
		}
	}
	return false, 0 // No convergence detected
}


// createSuperpositions creates superpositions when convergence is achieved
func (pv *PositronicVariable) createSuperpositions() {
    pv.mu.Lock()
    defer pv.mu.Unlock()

    // Collect all unique states from the timeline, skipping the initial value
    stateSet := make(map[interface{}]struct{})

    for i, tl := range pv.timeline {
        if len(tl) > 0 {
            if i == 0 {
                // Skip the initial value
                continue
            }
            stateSet[tl[0]] = struct{}{}
        }
    }

    // Convert the state set to a slice
    var states []interface{}
    for state := range stateSet {
        states = append(states, state)
    }

    // Create a superposition of all unique states
    superstate := quantum.Any(states...)

    // Append the superposition to the timeline
    pv.timeline = append(pv.timeline, []interface{}{superstate})
}


// processOutputs processes the collected outputs and prints the final result
func (pv *PositronicVariable) processOutputs(cycleLen int) {
    pv.mu.Lock()

    n := len(pv.outputLogs)
    if n < cycleLen {
        pv.mu.Unlock()
        return
    }

    // Collect outputs per format string
    outputMap := make(map[string][]interface{})

    // Use the outputs from the last cycle
    for i := n - cycleLen; i < n; i++ {
        logs := pv.outputLogs[i]
        for _, entry := range logs {
            key := entry.format
            argsList := outputMap[key]
            argsList = append(argsList, entry.args...)
            outputMap[key] = argsList
        }
    }

    // Prepare data to print after unlocking
    type printData struct {
        format   string
        superArg interface{}
    }
    var printList []printData

    // For each format string, create superposition of arguments
    for format, argsList := range outputMap {
        // Remove duplicates
        argsSet := make(map[interface{}]struct{})
        for _, arg := range argsList {
            argsSet[arg] = struct{}{}
        }

        var uniqueArgs []interface{}
        for arg := range argsSet {
            uniqueArgs = append(uniqueArgs, arg)
        }

        superArgs := quantum.Any(uniqueArgs...)

        // Collect the data to print
        printList = append(printList, printData{
            format:   format,
            superArg: superArgs,
        })
    }

    pv.mu.Unlock() // Release the mutex before printing

    // Now print the outputs
    for _, pd := range printList {
        fmt.Printf(pd.format, pd.superArg)
    }
}

// String implements the Stringer interface for PositronicVariable
func (pv *PositronicVariable) String() string {
	currentState := pv.CurrentState()
	return fmt.Sprintf("%v", currentState)
}
