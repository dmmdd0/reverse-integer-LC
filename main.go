// https://leetcode.com/problems/reverse-integer/

package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println(reverse(123456789))

}

func reverse(x int) int {
	xString := ""
	sign := 1

	if x <= -2147483648 || x >= 2147483647 {
		return 0
	}

	if x < 0 {
		sign = -1
		x = -x
	}
	for i := 1; i <= x; i = i * 10 {
		xString += strconv.Itoa(x / i % 10)
	}

	res, _ := strconv.Atoi(xString)

	if res <= -2147483648 || res >= 2147483647 {
		return 0
	}

	return res * sign
}
