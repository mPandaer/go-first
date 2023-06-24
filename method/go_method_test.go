package method

import (
	"fmt"
	"io"
	"reflect"
	"testing"
	"time"
)

//Go方法 receiver 的一些限制

//1. 基类型不能为接口和指针类型 *T/T 基类型都是T

type MyInt *int

//Invalid receiver type 'MyInt' ('MyInt' is a pointer type)
//func (myInt *MyInt) add(other MyInt) MyInt{
//	return *myInt + other
//}

type MyReader io.Reader

//Invalid receiver type 'MyReader' ('MyReader' is an interface type)
//func (r MyReader) test() {
//
//}

//方法和类型必须在同一包下
//所以不能给原生类型增加方法
//不能夸包给其他类型增加方法

//方法的本质

type T struct {
	num int32
}

func (t T) Get() int32 {
	return t.num
}

func (t *T) Set(other int32) {
	t.num = other
}

func TestMethodBasic(t *testing.T) {
	my := T{}
	num := my.Get()
	println(num)
	my.Set(22)
	fmt.Printf("%T\n", my)
	println(my.Get())
}

func TestMethodBasic1(t *testing.T) {
	my := T{}
	num := T.Get(my)
	println(num)
	(*T).Set(&my, 22)
	fmt.Printf("%T\n", my)
	println(my.Get())
}

//一个坑

type field struct {
	name string
}

func (f field) print() {
	fmt.Println(f.name)
}

func TestBug(t *testing.T) {
	data1 := []*field{{"one"}, {"two"}, {"three"}}
	for _, v := range data1 {
		//go v.print()
		time.Sleep(1 * time.Second)
		go field.print(*v)
	}

	data2 := []field{{"four"}, {"five"}, {"six"}}
	for _, v := range data2 {
		time.Sleep(1 * time.Second)
		go field.print(v)
	}

	time.Sleep(3 * time.Second)

}

// 方法集合
func getMethodSet(t interface{}) {
	dynType := reflect.TypeOf(t)

	if dynType == nil {
		fmt.Println("there is no dynamic type")
		return
	}

	n := dynType.NumMethod()

	if n == 0 {
		fmt.Printf("%s's method is empty\n", dynType)
		return
	}

	fmt.Printf("%s's method set:\n", dynType)

	for j := 0; j < n; j++ {
		fmt.Println("-", dynType.Method(j).Name)
	}
	fmt.Println()
}

type MM struct{}

func (MM) M1() {}
func (MM) M2() {}

type SS MM

func TestMethodSet(t *testing.T) {
	var s SS
	getMethodSet(MM(s))
}

type IT struct {
	I
}

type I interface {
	M1()
}

type T1 int

type t2 struct {
	n int
	m int
}

type S1 struct {
	T1
	*t2
	I
	a int
	b string
}

type S2 struct {
	T1 T1
	t2 *t2
	I  I
	a  int
	b  string
}

func TestS1S2(t *testing.T) {
	var s1 S1
	var s2 S2
	getMethodSet(s1)
	println("--------------")
	getMethodSet(s2)
}
