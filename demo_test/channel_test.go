package demo

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	ch := make(chan string, 4)
	go func() {
		str := <-ch
		fmt.Println(str)
	}()
	go func() {
		str := <-ch
		fmt.Println(str)
	}()
	go func() {
		str := <-ch
		fmt.Println(str)
	}()

	ch <- "hello"
	ch <- "hello"
	time.Sleep(time.Second)
}

func TestBroker(t *testing.T) {
	b := &Broker{
		consumers: make([]*Consumer, 0, 10),
	}
	c1 := &Consumer{
		ch: make(chan string, 1),
	}
	c2 := &Consumer{
		ch: make(chan string, 1),
	}

	b.Subscribe(c1)
	b.Subscribe(c2)

	b.Produce("hello")
	fmt.Println(<-c1.ch)
	fmt.Println(<-c2.ch)
}

// 每一个消费者订阅时，创建一个子 channel

type Consumer struct {
	ch chan string
}

type Broker struct {
	consumers []*Consumer
}

func (b *Broker) Produce(msg string) {
	for _, c := range b.consumers {
		c.ch <- msg
	}
}

func (b *Broker) Subscribe(c *Consumer) {
	b.consumers = append(b.consumers, c)

}

// 轮训所有消费者

type Broker1 struct {
	ch        chan string
	consumers []func(s string)
}

func (b *Broker1) Produce(msg string) {
	b.ch <- msg
}

func (b *Broker1) Subscribe(consume func(s string)) {
	b.consumers = append(b.consumers, consume)

}

func (b *Broker1) Start() {
	go func() {
		s := <-b.ch
		for _, c := range b.consumers {
			c(s)
		}
	}()
}
