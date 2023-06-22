package constants

import (
	"fmt"
	"testing"
)

//	var m1 = go_type.M3{
//		cpu:12,
//		mem:23,
//	}
//
// 普通的常量
const LOCALHOST string = "127.0.0.1"

// 无类型常量
const PI = 3.14

// 常量块 => 用作枚举
const (
	_ = iota //和常量块绑定
	IPV4
	IPV6
)

const (
	_ = iota
	HTTP
	HTTPS
	WEBSOCKET
	TCP
	UDP
	IP
)

func TestConst(t *testing.T) {
	fmt.Println(LOCALHOST)
	fmt.Println(PI)
	fmt.Println(IPV4)
	fmt.Println(IPV6)
	fmt.Println("-------")
	fmt.Println(HTTP)
	fmt.Println(HTTPS)
	fmt.Println(TCP)
	fmt.Println(UDP)
	fmt.Println(WEBSOCKET)
	fmt.Println(IP)
}
