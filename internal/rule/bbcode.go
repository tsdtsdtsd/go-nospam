package rule

import (
	"regexp"

	"github.com/tsdtsdtsd/go-nospam/internal/user"
)

// BBCode desfines the BBCode rule
type BBCode struct {
}

// Evaluate runs the rule and sends results to given channel
func (b *BBCode) Evaluate(c chan Result, ud *user.Data) {

	reg, err := regexp.Compile("(?is)\\[url[=\\]].*\\[\\/url\\]")
	if err != nil {
		c <- Result{0, err}
		return
	}

	if !reg.MatchString(ud.Text) {
		c <- Result{0, err}
		return
	}

	c <- Result{1, err}
}
