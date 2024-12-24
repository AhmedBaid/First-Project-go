package functions

import (
	"fmt"
	"strconv"
)

func HextoByte(s string) int {
	decimal, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return int(decimal)
}
