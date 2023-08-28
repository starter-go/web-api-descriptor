package main

import (
	"os"

	"github.com/starter-go/starter"
)

func main() {
	args := os.Args
	starter.Configen(args)
}
