// Package rule contains spam checker rules
package rule

import "github.com/tsdtsdtsd/go-nospam/internal/user"

// Rule describes an evaluatable rule
type Rule interface {
	Evaluate(c chan Result, ud *user.Data)
}

// Result is the result of a rules evaluation
type Result struct {
	Score float32
	Err   error
}
