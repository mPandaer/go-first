package control_structure

import (
	"fmt"
	"testing"
	"time"
)

// 经典for循环方式
func TestClassic(t *testing.T) {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
}

// 缺省for循环方式
func TestFor(t *testing.T) {
	//没有自用变量
	i := 10
	for ; i > 0; i-- {

	}
	fmt.Println(i)

	//没有后置表达式
	for j := 10; j > 0; {
		fmt.Println(j)
		j--
	}

	//没有前置和后置
	for true {
		i++
		if i > 3 {
			break
		}
	}

	//死循环
	for {
		break
	}
}

// 嵌套循环的label continue是否会执行内层循环的后置语句
func TestNesting(t *testing.T) {
	outer := 10
	inner := 10

outerLoop:
	for i := 0; i < outer; i++ {
		fmt.Println("i: ", i)
		for j := 0; j < inner; j++ {
			fmt.Println("j: ", j)
			if j%2 == 0 {
				continue outerLoop
			}
		}
	}
}

//for range 的三个问题

// 1. range 副本问题 存在
func TestCopybook(t *testing.T) {
	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int
	fmt.Println("original a = ", a)

	for i, v := range a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}

	fmt.Println("after r", r)
	fmt.Println("after a", a)
}

// 自用变量重用问题 存在
func TestVar(t *testing.T) {
	arr := [5]int{1, 2, 3, 4, 5}
	for i, v := range arr {
		go func(i int, v int) {
			time.Sleep(time.Second * 1)
			fmt.Printf("arr[%d]: %d\n", i, v)
		}(i, v) //解决方案
	}
	time.Sleep(time.Second * 10)
}

// map增加 删除问题 存在
func TestMap(t *testing.T) {
	var m = map[string]int{
		"tony": 21,
		"tom":  22,
		"jim":  23,
	}

	counter := 0
	for k, v := range m {
		if counter == 0 {
			delete(m, "tony")
		}
		counter++
		fmt.Println(k, v)
	}
	fmt.Println("counter is ", counter)
	counter = 0
	for k, v := range m {
		if counter == 0 {
			delete(m, "tony")
		}
		counter++
		fmt.Println(k, v)
	}
}
