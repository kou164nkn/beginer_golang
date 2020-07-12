package main

import (
	"fmt"
	"time"

	// サードパーティ製は空行を挟んで分けることが多い
	"github.com/kou164nkn/greeting" // インポートパスを記述する (事前にgo get~でGOPATHに保存しておく)
  greeting2 "github.com/kou164nkn/greeting/v2"
	/*
		パッケージ名のエイリアス
		- インポートパスの左側に変えたい名前を書く
		- 同じパッケージ名のパッケージを使いたい場合に書く
		- インポートパスとパッケージ名が一致していない場合に用いる

  	mysync "github.com/tenntenn/sync" // syncパッケージと名前が衝突している
		greeting2 "github.com/tenntenn/greeting/v2" // インポートパスとパッケージ名が一致していない
	*/
)

// パッケージ変数　関数間で変数が共有可能
var msg string = "hello" // :=記法は関数内でしか使えない

func f() { fmt.Println(msg) }

func main() {
	fmt.Println(greeting.Do())

	f() // => "hello"
	msg = "hi, gophers"
	f() // => "hi gophers"

	fmt.Println(greeting2.Do(time.Now()))
}