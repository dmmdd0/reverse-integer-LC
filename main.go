// https://leetcode.com/problems/reverse-integer/

package main

import "fmt"

func main() {

	fmt.Println(reverse(123456789))

}

func reverse(x int) int {
	r := 0
	for i := 1; i < x; i = i * 10 {
		r += (x / i % 10) * i
		fmt.Println(i, x/i%10, (x/i%10)*i, r)
	}
	return x
}
