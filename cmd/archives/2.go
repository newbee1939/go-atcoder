package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)

	// 偶数の場合
	if (a*b)%2 == 0 {
		fmt.Println("Even")
	}
	// 奇数の場合
	if (a*b)%2 == 1 {
		fmt.Println("Odd")
	}
}

// 問題:https://atcoder.jp/contests/abs/tasks/abc086_a

// a 人を b 人ずつグループに分けると何グループできるか？
// (a + b - 1) / b
