package main

import (
	"fmt"
	"strings"
)

func main() {
	var v string
	fmt.Scanf("%s", &v)

	ar := strings.Split(v, "")
	c := 0
	for _, s := range ar {
		if s == "1" {
			c++
		}
	}

	fmt.Println(c)
}

// 問題:https://atcoder.jp/contests/abs/tasks/abc081_a
// 1の数を出力する
// 101 -> 2
