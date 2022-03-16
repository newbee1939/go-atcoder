// 標準入力の例（一行ずつ読み込む）（冗長さを無くした形）
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
	for i := 0; i < n; i++ {
		var s = strings.Split(nextLine(), " ")
		fmt.Println("hello = " + s[0] + " , world = " + s[1])
	}
}

// 標準入力を受け付けて、入力された文字を出力する
func nextLine() string {
	// 標準入力を受け付け
	sc.Scan()
	// 入力したテキストを返す
	return sc.Text()
}
