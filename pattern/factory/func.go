package factory

import "fmt"

type PersonFunc struct {
	name string
	age  int
}

func NewPersonFactory(age int) func(name string) PersonFunc {
	return func(name string) PersonFunc {
		return PersonFunc{
			name: name,
			age:  age,
		}
	}
}

func main() {
	newPerson := NewPersonFactory(18)
	person := newPerson("xia")
	fmt.Println(person)
}
