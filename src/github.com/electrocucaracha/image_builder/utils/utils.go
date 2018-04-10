package utils

import (
	"log"
	"os"
)

func Check(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
