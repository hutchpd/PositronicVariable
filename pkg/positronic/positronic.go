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

	if entropy > 0 {
		// Forward time: Append the new value to the timeline
		pv.timeline = append(pv.timeline, []interface{}{value})
	} else {
		// Backward time: Remove the last value from the timeline
		if len(pv.timeline) > 1 {
			pv.timeline = pv.timeline[:len(pv.timeline)-1]
		}
	}
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

	for iterations := 0; iterations < maxIterations; iterations++ {
		if entropy > 0 {
			// Forward time: Reinitialize and run the program
			pv.Reinitialize(pv.timeline[0][0])
		} else {
			// Backward time: Do not reinitialize
			if pv.convergence {
				// Timelines have converged; create superpositions
				pv.createSuperpositions()
				break // Convergence achieved
			}
		}

		// Run the program, passing the current entropy
		program(pv, entropy)

		// Check for convergence after backward run
		if entropy < 0 && pv.checkConvergence() {
			pv.convergence = true
		}

		// Reverse the arrow of time
		entropy = -entropy
	}
}

// checkConvergence checks if the timelines have converged
func (pv *PositronicVariable) checkConvergence() bool {
	pv.mu.Lock()
	defer pv.mu.Unlock()

	if len(pv.timeline) < 2 {
		return false
	}

	currTL := pv.timeline[len(pv.timeline)-1]
	prevTL := pv.timeline[len(pv.timeline)-2]

	// Ensure both timelines have the same length
	if len(currTL) != len(prevTL) {
		return false
	}

	// Check if states are equal across the last two timelines
	for i := range currTL {
		if currTL[i] != prevTL[i] {
			return false
		}
	}

	return true
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
