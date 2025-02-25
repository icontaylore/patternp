package obs1

import "fmt"

type ClientOne struct {
	Name string
}

func (c *ClientOne) Update(temp, humid, press float64) {
	fmt.Println(c.Name)
}
