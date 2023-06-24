package handle_error

import (
	"errors"
	"fmt"
	"testing"
)

type My struct {
	name string
	age  int
}

// 测试类型断言
func TestTypeAssertion(t *testing.T) {
	var my interface{} = My{
		name: "pandaer",
		age:  19,
	}
	// myType, ok := my.(My) 类型断言
	if myType, ok := my.(My); ok {
		fmt.Println("yes", myType)
		return
	}
	fmt.Println("no")

}

// 透明处理策略 - 只关心有没有错误，不关系到底是什么错误
func throwError(b bool) error {
	if b {
		return errors.New("have a error")
	}
	return nil
}

func TestErrorHandle(t *testing.T) {
	if err := throwError(true); err != nil {
		fmt.Println("存在错误")
		return
	}
	fmt.Println("没有错误")
}

//需要知道具体的错误，基于唯一的错误字符串进行错误判断 - 哨兵错误处理机制
/**
type error interface {
	Error() string
}
*/

func throwErrorWithInt(order int) error {
	switch order {
	case 1:
		return errors.New("1")
	case 2:
		return errors.New("2")
	case 3:
		return fmt.Errorf("fmt: %w", errors.New("3")) //包装类型的错误
	default:
		return nil
	}
}

// 哨兵错误处理
func TestSentinelHandler(t *testing.T) {
	err := throwErrorWithInt(3)
	switch err.Error() {
	case "1":
		fmt.Println("错误1")
	case "2":
		fmt.Println("错误2")
	case "3":
		fmt.Println("错误3")
	case "fmt: 3":
		fmt.Println("包装错误3")
	}
}

// 优化，降低耦合度
var (
	Err1 = errors.New("1")
	Err2 = errors.New("2")
	Err3 = errors.New("3")
)

func throwErrorWithInt1(order int) error {
	switch order {
	case 1:
		return Err1
	case 2:
		return Err2
	case 3:
		return fmt.Errorf("fmt: %w", Err3) //包装类型的错误
	default:
		return nil
	}
}

func TestSentinelHandler1(t *testing.T) {
	err := throwErrorWithInt1(3)
	switch err {
	case Err1:
		fmt.Println("错误1")
	case Err2:
		fmt.Println("错误2")
	case Err3:
		fmt.Println("错误3")
	default:
		fmt.Println("其他错误")
	}
}

// 再优化 go1.13 利用提供的Is函数可以递归包装错误 推荐方式
func TestSentinelHandler2(t *testing.T) {
	err := throwErrorWithInt1(3)
	switch {
	case errors.Is(err, Err1):
		fmt.Println("错误1")
	case errors.Is(err, Err2):
		fmt.Println("错误2")
	case errors.Is(err, Err3):
		fmt.Println("错误3")
	default:
		fmt.Println("其他错误")
	}
}

// 根据类型进行错误判断 As函数
type MyError struct {
	e string
}

func (e *MyError) Error() string {
	return e.e
}

func TestTypeHandle(t *testing.T) {
	var err = &MyError{"MyError demo"}
	err1 := fmt.Errorf("wrap err1:%w", err)
	err2 := fmt.Errorf("wrap err2:%w", err1)
	var e *MyError
	fmt.Println("init e:", e)
	if errors.As(err2, &e) {
		println("MyError is on the chain of err2")
		println(e == err)
		return
	}
	println("no")
	//http.Serve
}

//基于错误特征行为的错误处理 参考 net.Error

/**
函数的健壮性以及简洁性
*/

//健壮性 1.不相信任何输入 2.不忽略任何错误 3.不假定异常不会发生

//panic的妙用 -- RuntimeException

func zoo() {
	println("call zoo")
	println("exit zoo")
}

func bar() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("%T\n", e)
			//fmt.Println(e)
		}
	}()

	println("call bar")
	panic(1)
	zoo()
	println("exit bar")
}

func foo() {
	println("call foo")
	bar()
	println("exit foo")

}

func TestPanic(t *testing.T) {
	println("call main")
	foo()
	println("exit main")
}

// 测试Defer是放在栈中的
func deferStack() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
}

func TestDeferStack(t *testing.T) {
	deferStack()
}

// defer带来的性能影响
func sum(count int) int {
	for i := 1; i < count; i++ {
		count += i
	}
	return count
}

func sumWithDefer() {
	defer func() {
		sum(100)
	}()
}

func sumWithNoDefer() {
	sum(100)
}

func BenchmarkSumWithDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumWithDefer() //1422065542
	}
}

func BenchmarkSumWithNoDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumWithNoDefer() //1419190875
	}
}
