package structural

import (
	"fmt"
	"testing"
)

func TestTemplate(t *testing.T) {
	xiHongShi := &XiHongShi{}
	doCook(xiHongShi)
	fmt.Println("the other")
	chaoJiDan := &ChaoJiDan{}
	doCook(chaoJiDan)
}
