package go_type

import (
	"fmt"
	"testing"
	"unsafe"
)

// 数组类型的等价性 类型 + 长度
func foo(arr [5]int) {}
func TestEquivalence(t *testing.T) {
	var arr1 [5]int
	//var arr2 [5]byte
	//var arr3 [4]int
	foo(arr1)
	//foo(arr2)
	//foo(arr3)
}

// 三种显式初始化的方式
func TestArrayInit(t *testing.T) {
	ipAddr := [3]string{"127.0.0.1", "localhost", "::1"}
	name := [...]string{"pandaer", "bobo"}
	nums := [...]int{
		9: 666,
	}
	fmt.Println(ipAddr)
	fmt.Println(name)
	fmt.Println(nums)

}

// 多维数据解析
func TestMultiDimensionalArrays(t *testing.T) {
	//var ipAddr [3][3][3]int
	//ipAddr[0] // [3][3]int
	//ipAddr[0][0] //[3]int
	//ipAddr[0][0][0] //int
}

// 数组底层实现 数组被看成一个整体，不是引用
func dataNature(arr [5]int, arrPtr *[5]int, arrSlice []int) {
	fmt.Println("arrSize: ", unsafe.Sizeof(arr))
	fmt.Println("arrPtrSize: ", unsafe.Sizeof(arrPtr))
	fmt.Println("arrSliceSize: ", unsafe.Sizeof(arrSlice))
}
func TestArrSize(t *testing.T) {
	arr := [5]int{1, 2, 3, 10, 15}
	arrPtr := &arr
	arrSlice := arr[:]
	dataNature(arr, arrPtr, arrSlice)
	//arrSize:  40
	//arrPtrSize:  8
	//arrSliceSize:  24
}

// 切片 Go为我们实现的可变数组
func TestBaseUse(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("len: ", len(slice))
	slice = append(slice, 7)
	fmt.Println("len: ", len(slice))
}

// 切片的几种创建方式
func TestSliceNew(t *testing.T) {
	//根据make函数来创建
	slice1 := make([]int, 5, 10)
	fmt.Println(slice1)

	//基于数组创建一个切片 [low:high:max] len=high-low;cap=max-low
	arr := [...]int{1, 2, 3, 4, 5, 6}
	slice2 := arr[1:3:5]
	fmt.Println(slice2)
	slice2 = append(slice2, 100)
	slice2 = append(slice2, 101)
	//超出范围，另起炉灶
	slice2 = append(slice2, 102)
	slice2[0] = 666
	//---------------
	fmt.Println(slice2)
	fmt.Println(arr)

	//基于切片创建切片
	slice3 := slice2[0:2:4]
	fmt.Println(slice3)
	slice3 = append(slice3, 999)
	slice3 = append(slice3, 998)
	slice3 = append(slice3, 997)
	//slice3 = append(slice3, 888)
	slice3[0] = 555
	fmt.Println(slice2)
	fmt.Println(slice3)

}
