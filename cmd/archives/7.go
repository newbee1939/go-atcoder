package main

import "fmt"

func main() {
	var A int //500円の枚数
	var B int //100円の枚数
	var C int //50円の枚数
	var X int //合計金額
	fmt.Scan(&A)
	fmt.Scan(&B)
	fmt.Scan(&C)
	fmt.Scan(&X)

	// 全検索
	total := 0
	count := 0
	for a := 0; a <= A; a++ {
		for b := 0; b <= B; b++ {
			for c := 0; c <= C; c++ {
				total = 500*a + 100*b + 50*c
				if total == X {
					count++
				}
			}
		}
	}

	fmt.Println(count)
}

// 問題：https://atcoder.jp/contests/abs/tasks/abc087_b
// 解答：https://qiita.com/drken/items/fd4e5e3630d0f5859067#5-%E9%81%8E%E5%8E%BB%E5%95%8F%E7%B2%BE%E9%81%B8-10-%E5%95%8F
// 全探索系の問題
// 変に数学的に解けないかを考えるのではなく、素直に全てのパターンを検索するのが大事
