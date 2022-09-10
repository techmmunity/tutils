package main

import (
	"log"

	tutils "github.com/techmmunity/tutils/pkg"
)

func main() {
	result, err := tutils.Find([]string{}, func(c string) bool {
		return c == "foo"
	})

	if err != nil {
		log.Print(err.Error())
	}

	log.Print(result)
}
