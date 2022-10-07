package factory

import "fmt"

// Person 与简单工厂的唯一区别是返回接口而不是结构体
// 在不公开内部实现的情况下，让调用者使用提供的各种功能
type Person interface {
	Greet()
}

// 定义一个不可导出的结构体 person
// 在通过 NewPersonAbstract 创建实例时返回接口，而不是结构体
type person struct {
	name string
	age  int
}

func (p person) Greet() {
	fmt.Printf("my name is %s", p.name)
}

func NewPersonAbstract(name string, age int) Person {
	return person{
		name: name,
		age:  age,
	}
}
