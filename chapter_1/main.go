package main

// fmtパッケージのインポート
import "fmt"

func main() {
	// println("コメントアウト")

	/*
		ブロックコメント
	*/

	// fmtパッケージの標準出力関数を呼び出す
	fmt.Print("Hello, ") // 改行されない
	fmt.Println("World") // 改行される
	fmt.Println("A", 100, true, 1.5) // スペース区切りで出力

	// 書式を指定して標準出力を行う
	fmt.Printf("Hello, World\n")     // \nで改行する
	fmt.Printf("%d-%s", 100, "偶数\n") // %dは整数、%sは文字列

	var price int
	fmt.Print("値段>")
	fmt.Scan(&price)
	fmt.Printf("%d円\n", price)
}