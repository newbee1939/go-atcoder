package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	var s string
	fmt.Scan(&s)
	ar := strings.Split(s, "")

	fmt.Println(ar[n-1])
}

// 問題:https://atcoder.jp/contests/abc244/tasks/abc244_a
// 英小文字からなる長さ N の文字列 S が与えられます。S の末尾の文字を出力してください。
// 5
// abcde
// eが出力される
