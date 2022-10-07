package demo

import (
	"fmt"
	"sync"
	"testing"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() any {
			return &User{}
		},
	}
	u1 := pool.Get().(*User)
	u1.ID = 12
	u1.Name = "tom"
	// do
	// 放回去之前要先重置
	u1.Reset()
	pool.Put(u1)

	u2 := pool.Get().(*User)
	fmt.Println(u2)
}

type User struct {
	ID   uint64
	Name string
}

func (u *User) Reset() {
	u.ID = 0
	u.Name = ""
}
