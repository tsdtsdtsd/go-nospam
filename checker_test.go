package nospam_test

import (
	"testing"

	"github.com/tsdtsdtsd/go-nospam"
	"github.com/tsdtsdtsd/go-nospam/internal/user"
)

type userDataCheck struct {
	Description string
	Data        user.Data
	ShouldScore float32
}

var userDataChecks = []userDataCheck{

	userDataCheck{
		Description: "BBCode spam",
		ShouldScore: 1,
		Data: user.Data{
			Name:  "Peter Griffin",
			Email: "peter@griffinmail.fun",
			Text:  "[url=www.thingy156.net]thing[/url]",
			IP:    "192.168.10.15",
		},
	},
	userDataCheck{
		Description: "Valid comment",
		ShouldScore: 0,
		Data: user.Data{
			Name:  "Peter Griffin",
			Email: "peter@griffinmail.fun",
			Text:  "Hehehehehehehehe",
			IP:    "192.168.10.15",
		},
	},
}

func TestDev(t *testing.T) {

	checker := nospam.NewCheckerWithConfig(
		&nospam.Options{},
	)

	for _, check := range userDataChecks {
		rating, err := checker.Evaluate(&check.Data)
		if err != nil {
			t.Error(err)
		}

		if rating != check.ShouldScore {
			t.Errorf("%s: rating should be %f but is %f", check.Description, check.ShouldScore, rating)
		}
	}

}

var (
	rating float32
)

func BenchmarkCheckerEvaluate(b *testing.B) {

	checker := nospam.NewCheckerWithConfig(
		&nospam.Options{},
	)

	for i := 0; i < b.N; i++ {
		var err error
		rating, err = checker.Evaluate(&userDataChecks[0].Data)

		if err != nil {
			b.Error(err)
		}
	}
}
