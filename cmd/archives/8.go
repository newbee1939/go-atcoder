package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	s := strings.Split(sc.Text(), " ")

	N, _ := strconv.Atoi(s[0])
	A, _ := strconv.Atoi(s[1])
	B, _ := strconv.Atoi(s[2])

	// 1以上N以下の整数のうち、各桁の和がA以上B以下であるものの総和を求める
	answer := 0
	for i := 1; i <= N; i++ {
		sum := getSum(i)
		if sum >= A && sum <= B {
			answer += i
		}
	}

	fmt.Println(answer)
}

func getSum(i int) int {
	sum := 0
	nums := strings.Split(strconv.Itoa(i), "")
	for _, num := range nums {
		num, _ := strconv.Atoi(num)
		sum += num
	}

	return sum
}

// 問題：https://atcoder.jp/contests/abs/tasks/abc083_b
// 解答：https://qiita.com/drken/items/fd4e5e3630d0f5859067#5-%E9%81%8E%E5%8E%BB%E5%95%8F%E7%B2%BE%E9%81%B8-10-%E5%95%8F
