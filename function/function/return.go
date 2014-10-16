package function

type Childhood struct {
	Age int
}

func (c *Childhood) MyAge() int {
	return c.Age
}
