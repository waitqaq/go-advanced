package structural

import "fmt"

// 模版模式
// 定义一个操作中算法的骨架，而将一些步骤延迟到子类中，
// 这种方法让子类在不改变一个算法结构的情况下，能重新定义该算法的某些特定步骤

// 也就是将一个类中能够公共使用的方法防止在抽象类中实现，将不能公共使用的方法作为抽象方法，强制子类去实现

type Cooker interface {
	fire()
	cooke()
	outFire()
}

// CookMenu 类似于一个抽象类
type CookMenu struct {
}

func (CookMenu) fire() {
	fmt.Println("开火")
}

// 做菜，交给具体的子类实现
func (CookMenu) cooke() {

}

func (CookMenu) outFire() {
	fmt.Println("关火")
}

// 封装具体步骤
func doCook(cook Cooker) {
	cook.fire()
	cook.cooke()
	cook.outFire()
}

type XiHongShi struct {
	CookMenu
}

func (*XiHongShi) cooke() {
	fmt.Println("做西红柿")
}

type ChaoJiDan struct {
	CookMenu
}

func (*ChaoJiDan) cooke() {
	fmt.Println("做炒鸡蛋")
}
