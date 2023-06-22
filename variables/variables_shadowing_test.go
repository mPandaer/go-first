package variables

import (
	"errors"
	"fmt"
	"testing"
)

var a int = 2020

func checkYear() error {
	err := errors.New("wrong year")
	switch a, _ = getYear(); a {
	case 2020:
		fmt.Println("it is", a, err)
	case 2021:
		fmt.Println("it is", a)
		err = nil
	}
	fmt.Println("after check, it is", a)
	return err
}

// go_type new int //屏蔽了预定义标识符new

func getYear() (int, error) {
	var b int16 = 2021
	return int(b), nil
}

func TestShadowing(t *testing.T) {
	err := checkYear()

	if err != nil {
		fmt.Println("call checkYear error:", err)
		return
	}
	fmt.Println("call checkYear ok")
}
