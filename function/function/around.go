package function

import (
	"fmt"
)

func Wrap(fn func()) {
	fmt.Println("BEGIN")
	fn()
	fmt.Println("AFTER")
}

func RunCho() {
	Wrap(func() {
		fmt.Println("Run Cho!")
	})
}
