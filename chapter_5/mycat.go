package main

import (
  "fmt"
  "os"
  "flag"
  "bufio"
  "path/filepath"
)

var line_flag bool
var count int = 1

func init() {
  flag.BoolVar(&line_flag, "n", false, "flag line number")
}

func main() {
  flag.Parse()
  files := flag.Args()

  for _, file := range files {
    search_file(file)
  }
}

func search_file(file string) {
  err := filepath.Walk("/beginer_golang/chapter_5/",
                       func(path string, info os.FileInfo, err error) error {
                         if filepath.Base(path) == file {
                           print_file(file)
                         }
                           return nil
                         })
  if err != nil {
    fmt.Fprintln(os.Stderr, "検索に失敗しました", err)
  }
}

func print_file(file string) {
  sf, err := os.Open(file)

  if err != nil {
    fmt.Fprintln(os.Stderr, "読み込みに失敗しました: ", err)
    os.Exit(1)
  }

  src := bufio.NewScanner(sf)
  for src.Scan() {
    if line_flag {
      fmt.Printf("%d: ", count)
      count++
    }
    fmt.Fprintln(os.Stdout, src.Text())
  }

  defer sf.Close()
}
