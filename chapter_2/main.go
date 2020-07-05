package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {

	// 変数定義

	str := "Hello, World"

	/*
		var str string = "Hello, World"

		var str string
		str = "Hello, World"

		var str = "Hello, World"

		str := "Hello, World"

		var (
			str  = "Hello, World"
			str2 = "Hey!"
		)
	*/

	println(str)


	// 定数定義

	const cons_str = "Hello, " + "World"

	/*
		const n int = 100
		const m = 100
		const (
			x = 100
			y = 200
		)
	*/

	println(cons_str)


	// グループ化された名前付き定数定義において右辺は省略可能
	// 2つ目以降の右辺を省略した場合、1つ目の定数の右辺と同じ値となる
	const (
		a = 1 + 2
		b
		c
	)
	fmt.Println(a, b, c)


	// iota (連続した定数を作るための仕組み)
	// グループ化された名前付き定数定義で使われる
	// 0から始まり1ずつ加算される値として使われる

	const (
		a_2 = iota // iota = 0
		b_2
	)
	const (
		c_2 = 1 << iota // 式中でも利用可能(1ビットずつずれる)
		d_2 // 2
		e_2 // 4
	)
	fmt.Println(a_2, b_2, c_2, d_2, e_2)


	// 条件分岐(if)
	// x:= 100で定義した変数xはifとelseのブロック内で使える
	if x := 100; x == 1 {
		println("xは1")
	} else if x == 2 {
		println("xは2")
	} else {
		fmt.Printf("xは1でも2でもない, x=%v\n", x)
	}


	// 条件分岐(switch)
	switch a:= 2; a {
	case 1, 2:
		fmt.Println("a is 1 or 2")
	default:
		fmt.Println("default")
	}

	var target = 10
	switch {
	case a == 1: // caseに式が使える
		fmt.Println("a is 1")
	default:
		fmt.Printf("default, target is %v\n", target)
	}


	// 繰り返し(for)
	for i:= 0; i < 5; i++ {
		fmt.Println(i)
	}

	for i, v := range []int{1, 2, 3} {
		fmt.Printf("i is %v\n", i) // => 0, 1, 2
		fmt.Printf("v is %v\n", v) // => 1, 2, 3
	}


	// breakによるループの抜け出し
	i := 1
	for {
		if i%2 == 0 {
			fmt.Println("break!")
			break
		}
		fmt.Println(i)
		i++
	}

	j := 2
LOOP:
	for {
		switch {
		case j%2 == 1:
			fmt.Println("break LOOP!") // ラベル(LOOP)を指定したbreak
			break LOOP
		}
		fmt.Println(j)
		j++
	}


	// 1-5の数値を偶数・奇数に分類する
	for i := 1; i <= 5; i++ {
		if i%2 == 0 {
			fmt.Printf("%v - 偶数\n", i)
		} else {
			fmt.Printf("%v - 奇数\n", i)
		}
	}


	// おみくじプログラム
	rand.Seed(time.Now().UnixNano())
	switch result := rand.Intn(6) + 1; result {
	case 1:
		fmt.Println("凶")
	case 2, 3:
		fmt.Println("吉")
	case 4, 5:
		fmt.Println("中吉")
	case 6:
		fmt.Println("大吉")
	}
}