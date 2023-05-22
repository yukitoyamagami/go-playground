package main

import (
	// 一時的なI/Oを扱うための標準ライブラリ
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := input()
	fmt.Println("入力値は:", input, "です")
	inferType(input)
}

func input() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("値を入力してください:")

	if scanner.Scan() {
		return scanner.Text()
	}

	if scanner.Err() != nil {
		fmt.Println("Error reading from input:", scanner.Err())
	}
	return ""
}

func inferType(input string) {
	_, err := strconv.Atoi(input)
	if err == nil {
		fmt.Println("入力値は数値です")
		return
	}

	_, err = strconv.ParseFloat(input, 64)
	if err == nil {
		fmt.Println("入力値は浮動小数点数です")
		return
	}

	_, err = strconv.ParseBool(input)
	if err == nil {
		fmt.Println("入力値は真偽値です")
		return
	}

	fmt.Println("入力値は文字列です")
}
