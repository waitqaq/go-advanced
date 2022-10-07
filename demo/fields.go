package demo

import (
	"errors"
	"fmt"
	"reflect"
)

func IterateFields(val any) {
	res, err := iterateFields(val)
	if err != nil {
		fmt.Println(err)
	}
	for k, v := range res {
		fmt.Println(k, v)
	}
}

func iterateFields(val any) (map[string]any, error) {
	if val == nil {
		return nil, errors.New("不能为 nil")
	}
	// 获取 val 的类型和值
	typ := reflect.TypeOf(val)
	refVal := reflect.ValueOf(val)
	if typ.Kind() != reflect.Ptr && typ.Kind() != reflect.Struct {
		return nil, errors.New("非法类型")
	}
	// 如果是指针，拿到指针指向的对象
	for typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
		refVal = refVal.Elem()
		if typ.Kind() != reflect.Ptr {
			return nil, errors.New("非法指针")
		}
	}
	// 字段数目
	numField := typ.NumField()
	res := make(map[string]any, numField)
	for i := 0; i < numField; i++ {
		fdType := typ.Field(i)
		res[fdType.Name] = refVal.Field(i).Interface()
	}
	return res, nil
}

// SetField 只能修改指针的值
func SetField(entity any, field string, newVal any) error {
	val := reflect.ValueOf(entity)
	typ := val.Type()
	// 必须是一级指针 *User
	if typ.Kind() != reflect.Ptr || typ.Elem().Kind() != reflect.Struct {
		return errors.New("非法类型")
	}
	// 通过指针拿到结构体
	typ = typ.Elem()
	val = val.Elem()
	// 拿出值对应的字段名
	fd := val.FieldByName(field)
	// 判断值的字段名是否在类型的字段名中
	if _, found := typ.FieldByName(field); !found {
		return errors.New("字段不存在")
	}
	if !fd.CanSet() {
		return errors.New("不可修改字段")
	}
	fd.Set(reflect.ValueOf(newVal))
	return nil
}
