package demo_test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestParentCtx(t *testing.T) {
	ctx := context.Background()
	//dlCtx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Minute))
	childCtx := context.WithValue(ctx, "map", map[string]string{})
	ccChild := context.WithValue(childCtx, "key1", "value1")
	m := ccChild.Value("map").(map[string]string)
	m["key1"] = "val1"
	val := childCtx.Value("key1")
	fmt.Println(val)
	val = childCtx.Value("map")
	fmt.Println(val)
}

func TestContext(t *testing.T) {
	ctx := context.Background()
	valCtx := context.WithValue(ctx, "abc", 123)
	val := valCtx.Value("abc")
	fmt.Println(val)
	//timeoutCtx, cancel := context.WithTimeout(ctx, time.Second)
	//defer cancel()
	//dl, ok := timeoutCtx.Deadline()
	//err := timeoutCtx.Err()
	//fmt.Println(dl, ok)
}

func TestContext_timeout(t *testing.T) {
	bg := context.Background()
	timeoutCtx, cancel1 := context.WithTimeout(bg, time.Second)
	subCtx, cancel2 := context.WithTimeout(timeoutCtx, 3*time.Second)
	go func() {
		// 一秒后会过期，输出 timeout
		<-subCtx.Done()
		fmt.Printf("timeout")
	}()
	time.Sleep(2 * time.Second)
	cancel2()
	cancel1()
}

func TestBusinessTimeout(t *testing.T) {
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	end := make(chan struct{}, 1)
	go func() {
		MyBusiness()
		end <- struct{}{}
	}()
	ch := timeoutCtx.Done()
	select {
	case <-ch:
		fmt.Println("timeout")
	case <-end:
		fmt.Println("business end")
	}
}
func MyBusiness() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("hello world")
}
