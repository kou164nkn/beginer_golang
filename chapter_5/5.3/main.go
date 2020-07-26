package main

import (
  "fmt"
  "os"
  "bufio"
  "path/filepath"
)

/*
  標準入力      os.Stdin
  標準出力      os.Stdout
  標準エラー出力 os.Stderr
*/

func main() {
  try_defer()
  try_scanner()
  try_file()
  try_dir_walk()
  try_err()
}

func try_defer() {
  /*
    defer 関数の遅延実行
      - 関数終了時に実行される
      - 引数の評価はdefer呼び出し時
      - スタック形式で呼び出される(最後に呼び出したものから実行)
      - 予約した関数呼び出しはreturn時に実行される

    下の例だと
    => hello
    => world
    => !!!
  */

  msg := "!!!"
  defer fmt.Println(msg)
  msg = "world"
  defer fmt.Println(msg)
  fmt.Println("hello")
}

// 1行ずつ読み込む関数
func try_scanner() {
  // 標準入力から読み込む
  fmt.Printf("%s", "Please input > ")
  scanner := bufio.NewScanner(os.Stdin)

  // 1行ずつ読み込んで繰り返す
  for scanner.Scan() {
    // exitが入力されたら終了
    if scanner.Text() == "exit" {
      return
    }
    // 1行分を出力する
    fmt.Println(scanner.Text())
    fmt.Printf("%s", "Please input > ")
  }

  // まとめてエラー処理する
  if err := scanner.Err(); err != nil {
    fmt.Fprintln(os.Stderr, "読み込みに失敗しました", err)
  }
}


// ファイルパスを取り扱う関数
func try_file() {
  // パスを結合する
  path := filepath.Join("dir", "main.go")

  // 拡張子を取得
  fmt.Println(filepath.Ext(path))
  // ファイル名を取得
  fmt.Println(filepath.Base(path))
  // ディレクトリ名を取得
  fmt.Println(filepath.Dir(path))
}


func try_dir_walk() error {
  // Goファイルを探し出す
  // 第1引数にディレクトリ名を、第2引数にファイルorディレクトリが見つかった際の処理を記載
  err := filepath.Walk("check.go",
    func(path string, info os.FileInfo, err error) error {
      if filepath.Ext(path) == ".go" {
        fmt.Println(path)
      }
      return nil
    })

  if err != nil {
    return err
  }
  return nil
}


// 標準エラー出力で表示し、プログラムを終了させる関数
func try_err() {
  fmt.Fprintf(os.Stderr, "エラー")
  os.Exit(1) // 終了コードを指定してプログラムを終了. デフォルトは0(成功)

  /*
    log.Fatal 標準エラー出力にエラーメッセージを表示し、os.Exit(1)で異常終了させる

    if err := f(); err != nil {
      log.Fatal(err)
    }
  */
}

/*
  // 読み込み用にファイルを開く
  sf, err := os.Open(src)
  if err != nil {
    return err
  }

  // 関数終了時に閉じる
  defer sf.Close()
*/

/*
  // 書き込み用にファイルを開く
  df, err := os.Create(dst)
  if err != nil {
    return err
  }

  // 関数終了時に閉じる
  defer func() {
    if err := df.Close(); err != nil {
      rerr = err
    }
  }
*/
