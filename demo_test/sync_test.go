package demo

import (
	"fmt"
	"go/demo"
	"testing"
)

func TestDeferRLock(t *testing.T) {
	sm := demo.SafeMap[string, string]{
		Values: make(map[string]string, 4),
	}
	go func() {
		sm.LoadOrStore("a", "b")
	}()
	go func() {
		sm.LoadOrStore("a", "c")
	}()
	fmt.Println("hello")

}
