package control_structure

import (
	"fmt"
	"testing"
)

// 普通结构
func TestCommon(t *testing.T) {
	num := 100
	if num > 90 {
		fmt.Println("优秀")
	} else {
		fmt.Println("不好")
	}
}

// 声明自用变量+多分支结构
func TestMul(t *testing.T) {
	if a, c := 1, 2; false {
		println(a)
	} else if b := 3; false {
		println(a, b)
	} else {
		println(a, b, c)
	}
}
