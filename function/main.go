package main

import (
	"./function"
	"strconv"
)

func main() {
	function.Echo("Hello World")
	boo := function.Boo{}
	boo.Echo()

	function.RunCho()

	childhood := function.Childhood{Age: 10}
	function.Echo(strconv.Itoa(childhood.MyAge()))
}
