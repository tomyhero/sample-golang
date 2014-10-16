package main

import (
	"./function"
)

func main() {
	function.Echo("Hello World")
	boo := function.Boo{}
	boo.Echo()

	function.RunCho()
}
