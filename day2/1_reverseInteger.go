package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	//fmt.Println(reverse(123))
	//fmt.Println(reverse(-123))
	//fmt.Println(reverse(120))
	//fmt.Println(reverse(0))
	//fmt.Println(reverse(1534236469))

	println(quotientReverse(123))
	println(quotientReverse(-123))
	println(quotientReverse(120))
	println(quotientReverse(0))
	println(quotientReverse(1534236469))

}

/*題目描述為給定一個 32 位元帶有正負符號的整數，要求返回他的反轉數。
題目有補充說明可接受的數值範圍為 : [−2³¹, 2³¹ − 1]，反轉後的數不在此範圍內則返回0。

*/

func reverse(x int) int {

	sign := 1
	if x < 0 {
		sign = -1
		x = -x
	}

	s := strconv.Itoa(x)
	bs := []byte(s)

	head := 0
	tail := len(bs) - 1

	for head < tail {
		bs[head], bs[tail] = bs[tail], bs[head]
		head++
		tail--
	}

	rs := string(bs)

	if sign == -1 {
		rs = "-" + rs
	}

	y, err := strconv.Atoi(rs)
	if err != nil {
		fmt.Println(err, rs)
	}

	if (y > math.MaxInt32) || (y < math.MinInt32) {
		return 0
	}

	return y
}
