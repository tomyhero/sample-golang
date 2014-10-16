package function

import (
	"fmt"
)

type Boo struct {
	Name string
}

func (b *Boo) Echo() {
	fmt.Println("Boo")
}
