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
	var a, _ = strconv.Atoi(nextLine())

	var tmp = strings.Split(nextLine(), " ")
	var b, _ = strconv.Atoi(tmp[0])
	var c, _ = strconv.Atoi(tmp[1])

	var s = nextLine()

	fmt.Println(a+b+c, s)
}

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

// 【メモ】
// 問題：https://atcoder.jp/contests/abs/tasks/practice_1
// a
// b c
// s
// a+b+c と s を空白区切りで 1 行に出力せよ。

// 【回答例】
// func main() {
// 	var a, b, c int
// 	var s string
// 	fmt.Scanf("%d", &a)
// 	fmt.Scanf("%d %d", &b, &c)
// 	fmt.Scanf("%s", &s)
// 	fmt.Printf("%d %s\n", a+b+c, s)
// }
