package factory

import "fmt"

// PersonSimple 简单工厂模式
// 确保创建的实例具有需要的参数，保证实例的方法可以按预期进行
type PersonSimple struct {
	Name string
	Age  int
}

func (p PersonSimple) Greet() {
	fmt.Printf("my name is : %s", p.Name)
}

func NewPersonSimple(name string, age int) *PersonSimple {
	return &PersonSimple{
		Name: name,
		Age:  age,
	}
}
