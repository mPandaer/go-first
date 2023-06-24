package function

import (
	"fmt"
	"net/http"
	"testing"
)

// 经典的函数声明
func sum(init int, other ...int) int {
	for _, v := range other {
		init += v
	}
	return init
}

func TestSumFunc(t *testing.T) {
	res := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(res)
}

// 函数参数的问题
func myAppend(s1 []int, elems ...int) []int {
	fmt.Printf("%T\n", elems)
	if len(elems) == 0 {
		fmt.Println("no elem")
		return s1
	}
	s1 = append(s1, elems...) //经典的纯函数设计
	return s1
}

func TestMyAppend(t *testing.T) {
	s1 := make([]int, 3, 6)
	s1[0] = 1
	s1[1] = 2
	s1[2] = 3
	s1 = myAppend(s1, 4, 5, 6)
	fmt.Println(s1)
}

// 函数的多返回值设计 + 具名返回值
func mulReturnFunc(slice []int) (length int, err error) {
	if len(slice) == 0 {
		err = fmt.Errorf("长度为0")
		return
	}
	length = len(slice)
	return
}

func TestMulReturn(t *testing.T) {
	length, err := mulReturnFunc([]int{1, 2, 3})
	fmt.Println(length, err)
}

//在go中函数是一等公民 可以被变量引用，作为返回值，有自己的类型

// 妙用1 利用类型的底层类型一致之间可以互相转换。可以将没有实现某类型接口的类型显示的转换为实现了某接口的相同类型
func greeting(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello,Gopher!\n")
}

func TestFuncType(t *testing.T) {
	http.ListenAndServe(":8080", http.HandlerFunc(greeting))
}

/**
//ListenAndServe的声明
func ListenAndServe(addr string, handler Handler) error {
	server := &Server{Addr: addr, Handler: handler}
	return server.ListenAndServe()
}


//HandlerFunc的定义
type HandlerFunc func(ResponseWriter, *Request)

// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}

*/

//妙用2 利用闭包简化计算

func times(x, y int) int {
	return x * y
}

func partialTimes(x int) func(int) int {
	return func(y int) int {
		return x * y
	}
}

func TestClosures(t *testing.T) {
	twoTimes := partialTimes(2)
	fiveTimes := partialTimes(5)

	fmt.Println("2 * 2 =", twoTimes(2))
	fmt.Println("5 * 2 =", fiveTimes(2))
}

// 利用函数是一等公民的特性实现AOP机制
func logging(function func(int) int, args ...int) func() {
	return func() {
		fmt.Println("函数调用前 args:", args)
		res := function(args[0])
		fmt.Println("函数调用后 res:", res)
	}
}

func TestAOP(t *testing.T) {
	add10 := func(init int) int {
		return init + 10
	}
	logging(add10, 5)()

}
