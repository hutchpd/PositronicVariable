package positronic_test

import (
    "testing"
    "github.com/hutchpd/positronic-variables/pkg/positronic"
)

func TestNewPositronicVariable(t *testing.T) {
    pv := positronic.NewPositronicVariable(0)
    if pv.CurrentState() != 0 {
        t.Errorf("Expected initial state to be 0, got %v", pv.CurrentState())
    }
}

func TestReverseArrowOfTime(t *testing.T) {
    entropy := 1
    positronic.ReverseArrowOfTime(&entropy)
    if entropy != -1 {
        t.Errorf("Expected entropy to be -1, got %d", entropy)
    }
}
