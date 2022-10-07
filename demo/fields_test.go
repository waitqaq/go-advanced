package demo

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIterateFields(t *testing.T) {

	u1 := &User1{
		Name: "xia",
	}
	u2 := &u1

	tests := []struct {
		// 名字(测试用例叫什么)
		name string
		// 输入
		val any
		// 输出部分
		wantRes map[string]any
		wantErr error
	}{
		{
			name:    "nil",
			val:     nil,
			wantErr: errors.New("不能为 nil"),
		},
		{
			name:    "user",
			val:     User1{Name: "tom"},
			wantErr: nil,
			wantRes: map[string]any{
				"Name": "tom",
			},
		},
		{
			// 指针
			name: "pointer",
			val:  &User1{Name: "jerry"},
			// 要支持指针
			wantErr: nil,
			wantRes: map[string]any{
				"Name": "jerry",
			},
		},
		{
			// 多重指针
			name:    "pointer",
			val:     u2,
			wantErr: nil,
			wantRes: map[string]any{
				"Name": "xia",
			},
		},
		{
			// 非法输入
			name:    "slice",
			val:     []string{},
			wantErr: errors.New("非法类型"),
		},
		{
			// 非法指针输入
			name:    "pointer to map",
			val:     &(map[string]string{}),
			wantErr: errors.New("非法指针"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := iterateFields(tt.val)
			// 断言错误
			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				return
			}
			// 断言结果
			assert.Equal(t, tt.wantRes, res)
		})
	}
}

type User1 struct {
	Name string
}

type User2 struct {
	Name string
	age  int
}

func TestSetField(t *testing.T) {
	tests := []struct {
		// 用例名
		name string
		// 入参
		entity any
		field  string
		newVal any
		// 出参
		wantErr error
	}{
		{
			name:    "struct",
			entity:  User2{},
			field:   "Name",
			wantErr: errors.New("非法类型"),
		},
		{
			name:    "private field",
			entity:  &User2{},
			field:   "age",
			wantErr: errors.New("不可修改字段"),
		},
		{
			name:    "invalid field",
			entity:  &User2{},
			field:   "invalid_field",
			wantErr: errors.New("字段不存在"),
		},
		{
			name:    "pass",
			entity:  &User2{},
			field:   "Name",
			newVal:  "tom",
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := SetField(tt.entity, tt.field, tt.newVal)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
