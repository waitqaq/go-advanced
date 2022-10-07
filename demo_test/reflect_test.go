package demo

import (
	"fmt"
	"reflect"
	"testing"
)

func OutputFields(t *testing.T) {

}

func TestReflectPanic(t *testing.T) {
	typ := reflect.TypeOf(User1{})
	//typ.NumField()
	if typ.Kind() == reflect.Struct {
		fmt.Println("结构体")
	} else if typ.Kind() == reflect.Ptr {
		fmt.Println("指针")
	}
}

type User1 struct {
	Name string
}
