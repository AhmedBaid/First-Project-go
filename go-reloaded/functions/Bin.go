package functions

import (
	"fmt"
	"strconv"
)
// binaire to decimal
func Bin(s string) int {
	decimal, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		fmt.Println("error:", err)
		return 0
	}
	return int(decimal)
}
