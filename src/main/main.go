package main

import (
	"os"

	webapid "github.com/bitwormhole/web-api-descriptor"
	"github.com/starter-go/starter"
)

func main() {
	a := os.Args
	m := webapid.Module()
	i := starter.Init(a)
	i.MainModule(m)
	i.WithPanic(true).Run()
}
