package go_type

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

// 字节视角
func byteString() {
	var s = "中国人"
	fmt.Println("len: ", len(s))
	for i := 0; i < len(s); i++ {
		fmt.Printf("0x%x ", s[i])
	}
	fmt.Println()
}

// 字符视角
func runeString() {
	var s = "中国人"
	fmt.Println("len: ", utf8.RuneCountInString(s))
	for _, char := range s {
		fmt.Printf("0x%x ", char)
	}
	fmt.Println()
}

// 证明rune只不过是int32的别名
func aliasRune() {
	a := rune(2)
	b := int32(1)
	fmt.Println(a + b)
}

//字符的编解码

// 编码
func encodeRune(char rune) []byte {
	buf := make([]byte, 3)
	_ = utf8.EncodeRune(buf, char)
	return buf
}

// 解码
func decodeRune(buf []byte) rune {
	char, _ := utf8.DecodeRune(buf)
	return char
}

// 类型互转
func TestTypeConversion(t *testing.T) {
	name := "pandaer"
	nameBytes := []byte(name)
	nameRunes := []rune(name)
	fmt.Println(name)
	fmt.Println(nameBytes)
	fmt.Println(nameRunes)
	fmt.Println("----------")
	fmt.Println(string(nameRunes))
	fmt.Println(string(nameBytes))
}

func TestRune(t *testing.T) {
	a := '国'
	buf := encodeRune(a)
	fmt.Println("encode: ", buf)
	fmt.Println("decode: ", string(decodeRune(buf)))
	//strings.Builder{}
	//a bool()
}

func TestString(t *testing.T) {
	byteString()
	println()
	runeString()
	println()
	aliasRune()

}
