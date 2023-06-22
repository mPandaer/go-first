package variables

import (
	"fmt"
	"testing"
)

// 最基础的变量声明
var name string = "pandaer"

// 变量块声明
var (
	cpu   string  = "Intel13500h"
	mem   string  = "16G"
	price float32 = 4799.99
)

// 单行同类型声明
var computer, pad, phone string = "honor x 16pro", "huawei pad 11", "redmi k50"

// var computer1,pad1 ,phone1 string = "honor x 16pro","huawei pad 11","redmi k50"

// 变量块和变量行组合
var (
	book1, name1, price1 string = "book1", "name1", "price1"
	book2, name2, price2 string = "book2", "name2", "price2"
)

// go编译器的类型推断的变量声明语法糖
var id = "20230620"

// 结合变量行 实现同一行不同类型
var mac_cpu, mac_mem, mac_disk = "M1", 16, 512

func TestVariables(t *testing.T) {
	fmt.Println(name)
	fmt.Println(cpu)
	fmt.Println(mem)
	fmt.Println(price)
	fmt.Println(computer)
	fmt.Println(pad)
	fmt.Println(phone)

	fmt.Println(book1)
	fmt.Println(name1)
	fmt.Println(price1)
	fmt.Println(book2)
	fmt.Println(name2)
	fmt.Println(price2)

	fmt.Println(id)

	fmt.Println(mac_cpu)
	fmt.Println(mac_mem)
	fmt.Println(mac_disk)

}
