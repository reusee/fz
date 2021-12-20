package main

import (
	"github.com/reusee/fz"
)

func run() (err error) {
	defer he(&err)

	scope := NewScope()

	var execute fz.Execute
	scope.Assign(&execute)
	ce(execute())

	return
}

func main() {
	run()
}
