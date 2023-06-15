package main

import (
	// 一時的なI/Oを扱うための標準ライブラリ
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input_value, err := input()
	if err != nil {
		fmt.Println("エラーが発生しました:", err)
	} else {
		fmt.Println("入力値は:", input_value, "です")
	}

	inferType(input_value)
}

func input() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("値を入力してください:")

	if scanner.Scan() {
		return scanner.Text(), nil
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return "", errors.New("no input provided")
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
