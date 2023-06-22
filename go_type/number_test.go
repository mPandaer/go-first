package go_type

import (
	"fmt"
	"testing"
)

func TestGoType(t *testing.T) {
	f1 := float32(16777216.0)
	f2 := float32(16777216.0)
	fmt.Println(f1 == f2)
	fmt.Printf("f1: %X\n", f1)
	fmt.Printf("f2: %X\n", f2)
}
