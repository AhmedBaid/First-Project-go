package functions

import (
	"fmt"
	"strconv"
)

// hexadecimal to decimal
func HextoByte(s string) string {
	decimal, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		fmt.Println(err)
		return s
	}
	return strconv.Itoa(int(decimal))
}
