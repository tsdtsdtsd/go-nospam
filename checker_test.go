package nospam_test

import (
	"testing"

	"github.com/tsdtsdtsd/go-nospam"
	"github.com/tsdtsdtsd/go-nospam/internal/user"
)

func TestDev(t *testing.T) {

	checker := nospam.NewCheckerWithConfig(
		&nospam.Options{},
	)

	var expected float32 = 1

	ud := user.Data{
		Name:  "Peter Griffin",
		Email: "peter@gmail.com",
		Text:  "[url=www.thingy156.net]thing[/url]",
		IP:    "192.168.10.15",
	}

	rating, err := checker.Evaluate(&ud)
	if err != nil {
		t.Error(err)
	}

	if rating != expected {
		t.Errorf("rating should be %f but is %f", expected, rating)
	}
}

var (
	rating float32
)

func BenchmarkDev(b *testing.B) {

	checker := nospam.NewCheckerWithConfig(
		&nospam.Options{},
	)

	ud := user.Data{
		Name:  "Peter Griffin",
		Email: "peter@gmail.com",
		Text:  "[url=www.thingy156.net]thing[/url]",
		IP:    "192.168.10.15",
	}

	for i := 0; i < b.N; i++ {
		var err error
		rating, err = checker.Evaluate(&ud)

		if err != nil {
			b.Error(err)
		}
	}
}
