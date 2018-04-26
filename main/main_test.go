package main

import (
	"testing"
	"fmt"
)

func TestIntToHex(t *testing.T) {
	var num int64 = 1000
	byteArr := IntToHex(num)
	fmt.Println(byteArr)
	bb := []byte("1")
	fmt.Println(bb)
}
