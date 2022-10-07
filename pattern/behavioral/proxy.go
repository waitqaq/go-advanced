package behavioral

import "fmt"

// 代理模式
// 可以为另一个对象提供一个替身或占位符，以控制对这个对象的访问

type Seller interface {
	sell(name string)
}

// Station 火车站
type Station struct {
	stock int // 库存
}

func (station *Station) sell(name string) {
	if station.stock > 0 {
		station.stock--
		fmt.Printf("代理点中： %s 买了一张票，剩余：%d \n", name, station.stock)
	} else {
		fmt.Println("票已售空")
	}
}

type StationProxy struct {
	station *Station // 持有一个火车站对象
}

func (proxy *StationProxy) sell(name string) {
	if proxy.station.stock > 0 {
		proxy.station.stock--
		fmt.Printf("代理点中：%s 买了一张票，剩余：%d", name, proxy.station.stock)
	} else {
		fmt.Println("票已售空")
	}
}

// StationProxy 代理了 Station，代理类中持有被代理类对象，并且和代理类对象实现了同一接口
