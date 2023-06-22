package go_type

import (
	"fmt"
	"testing"
)

//map key的限制 key必须是元素可比较的 == !=运算符支持

// map不支持零值可用，必须显式初始化
func TestMapInit(t *testing.T) {
	var m1 = map[int]string{}
	m1[1] = "pandaer"
	DPrint(m1)
	m2 := map[int]string{
		1: "pandaer",
	}
	DPrint(m2)
	//复杂结构体的优化
	m3 := map[Point]string{
		Point{X: 121, Y: 128}: "北京",
	}
	DPrint(m3)
	m4 := map[Point]string{
		{22, 33}: "北京",
	}
	DPrint(m4)

	//通过make
	//m5 := make(map[Point]string)
	//m6 := make(map[Point]string,10) 指定容量适合 知道大概得容量的场景

}

type Point struct {
	X int
	Y int
}

func DPrint(origin interface{}) {
	fmt.Println(origin)
	fmt.Println("-------------------------")
}

// map的基本使用
func TestMapBaseUser(t *testing.T) {
	m := make(map[string]int)
	//插入数据
	m["key1"] = 10
	//获取键值对数量
	fmt.Println("len: ", len(m))
	//通过key读取value
	//v1 := m["key1"] //这种方式无法判断是否真的读取到零值了，没有读取到默认类型零值
	//v2 := m["key2"]
	//另外一种方式
	v2, ok := m["key2"]
	fmt.Println(ok)
	fmt.Println(v2)

	//删除
	fmt.Println(m)
	delete(m, "key1")
	fmt.Println(m)

	//map的遍历
	m["key1"] = 1
	m["key2"] = 2
	m["key3"] = 3
	for k, v := range m {
		fmt.Println("[", k, "]", ": ", v)
	}

	//验证遍历方式是否是迭代器 否 底层是利用hash的机制
	for i := 0; i < 10; i++ {
		for k, v := range m {
			fmt.Println("[", k, "]", ": ", v)
		}
	}
}
