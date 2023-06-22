package control_structure

import (
	"fmt"
	"math/rand"
	"testing"
)

//switch分支结构

// 普通用法
func readByExt(ext string) {
	switch ext {
	case "json":
		fmt.Println("json file")
	case "jpg", "jpeg", "png", "gif":
		fmt.Println("image file")
	case "txt", "md":
		fmt.Println("text file")
	default:
		fmt.Println("unsupported file:", ext)
	}
}

func TestSwitchBase(t *testing.T) {
	ext := "json"
	readByExt(ext)
}

//表达式是布尔值的Switch方式 对标IF的多分支

func TestBoolSwitch(t *testing.T) {
	//完整的方式
	switch num := rand.Int(); true {
	case num >= 90:
		fmt.Println("优秀")
	case num > 80:
		fmt.Println("合格")
	case num > 70:
		fmt.Println("及格")
	default:
		fmt.Println("不及格")
	}

	//语法糖方式
	switch num := rand.Int(); {
	case num >= 90:
		fmt.Println("优秀")
	case num > 80:
		fmt.Println("合格")
	case num > 70:
		fmt.Println("及格")
	default:
		fmt.Println("不及格")
	}
}

// switch的表达式列表
func TestExpressions(t *testing.T) {
	ip := "1.1.1.1"
	switch ip {
	case "127.0.0.1", "localhost", "::1":
		fmt.Println("当前主机")
	case "192.168.0.1":
		fmt.Println("路由器")
	default:
		fmt.Println("其他主机")
	}
}

//fallthrough的使用和注意事项 不会对下一个case表达式求值

func switchExpr() int {
	fmt.Println("switchExpr()...")
	return 1
}

func case1() int {
	fmt.Println("case1()...")
	return 1
}

func case2() int {
	fmt.Println("case2()...")
	return 2
}

func TestFallThrough(t *testing.T) {
	switch switchExpr() {
	case case1():
		fmt.Println("exec case1")
		fallthrough
	case case2():
		fmt.Println("exec case2")
		fallthrough
	default:
		fmt.Println("exec default")
	}

	//switchExpr()...
	//case1()...
	//exec case1
	//exec case2
	//exec default
}

// 特有的type switch
func TestTypeSwitch(t *testing.T) {
	var i interface{} = 13.6
	switch i.(type) {
	case int:
		fmt.Println("int type")
	case string:
		fmt.Println("string type")
	case float32:
		fmt.Println("float32 type")
	case float64:
		fmt.Println("float64 type")
	default:
		fmt.Println("other type")
	}
}

// for switch select 都可以捕获break
func TestBreakSwitch(t *testing.T) {
loop:
	for i := 1; i < 10; i++ {
		fmt.Println(i)
		switch i % 2 {
		case 0:
			break loop
		case 1:
			//
		}
	}
}
