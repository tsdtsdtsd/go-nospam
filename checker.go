package nospam

import (
	"github.com/tsdtsdtsd/go-nospam/internal/rule"
	"github.com/tsdtsdtsd/go-nospam/internal/user"
)

// Checker checks for spam
type Checker struct {
	options *Options
}

// NewChecker returns a new Checker
func NewChecker() *Checker {
	return &Checker{}
}

// NewCheckerWithConfig returns a new Checker with given options set
func NewCheckerWithConfig(opt *Options) *Checker {
	return &Checker{opt}
}

// SetOptions overwrites the Checkers options with the given ones.
// It DOES NOT merge values.
func (c *Checker) SetOptions(opt *Options) {
	c.options = opt
}

// Evaluate runs all rules
func (c *Checker) Evaluate(ud *user.Data) (float32, error) {

	// List of rules
	rules := []rule.Rule{
		&rule.BBCode{},
	}

	// Create result channel and run all rules in routines
	resultChan := make(chan rule.Result, len(rules))
	for _, r := range rules {
		go r.Evaluate(resultChan, ud)
	}

	var total float32
	var r rule.Result

	// Fetch results
	for i := 0; i < len(rules); i++ {
		r = <-resultChan
		if r.Err != nil {
			return 0, r.Err
		}

		// For now, just accumulate scores
		total += r.Score
	}

	return total, nil
}
