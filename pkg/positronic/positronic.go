package positronic

import (
	"fmt"
	"sync"
	"github.com/hutchpd/QuantumSuperPosition-Go/pkg/quantum"
)

// PositronicVariable represents a variable with timelines and convergence tracking
type PositronicVariable struct {
	timeline    [][]interface{}
	convergence bool
	mu          sync.Mutex
}

// NewPositronicVariable initializes a new positronic variable
func NewPositronicVariable(initialValue interface{}) *PositronicVariable {
	return &PositronicVariable{
		timeline: [][]interface{}{{initialValue}},
	}
}

// Reinitialize resets the positronic variable to the initial value
func (pv *PositronicVariable) Reinitialize(value interface{}) {
	pv.mu.Lock()
	defer pv.mu.Unlock()

	// Reset the timeline to contain only the initial value
	pv.timeline = [][]interface{}{{value}}
	pv.convergence = false
}

// Assign assigns a new value to the positronic variable
func (pv *PositronicVariable) Assign(value interface{}, entropy int) {
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
func (pv *PositronicVariable) RunProgram(program func(*PositronicVariable, int)) {
    entropy := 1 // Start with forward time
    maxIterations := 100 // Prevent infinite loops

    // Reinitialize before starting
    pv.Reinitialize(pv.timeline[0][0])

    for iterations := 0; iterations < maxIterations; iterations++ {
        // Run the program, passing the current entropy
        program(pv, entropy)

        // Check for convergence after backward run
        if entropy < 0 && pv.checkConvergence() {
            pv.convergence = true
            pv.createSuperpositions()
            break // Convergence achieved
        }

        // Reverse the arrow of time
        entropy = -entropy
    }
}

// checkConvergence checks if the timelines have converged
func (pv *PositronicVariable) checkConvergence() bool {
    pv.mu.Lock()
    defer pv.mu.Unlock()

    n := len(pv.timeline)
    maxCycleLength := 10 // Our definition of a reasonable maximum cycle length

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
            return true // Convergence detected
        }
    }
    return false // No convergence detected
}

// createSuperpositions creates superpositions when convergence is achieved
func (pv *PositronicVariable) createSuperpositions() {
	pv.mu.Lock()
	defer pv.mu.Unlock()

	// Collect all unique states from the entire timeline
	stateSet := make(map[interface{}]struct{})

	for _, tl := range pv.timeline {
		if len(tl) > 0 {
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

// String implements the Stringer interface for PositronicVariable
func (pv *PositronicVariable) String() string {
	currentState := pv.CurrentState()
	return fmt.Sprintf("%v", currentState)
}
