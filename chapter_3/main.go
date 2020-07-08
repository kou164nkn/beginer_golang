package main

import (
	"fmt"
	"net/http"
)

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


	// スライスの要素の和
	slice := []int{19, 86, 1, 12}
	var sum_slice int
	for _, number := range slice {
		sum_slice += number
	}
	fmt.Println(sum_slice)


	/*
	  マップ型 キーと値をマッピングさせるデータ構造
		マップの初期化
			var m1 map[string]int             // ゼロ値はnil
			var m2 = make(map[string]int)     // makeで初期化
			var m3 = make(map[string]int, 10) // 容量指定できる
			m5 := map[string]int{}            // 空の場合
	*/
	m4 := map[string]int{"x": 10, "y": 20}  // リテラルで初期化

	fmt.Println(m4["x"]) // => 10
	m4["z"] = 30
	n, ok := m4["z"]
	fmt.Println(n, ok)  // => 30 true
	delete(m4, "z")
	n, ok = m4["z"]
	fmt.Println(n, ok) // => 0 false


	// コンポジット型を要素に持つコンポジット型
	// e.g. マップの値がスライス型
	pinp := map[string][]int {
		"x": []int{1, 2, 3},
		"y": []int{10, 20, 30},
	}
	fmt.Println(pinp) // => map[x:[1 2 3] y:[10 20 30]]


	/*
		ユーザ定義型
		type 型名 基底型 で定義する

		// 組み込み型を基にする
		type MyInt int
		// 他のパッケージの型を基にする
		type MyWriter io.Writer
		// 型リテラルを基にする
		type Person struct {
			Name string
		}
	*/

	// 基底型とユーザ定義型で相互キャスト可能
	type MyInt int
	var num_int int = 100
	num_myint := MyInt(num_int)
	num_int = int(num_myint)
	fmt.Println(num_int, num_myint) // => 100 100

	/*
		次の仕様を満たすユーザ定義型
		- とあるゲームの得点を集計する
		- ゲームの結果は0~100点まで1点刻みで記録
		- 集計は複数回のゲームの結果を元にユーザ毎に行う
	  どういうデータ皇族で1回のゲーム結果を表現すべきか？
	*/
	type GameScore struct {
		UserID string
		GameID int
		Score  int
	}

	/*
		型エイリアス(Go 1.9以上)
		- 完全に同じ型 (前述のユーザ定義型は全く別の型を新たに定義している)
		- キャスト不要
		- エイリアス側ではメソッドは定義できない
	*/
	type Applicant = http.Client
	fmt.Printf("%T\n", Applicant{}) // => http.Client


  func_result := add(2, 3)
	fmt.Println(func_result) // => 5

	hoge, fuga := swap(10, 20)
	fmt.Println(hoge, fuga) // => 20 10
	/*
		省略したい場合はブランク変数を用いる
			hoge, _ := swap(10, 20)
			fmt.Println(hoge) // => 20
			_, fuga := swap(10, 20)
			fmt.Println(fuga) // => 10
	*/

	hoge, fuga = swap2(10, 20)
	fmt.Println(hoge, fuga) // => 20 10


	// 無名関数(クロージャ)
	msg := "test"
	func() {
		fmt.Println(msg) // 関数外のmsgを参照できる
	}() // => test   無名関数を定義しすぐに呼び出している


	fs := make([]func() string, 2) // 長さが2で関数型を要素に持つスライスを生成
	fs[0] = func() string { return "hoge" } // fsの各要素の関数を代入
	fs[1] = func() string { return "fuga" }
	for _, f := range fs { // indexは省略し、要素だけを取得
		fmt.Println(f()) // 要素の関数を実行
	}

	/*
		クロージャのよくあるバグ
		定義と実行のタイミングを気を付ける
		- 関数外の変数(自由変数)を参照している場合, 実行のタイミングでは値が変わっている可能性がある
	*/
	fs2 := make([]func(), 3)
	for i := range fs2 {
		fs2[i] = func() { fmt.Println(i) } // fs2[0],[1]のiにもiに最後に入った2が入っている
	}
	for _, f := range fs2 { f() } // 0 1 2を期待したが2 2 2となる


	p_test := struct{age int; name string}{age: 10, name: "Gopher"}
	p_test2 := p_test
	p_test2.age = 20
	fmt.Println(p_test.age, p_test.name)   // => 10 "Gopher"　代入はコピーのため元の値に影響はない
	fmt.Println(p_test2.age, p_test2.name) // => 20 "Gopher"

	var x_pointa int
	f_pointa(&x_pointa) // &でポインタを取得する
	fmt.Println(x_pointa)

	// スライス, マップ, チャネルは内部でポインタが使われている
	ns_pointa := []int{10, 20, 30}
	ns_pointa2 := ns_pointa
	ns_pointa[1] = 200
	fmt.Println(ns_pointa[0], ns_pointa[1], ns_pointa[2])    // => 10 200 30
	fmt.Println(ns_pointa2[0], ns_pointa2[1], ns_pointa2[2]) // => 10 200 30


	for i := 1; i <=5; i++ {
		fmt.Print(i)
		if check(i) {
			fmt.Println("-偶数")
		} else {
			fmt.Println("-奇数")
		}
	}

	n_try, m_try := 10, 20
	swap_pointa(&n_try, &m_try)
	fmt.Printf("use pointa: %d %d\n", n_try, m_try) // => use pointa: 20 10


  // 353行目
	var hex Hex = 100
	fmt.Println(hex.String()) // => 64


  // 366行目
	var v T
	(&v).f() // => hi
	v.f()    // => hi 上と同じ意味
}

// func 関数名(引数) 戻り値 {} で定義する
func add(x int, y int) int {
	return x + y
}

// 型をまとめて定義可能で複数の戻り値を定義可能
func swap(x, y int) (int, int) {
	return y, x
}

// 名前付き戻り値
func swap2(x, y int) (x2, y2 int) { // x2, y2が名前付き戻り値
	y2, x2 = x, y
	return // 明示しない場合は戻り値用の変数の値が返される
}

/*
	ポインタ 変数の格納先を表す値
	- 値で渡される型の値に対して破壊的な操作を加える際に利用する
		- 破壊的な操作 = 関数を出てもその影響が残る
*/
func f_pointa(xp *int) { // intのポインタ型
	*xp = 100 // *でポインタの指す先に値を入れる
}

// 偶数・奇数を判定する関数
func check(i int) bool {
	if i%2 == 0 { return true }
	return false
}

// ポインタを使ったswap
func swap_pointa(n, m *int) {
	*n, *m = *m, *n
}


/*
	メソッド レシーバと紐づけられた関数
	- データとそれに対する操作を紐付けるために用いる
	- ドットでメソッドにアクセスする

	func (レシーバ レシーバの型) メソッド名 戻り値の型 {} で定義
*/
type Hex int
func (h Hex) String() string {
	return fmt.Sprintf("%x", int(h)) // fmt.Sprintf() string型と他の型をstring型として扱う
}


/*
	レシーバ メソッドに関連づけられた変数
	- メソッド呼び出し時には通常の引数と同じような扱いになる
		- コピーが発生する
	- ポインタを用いることでレシーバへの変更を呼び出し元に伝えることができる
		- レシーバがポインタの場合もドットでアクセスする
*/
type T int
func (t *T) f() { fmt.Println("hi") }