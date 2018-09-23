package main

import (
	"fmt"

	"github.com/tsdtsdtsd/go-nospam"
	"github.com/tsdtsdtsd/go-nospam/internal/user"
)

func main() {

	checker := nospam.NewCheckerWithConfig(
		&nospam.Options{},
	)

	rating, err := checker.Evaluate(
		&user.Data{
			Name:  "Peter Griffin",
			Email: "peter@griffinmail.fun",
			Text:  "Awesome! [url=www.thingy156.net]thing[/url]",
			IP:    "192.168.10.15",
		},
	)

	fmt.Println(rating, err)
}
