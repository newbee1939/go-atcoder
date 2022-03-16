// 標準入力の例（一行ずつ読み込む）

// [入力値]
// 2
// 2 5
// 3 4

// [出力値]
// hello = 2 , world = 5
// hello = 3 , world = 4

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// os.Stdin:Goの標準入力
	// Scannerを用意
	sc := bufio.NewScanner(os.Stdin)

	// トークンを読み込む（標準入力を受け付ける）（一回目）（ここで最初の「行数」の部分が読み込まれる）（一行しか読まれない）
	sc.Scan()
	// sc.Text():Scannerが読み込んだトークンを文字列として取り出す
	// ここでは最初に入力した行数がnにセットされる（AtoiでIntegerに変換）
	var n, _ = strconv.Atoi(sc.Text())
	// 最初に入力した行数分forを回す
	for i := 0; i < n; i++ {
		// 入力を受け付ける（二行目以降）
		sc.Scan()
		// " "で入力した文字を分割する
		var s = strings.Split(sc.Text(), " ")
		// それぞれが配列に入っているので一つずつ出力する
		fmt.Println("hello = " + s[0] + " , world = " + s[1])
	}
}
