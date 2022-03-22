package main

import (
	"fmt"
	"strings"
)

func main() {
	// SとRの合計文字数
	var n int
	fmt.Scan(&n)

	// SとRの文字列
	var s string
	fmt.Scan(&s)
	ar := strings.Split(s, "")

	x := 0
	x_dir := 1
	y := 0
	y_dir := 1
	r_c := 0
	for i := 0; i < n; i++ {
		// Rの加算
		if ar[i] == "R" {
			r_c++
		}
		// 向きの調整
		if ar[i] == "R" && r_c%2 == 1 {
			y_dir = y_dir * -1
		}
		if ar[i] == "R" && r_c%2 == 0 {
			x_dir = x_dir * -1
		}

		// 実際の移動
		if ar[i] == "S" && r_c%2 == 1 {
			y = y + y_dir
		}
		if ar[i] == "S" && r_c%2 == 0 {
			x = x + x_dir
		}
	}

	fmt.Println(x, y)
}

// 問題：https://atcoder.jp/contests/abc244/tasks/abc244_b
// 4
// SSRS
// 2 -1が出力される
// 一応正解した
