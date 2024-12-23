package main

import (
	"fmt"
	"strconv"
)

func HextoByte(s string)  {
	decimal, err := strconv.ParseInt(s, 16, 32)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("valeur byte est :", decimal)
}
