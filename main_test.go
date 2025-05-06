package main

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	var a = 0
	_ = 10 / a
}
