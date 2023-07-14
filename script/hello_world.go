package main

import (
	"fmt"
	"os/user"
	"strings"
	"strconv"
	"time"
)

func init() {
	fmt.Println("This is init function!!")
}

func hello_world() {
	fmt.Println("Hello World")
}

func main() {
	hello_world()
	fmt.Println("My Info is below:")
	fmt.Println(user.Current())
	fmt.Println("The time is", time.Now())

	const sample_text string = "Hello World"

	// 👇型を変換しないとASCIIコードが表示される
	fmt.Println("The first letter of the string is", string(sample_text[0]))
	// 文字の抽出
	fmt.Println(sample_text[6:11])
	// 文字の置換
	fmt.Println(strings.Replace(sample_text, "World", "Go", 1))
	// 検索 return bool
	fmt.Println(strings.Contains(sample_text, "World"))
	// 複数行の文字列
	fmt.Println(`Hello
	
World`)

	// 文字列の数字→数値への変換 return int error
	// _は以降使わない変数を表すときに使う。
	i, _ := strconv.Atoi("100")
	fmt.Println(i)

	// 配列とスライス
	// 配列は固定長、スライスは可変長
	a := [5]int{1, 2, 3, 4, 5}
	s := []int{1, 2, 3, 4, 5}

	fmt.Println(a)
	fmt.Println(append(s, 6))

	// makeとcap
	// make(型, 長さ, 容量)
	// 長さは要素数、容量は配列の長さ
	// 配列の長さを超えた要素を追加すると容量が増える
	a2 := make([]int, 3, 5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(a2), cap(a2), a2)

	// 下記の違いはメモリが確保されるか否か
	// a3 := make([]int, 0)　👈確保する
	// a3 := []int{}　👈確保しない

	c := make([]int, 5)
	// c := make([]int, 0, 5)
	for i := 0; i < 5; i++ {
		c = append(c, i)
		fmt.Println(c)
	}
	fmt.Println(c)
}
