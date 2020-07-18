package main

import (
	"fmt"
	"flag"
	"strings"
)

var msg = flag.String("msg", "default value", "description")
var n int

func init() {
	// ポインタを指定して設定を予約
	flag.IntVar(&n, "n", 1, "回数")
}

func main() {
	// ここで実際に設定される
	flag.Parse()
	fmt.Println(strings.Repeat(*msg, n))
}

/*
	$ go run main.go -msg=こんにちは -n=2
	=> こんにちはこんにちは
*/