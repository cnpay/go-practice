package test

import "fmt"
import "testing"

type Person struct {
	name string
	age  int
}

func (p Person) ToString() string {
	return fmt.Sprintf("p.name=%s,p.age=%d \n", p.name, p.age)
}

func Test_struct1(t *testing.T) {

	var liunix = Person{"liunix", 35}
	fmt.Printf("liunix.name=%s,liunix.age=%d \n", liunix.name, liunix.age)

	println(liunix.ToString())
}
