package main

import "fmt"

func main() {
	var f float64 = 10
	var n int = int(f)
	fmt.Println(n)

	var sum int
	sum = 5 + 6 + 3
	avg := float64(sum / 3)
	if avg > 3.5 {
		fmt.Println("good")
	}


	/*
		型リテラル
		- 型の具体的な定義を書き下ろした型の表現方法
		- コンポジット型などを表現するために使う
		- 変数定義やユーザ定義などで使用する
	*/

	// int型のスライスの型リテラルを使った変数定義
	var ns []int // []はスライスを表し, 後ろのintが要素の型を表す
	fmt.Println(ns) // => []

	// mapの型リテラルを使った変数定義
	var m map[string]int // keyがstring型でvalueがint型
	fmt.Println(m) // => map[]

	// mapのフィールドを指定して初期化
	lang := map[string]string{
		"rb" : "Ruby",
		"js" : "JavaScript",
	}
	fmt.Println(lang) // => map[js:JavaScript rb:Ruby]

	// 構造体の型リテラルを使った変数定義
	var p struct { // nameがstring型, ageがint型の構造体を生成する
		name string
		age  int
	}
	fmt.Println(p) // => { 0} string, int型のゼロ値の空文字と0が含まれている

	// 構造体のフィールドを指定して初期化
	p2 := struct {
		name string
		age  int
	}{
		name: "Gopher",
		age:  10,
	}
	fmt.Println(p2) // => {Gopher 10}

	// フィールドの参照
	p2.age++
	fmt.Println(p2.name, p2.age) // => Gopher 11


	// 配列の初期化
	var ar1 [5]int                       // ゼロ値で初期化
	var ar2 = [5]int{10, 20, 30, 40, 50} // 配列リテラルで初期化
	ar3 := [...]int{10, 20, 30, 40, 50}  // 要素数を値から推論
	ar4 := [...]int{5: 50, 10: 100}      // index=5が50, 10が100で他が0の要素数11の配列

	fmt.Println(ar1) // => [0 0 0 0 0]
	fmt.Println(ar2) // => [10 20 30 40 50]
	fmt.Println(ar3) // => [10 20 30 40 50]
	fmt.Println(ar4) // => [0 0 0 0 0 50 0 0 0 0 100]
	// 要素にアクセス
	fmt.Println(ar2[1]) // => 20
	// 配列ns4の長さ
	fmt.Println(len(ar4)) // => 11
	// スライス演算
	fmt.Println(ar2[1:2]) // => [20]


  // スライス(配列の一部を切り出したデータ構造)
	var ns1 []int // ゼロ値はnil
	ns1 = make([]int, 3 , 10) // 長さと容量(背後の配列の大きさ)を指定して初期化, 各要素はゼロ値
	var ns2 = []int{10, 20, 30, 40, 50} // スライスリテラルで初期化
	ns3 := []int{5: 50, 10: 100}        // index=5が50, 10が100で他が0の要素数11の配列

	fmt.Println(ns1) // => [0 0 0]
	fmt.Println(ns2) // => [10 20 30 40 50]
	fmt.Println(ns3) // => [0 0 0 0 0 50 0 0 0 0 100]
  // 要素にアクセス
	fmt.Println(ns2[3]) // => 40
	// 長さ
	fmt.Println(len(ns3)) // => 11
	// 容量
	fmt.Println(cap(ns1)) // => 10
	// 要素の追加, 容量が足りない場合は背後の配列が追加(元のおよそ2倍)
	ns1 = append(ns1, 10 ,20)
	fmt.Println(ns1) // => [0 0 0 10 20]

	var ns4 = make([]int, 2, 3) // 長さが2, 容量が3のスライスを初期化
	fmt.Println(ns4)      // => [0 0]
	fmt.Println(cap(ns4)) // => 3
	ns4 = append(ns4, 20, 30) // スライスの長さを容量より大きくする
	fmt.Println(cap(ns4)) // => 6  容量が先ほどの2倍になっている

	/*
		スライスはベースとなる配列が存在している
		ns := make([]int, 3, 10)
				 ||
		var array [10]int
		ns := array[0:3]  array[:3]とも記述可

		ns2 := []int{10, 20, 30, 40, 50}
				 ||
		var array2 [...]int{10, 20, 30, 40, 50}
		ns2 := array2[0:5]

		~イメージ~
		  スライス                配列
             ----------              0    1    2    3    4
		ポインタ  | 　※  　|             ------------------------
						 ----------      　　  | 10 | 20 | 30 | 40 | 50|
		長さ     |   5    |            -------------------------
						 ----------             /\
		容量     |   5    |              |
		        -----------             ※
	*/

	// 要素のカット
	a := []int{10, 20, 30, 40, 50}
	a = append(a[:1], a[3:]...)  // [10]に40, 50を追加
	fmt.Println(a) // => [10 40 50]

	// index'i'の要素の削除
	b := []int{10, 20, 30, 40, 50}
	b = append(b[:1], b[2:]...) // append(b[:i], b[i+1:]...)
	fmt.Println(b) // => [10 30 40 50]
}