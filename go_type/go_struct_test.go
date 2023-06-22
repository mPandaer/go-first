package go_type

import (
	"fmt"
	"testing"
	"unsafe"
)

//结构体类型

type M1 struct {
	cpu int32
	mem int64
}

type M2 = struct {
	cpu string
	gpu string
}

type T1 struct {
	*T2
}

type T2 struct {
	T1
}

func TestTypeDiff(t *testing.T) {
	fmt.Printf("%T\n", M1{})
	fmt.Printf("%T\n", M2{})
}

// 测试结构体导出字段和非导出字段
type T struct {
	F1 int
	F2 string
	f3 int
	F4 int
	F5 int
}

func TestStructExport(t *testing.T) {
	//var t1 = T{11,"hello",13,}
	var t2 = T{11, "hello", 13, 14, 15}
	//fmt.Println(t1)
	fmt.Println(t2)
}

// 测试结构体的内存布局
func TestStructMemLayout(t *testing.T) {
	m1 := M1{}
	m := M{}
	fmt.Println(unsafe.Sizeof(m1))
	fmt.Println(unsafe.Sizeof(m))
	//fmt.Println(unsafe.Sizeof(&m1))
}

//测试内存对齐的情况

type M struct {
	b byte
	i int64
	u uint16
}

type S struct {
	b byte
	u uint16
	i int64
}

func TestMemoryAlignment(t *testing.T) {
	m := M{}
	s := S{}
	fmt.Println(unsafe.Sizeof(m))
	fmt.Println(unsafe.Sizeof(s))
}
