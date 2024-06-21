package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// ファイルを開く
	file, err := os.Open("kaiki.txt")
	if err != nil {
		fmt.Println("ファイルを開く際にエラーが発生しました:", err)
		return
	}
	defer file.Close()

	// ファイルを読み込む
	scanner := bufio.NewScanner(file)
	var line string
	if scanner.Scan() {
		line = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("ファイルを読み取る際にエラーが発生しました:", err)
		return
	}

	// コンマで分割して整数に変換する
	parts := strings.Split(line, ",")
	if len(parts) != 2 {
		fmt.Println("予期しないデータ形式です")
		return
	}

	a, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		fmt.Println("最初の部分を整数に変換する際にエラーが発生しました:", err)
		return
	}

	b, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		fmt.Println("二番目の部分を整数に変換する際にエラーが発生しました:", err)
		return
	}

	// 結果を表示する
	fmt.Println("a:", a)
	fmt.Println("b:", b)
}
