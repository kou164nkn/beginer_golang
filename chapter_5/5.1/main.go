package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		ビルドしてから実行
			$ go build main.go
			$ ./main hello world
			// => [./main hello world]
	*/
	fmt.Println(os.Args)
}