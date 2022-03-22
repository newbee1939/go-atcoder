package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	n, _ := strconv.Atoi(nextLine())
	ar := strings.Split(nextLine(), " ")

	c := []int{}
	for i := 0; i < n; i++ {
		t, _ := strconv.Atoi(ar[i])

		tmp := 0
		for t%2 == 0 {
			t /= 2
			tmp++
		}
		// 各値のカウント数を格納する
		c = append(c, tmp)
	}
	min := c[0]
	for _, i := range c {
		if i < min {
			min = i
		}
	}
	fmt.Println(min)
}

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

// 問題：https://atcoder.jp/contests/abs/tasks/abc081_b
