package main

import (
	"fmt"
)

// var n = 1 // パッケージ変数とする場合はここに書く

/*
---------------
main
---------------
*/
func main() {
	// 配列に関するアレコレ
	execArray()

	// interface{}に関するアレコレ
	execInterface()

	// 関数に関するアレコレ
	execFunc()

	// 定数に関するアレコレ
	execConst()

	// for文に関するアレコレ
	execFor()

	// goroutineに関するアレコレ
	execGo()
}

/*
---------------
配列に関するアレコレ
---------------
*/
func execArray() {
	n := one() + two()
	fmt.Printf("数値 = %d\n", n)

	// 配列（配列は固定長。可変長としたい場合はスライスを使う）
	arr1 := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%v\n", arr1)
	fmt.Printf("%v\n", arr1[0])

	// 配列 > 要素数の省略
	arr2 := [...]string{"aaa", "bbb", "ccc"}
	fmt.Printf("%v\n", arr2[0])

	// 配列 > 要素への代入
	arr2[0] = "ddd"
	fmt.Printf("%v\n", arr2[0])
}

func one() int {
	return 1
}

func two() int {
	return 2
}

/*
---------------
interface{}に関するアレコレ
---------------
*/
func execInterface() {
	var i1 interface{}
	i1 = 1
	fmt.Printf("%v\n", i1)

	var i2 interface{}
	i2 = "interface"
	fmt.Printf("%v\n", i2)
}

/*
---------------
関数に関するアレコレ
---------------
*/
func execFunc() {
	// 基本
	plus := plus(1, 5)
	fmt.Printf("%v\n", plus)

	// 複数の戻り値
	q, r := div(19, 7)
	fmt.Printf("商 = %d 剰余 = %d\n", q, r)

	// 戻り値の破棄
	q2, _ := div(19, 7)
	fmt.Printf("商 = %d\n", q2)

	// 無名関数
	f := func(x, y int) int { return x + y }
	fmt.Printf("%d\n", f(2, 3))

	// 関数を返す関数
	rf := returnFunc()
	rf()

	// 関数を引数に取る関数
	callFunction(func() {
		fmt.Println("I'm a function!!!!")
	})
}

func plus(x, y int) int {
	return x + y
}

func div(a, b int) (int, int) {
	q := a / b
	r := a % b
	return q, r
}

func returnFunc() func() {
	return func() {
		fmt.Println("I'm a function")
	}
}

func callFunction(f func()) {
	f()
}

/*
---------------
定数に関するアレコレ
---------------
*/
func execConst() {
	const (
		X  = 1
		Y  = 2
		Z  = 3
		S1 = "あ"
		S2
	)

	fmt.Printf("%d\n", X+Y+Z)
	fmt.Printf("%v %v\n", S1, S2)
}

/*
---------------
for文に関するアレコレ
---------------
*/
func execFor() {
	i := 0
	for {
		fmt.Println(i)
		i++
		if i == 3 {
			break
		}
	}

	i2 := 0
	for i2 < 3 {
		println(i2)
		i2++
	}

	fruits := [3]string{"Apple", "Banana", "Cherry"}
	for i := range fruits {
		fmt.Printf("%s\n", fruits[i])
	}
}

/*
---------------
goroutineに関するアレコレ
---------------
*/
func execGo() {
	go sub()
	for {
		fmt.Println("main loop")
	}
}

func sub() {
	for {
		fmt.Println("sub loop")
	}
}
