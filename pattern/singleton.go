package pattern

import "sync"

// 单例模式
type singleton struct {
}

var ins *singleton
var once sync.Once

// GetInsOr once.Do 确保 ins 实例全局只被创建一次，
//以及当同时多个创建动作时，只有一个创建动作被执行
func GetInsOr() *singleton {
	once.Do(func() {
		ins = &singleton{}
	})
	return ins
}
