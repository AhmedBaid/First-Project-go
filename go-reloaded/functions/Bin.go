package functions

import (
	"fmt"
	"strconv"
)

// binaire to decimal
func Bin(s string) string {
	decimal, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		fmt.Println("error:", err)
		return s
	}
	return strconv.Itoa(int(decimal))
}
