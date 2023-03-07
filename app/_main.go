package _main

// GOではmain packageのmain関数から始まる

import (
	"app/foo" // ローカルパッケージの呼び出し
	"fmt"
	. "fmt"  // パッケージ名の省略（パッケージ内のフィールドに直接アクセスできる）
	fm "fmt" // パッケージ名の省略（短縮）
	"strconv"
	"time"
)

// 関数外で暗黙的な変数定義は使用することができない
//i5 := 500
// 関数外では明示的な変数定義は使用可能
var i5 int = 500

// 定数(頭文字を大文字にするとパッケージ外から呼び出すことができるpublic, 小文字ならprivate)
const Pi = 3.14
const (
	URL      = "https://xxx.co.jp"
	SiteName = "test"
)

// 値の省略
const (
	A = 1
	B
	C
	D = "A"
	E
	F
)

// iotaは連続する整数の連番を生成する
const (
	C0 = iota
	C1
	C2
)

// 構造体
type User struct {
	Name string
	Age  int
	//X, Y int
}

// 構造体のメソッド
func (u User) sayName() {
	fmt.Println(u.Name)
}

// データ更新する場合の構造体メソッド
func (u *User) SetName(name string) {
	u.Name = name
}

// 構造体の中に構造体を埋め込み
type T struct {
	User User
}

// 構造体の中に構造体埋め込み＋フィールド省略
type T2 struct {
	User
}

// 構造体のコンストラクタ
// 実際に構造体にコンストラクタ機能があるわけではないが、Goではよく使う手法
func NewUser(name string, age int) *User {
	return &User{Name: name, Age: age}
}

// 構造体とスライス
type Users []*User

// 独自型
type MyInt int

func (mi MyInt) Print() {
	fmt.Println(mi)
}

// interface
type Stringfy interface {
	ToString() string
}
type Person struct {
	Name string
	Age  int
}

func (p *Person) ToString() string {
	return fmt.Sprintf("Name=%v, Age=%v", p.Name, p.Age)
}

type Car struct {
	Number string
	Model  string
}

func (c *Car) ToString() string {
	return fmt.Sprintf("Number=%v, Mode=%v", c.Number, c.Model)
}

// カスタムエラー（GOではエラーはインターフェースとして定義される）↓の感じ
//type error interface {
//    Error() string
//}
type MyError struct {
	Message string
	ErrCode int
}

func (e *MyError) Error() string {
	return e.Message
}
func RaiseError() error {
	return &MyError{Message: "カスタムエラーが発生しました", ErrCode: 1234}
}

// fmt.Stringer(↓のように定義されている)
//type Stringer interface {
//    String() string
//}
type Point struct {
	A int
	B string
}

func (p *Point) String() string {
	return fmt.Sprintf("<<<%v, %v>>", p.A, p.B)
}

// entry point
func main() {
	fmt.Println("Hello World")
	fmt.Println(time.Now())

	// 明示的な変数定義
	var i int = 100
	fmt.Println(i)
	var s string = "Hello Go"
	fmt.Println(s)
	var t, f bool = true, false
	fmt.Println(t, f)
	var (
		i2, i21 int    = 200, 201
		s2      string = "Golang"
	)
	fmt.Println(i2, s2, i21)
	// 初期値定義なし
	var i3 int
	var s3 string
	fmt.Println(i3, s3)
	i3 = 300
	s3 = "Go"
	fmt.Println(i3, s3)

	// 暗黙的な変数定義
	i4 := 400
	fmt.Println(i4)
	i4 = 450
	fmt.Println(i4)
	// 再定義することはできない
	//i4 := 500
	//fmt.Println(i4)
	// 型を変更することはできない
	//i4 = "Hello"
	//fmt.Println(i4)
	fmt.Println(i5)
	// Goでは定義した変数は必ず使わなければならない(エラーになる)
	//var s5 string = "not use"

	// 関数呼び出し
	outer()
	// outer関数内の変数のスコープ範囲は関数に閉じられる
	//fmt.Println(s4)

	// int型
	// int型にはint8(-127~127), int16, int32, int64の4つの型があり、環境(32bit, 64bit)によって自動で決まる
	// ただし、明示的に型を指定することもできる
	var ii int8 = 100
	fmt.Println(ii)
	// 計算する際にintの型が違うと計算ができないので注意
	//fmt.Println(i + ii)
	// 現在の型を調べる(書式指定子)
	fmt.Printf("%T\n", ii)
	// 型変換
	fmt.Printf("%T\n", int32(ii))
	// 型変換を使えばさっきできなかった計算も行うことができる
	fmt.Println(i + int(ii))

	// float型(float64とfloat32しかない 暗黙的な変数定義した場合は自動的にfloat64になる)
	var fl float64 = 2.4
	fl2 := 3.2
	fmt.Printf("%T, %T\n", fl, fl2)
	zero := 0.0
	// 正の無限大 になる
	pinf := 1.0 / zero
	fmt.Println(pinf)

	// 負の無限大 になる
	ninf := -1.0 / zero
	fmt.Println(ninf)
	// NaN になる
	nan := zero / zero
	fmt.Println(nan)

	// bool型

	// string型
	var si string = `test
 test
      test`
	fmt.Println(si)
	// 1文字目
	fmt.Println(string(si[0]))

	// byte型
	byteA := []byte{72, 73}
	fmt.Println(byteA)
	// byte -> string
	fmt.Println(string(byteA))
	// string -> byte
	c := []byte("HI")
	fmt.Println(c)

	// 配列型(要素数を変更することができない)
	var arr1 [3]int
	fmt.Println(arr1)
	fmt.Printf("%T\n", arr1)
	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2)
	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)
	// 要素数を自動で設定
	arr4 := [...]string{"C", "D"}
	fmt.Println(arr4)
	fmt.Printf("%T\n", arr4)
	fmt.Println(arr2[0])
	arr2[2] = "C"
	fmt.Println(arr2)
	fmt.Println(len(arr1))

	// interface(全ての型と互換性がある -> どんな値でも入れられる)
	var x interface{}
	fmt.Println(x) // 初期値のnilは値を何も持っていない状態を表す
	x = 1
	fmt.Println(x)
	x = 3.14
	fmt.Println(x)
	x = "A"
	fmt.Println(x)
	x = [3]int{1, 2, 3}
	fmt.Println(x)
	// データ特有の処理（計算等は行えないので注意）。あくまでinterface型である
	x = 2
	//fmt.Println(x + 1)

	// 型変換
	var iii int = 1
	fl64 := float64(iii)
	fmt.Println(fl64)
	fmt.Printf("iii = %T\n", iii)
	fmt.Printf("fl64 = %T\n", fl64)
	iii2 := int(fl64)
	fmt.Printf("iii2 = %T\n", iii2)
	// 文字列からint型への変換
	var sss string = "100"
	fmt.Printf("sss = %T\n", sss)
	iiii, _ := strconv.Atoi(sss) // _は返り値2番目を使わないという意味(第二返り値はエラーが入る)。変数は必ず使わないといけないため
	fmt.Println(iiii)
	fmt.Printf("iiii = %T\n", iiii)
	// int型から文字列
	var iiii2 int = 200
	sss2 := strconv.Itoa(iiii2)
	fmt.Println(sss2)
	fmt.Printf("sss2 = %T\n", sss2)
	// 文字列からbyte型
	var h string = "Hello world"
	b := []byte(h)
	fmt.Println(b)
	// byte配列を文字列へ
	h2 := string(b)
	fmt.Println(h2)

	// 定数
	fmt.Println(Pi)
	// Pi = 3 上書きできない
	fmt.Println(URL, SiteName)
	fmt.Println(A, B, C, D, E, F)
	fmt.Println(C0, C1, C2)

	// 算術演算子
	fmt.Println(3 + 1)
	fmt.Println("ABC" + "DEF")
	n := 0
	n += 4
	fmt.Println(n)

	// 比較演算子
	fmt.Println(1 == 1)

	// 論理演算子
	fmt.Println(true && false == true)
	fmt.Println(true || false == true)

	// 関数
	fmt.Println(Plus(2, 4))
	div1, div2 := Div(9, 3)
	fmt.Println(div1, div2)
	div3, _ := Div(10, 4)
	fmt.Println(div3)
	double1 := Double(1000)
	fmt.Println(double1)
	NoReturn()

	// 無名関数
	fc := func(x, y int) int {
		return x + y
	}
	an := fc(10, 20)
	fmt.Println(an)
	an2 := func(x, y int) int {
		return x + y
	}(1, 2)
	fmt.Println(an2)

	// 関数を返す関数
	fc2 := ReturnFunc()
	fc2()

	// 関数を引数にとる関数
	CallFunction(func() {
		fmt.Println("I'm a function")
	})

	// クロージャ
	fc3 := Later()
	fmt.Println(fc3("Hello"))
	fmt.Println(fc3("World"))

	// ジェネレータ
	fc4 := integers()
	fmt.Println(fc4())
	fmt.Println(fc4())
	fmt.Println(fc4())

	// 条件分岐
	a := 0
	if a == 2 {
		fmt.Println("two")
	} else if a == 0 {
		fmt.Println("zero")
	} else {
		fmt.Println("I don't know")
	}
	// 簡易文つきIF文
	if a2 := 100; a2 == 100 {
		fmt.Println("One hundred")
	}
	// 簡易文つきIF文の注意(内部の変数が優先される。)
	a3 := 0
	if a3 := 3; true {
		fmt.Println(a3)
	}
	fmt.Println(a3)
	// これであれば上書きされる
	if a3 = 3; true {
		fmt.Println(a3)
	}
	fmt.Println(a3)

	// エラーハンドリング
	//var a4 string = "100"
	var a4 string = "aaaa"
	a5, err := strconv.Atoi(a4)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("a4 = %T\n", a5)
	}

	// for
	loop_i := 0
	for {
		loop_i++
		if loop_i >= 3 {
			break
		}
		fmt.Println("loop")
	}
	// 条件付きfor
	point := 0
	for point < 10 {
		fmt.Println(point)
		point++
	}
	// 古典的for
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// 配列for
	//arr := [3]int{1, 2, 3}
	//for i := 0; i < len(arr); i++ {
	//    fmt.Println(arr[i])
	//}
	arr := [3]int{1, 2, 3}
	for i, v := range arr {
		fmt.Println(i, v)
	}

	// switch
	n = 1
	switch n {
	case 1, 2:
		fmt.Println("1 or 2")
		break
	case 3, 4:
		fmt.Println("3 or 4")
		break
	default:
		fmt.Println("I don't know")
		break
	}
	// 代入式をswitchに入れることもできる（scopeはswitch文ないのみ）
	switch n := 3; n {
	case 1, 2:
		fmt.Println("1 or 2")
		break
	case 3, 4:
		fmt.Println("3 or 4")
		break
	default:
		fmt.Println("I don't know")
		break
	}

	// 型スイッチ
	var intf interface{} = 3
	intf1 := intf.(int) // interface型をint型で復元（型アサーション）
	fmt.Println(intf1 + 2)
	intf2, isFloat64 := intf.(float64) // 第二返り値に変換に失敗したかが返ってくる（第二返り値ないとRuntimeExceptionが発生）
	fmt.Println(intf2, isFloat64)
	// 型アサーションを使ったif文
	if intf == nil {
		fmt.Println("None")
	} else if i, isInt := intf.(int); isInt {
		fmt.Println(i, "intf is int")
	} else if s, isString := intf.(string); isString {
		fmt.Println(s, isString)
	} else {
		fmt.Println("I don't know")
	}
	// 型アサーションを使ったスイッチ（こっちの方がより可読性が高い）
	switch intf.(type) {
	case int:
		fmt.Println("int")
		break
	case string:
		fmt.Println("string")
		break
	default:
		fmt.Println("I don't know")
	}
	// 復元後の値を使用したい場合はこっち
	switch v := intf.(type) {
	case bool:
		fmt.Println(v, "bool")
		break
	case int:
		fmt.Println(v, "int")
		break
	case string:
		fmt.Println(v, "string")
		break
	default:
		fmt.Println("I don't know")
	}

	// ラベルつきfor
Loop:
	for {
		for {
			for {
				fmt.Println("start")
				break Loop
			}
			fmt.Println("処理をしない")
		}
		fmt.Println("処理をしない")
	}
	fmt.Println("end")

	// defer
	TestDefer()
	// 複数のdefer処理をまとめて登録する場合
	//defer func() {
	//   fmt.Println(1)
	//   fmt.Println(2)
	//   fmt.Println(3)
	//}()
	// よくあるのはファイルの解放処理に使用
	//file, err := os.Create("test.txt")
	//if err != nil {
	//    fmt.Println(err)
	//}
	//defer file.Close()
	//file.Write([]byte("Hello"))

	// panic & recover（基本的にあまり使わない）
	// panicはruntime errorを強制的に発生させてプログラムを終了させる機能
	//defer func() {
	//    // recoverはpanic状態でなければ値が返ってくる。panicで発生したruntime errorの状態から復帰することができる。
	//    // deferで使用する
	//    if x:= recover(); x != nil {
	//        fmt.Println(x)
	//    }
	//}()
	//panic("runtime error")
	//fmt.Println("start")

	// 並行処理（ゴルーチン）go文を使うだけで簡単に並行処理を作成することが可能
	//go sub()
	//go sub()
	//for {
	//    fmt.Println("main loop")
	//    time.Sleep(200 * time.Millisecond)
	//}

	// init（パッケージの初期化関数）
	fmt.Println("main")

	// スライス（要素数不定の配列）
	var sl []int
	fmt.Println(sl)
	// 明示的宣言
	var sl2 []int = []int{100, 200}
	fmt.Println(sl2)
	// 暗黙的宣言
	sl3 := []string{"A", "B"}
	fmt.Println(sl3)
	// make関数
	sl4 := make([]int, 5) // 要素数を５のスライスを生成
	fmt.Println(sl4)
	// 値の上書き
	sl2[0] = 1000
	fmt.Println(sl2)
	// 値の取り出し
	sl5 := []int{1, 2, 3, 4, 5}
	fmt.Println(sl5[0])
	fmt.Println(sl5[2:4])            // 3番目から4番目までの要素
	fmt.Println(sl5[:2])             // 2番目までの要素
	fmt.Println(sl5[2:])             // 3番目以降の要素
	fmt.Println(sl5[:])              // 全ての要素
	fmt.Println(sl5[len(sl5)-1])     // 最後の要素のみ
	fmt.Println(sl5[1 : len(sl5)-1]) // 最初と最後の要素以外
	// 値の追加
	sl5 = append(sl5, 6, 7, 8, 9) // 可変長引数
	fmt.Println(sl5)
	// 容量(メモリの確保量)。メモリを意識して開発するような場合に必要。おそらく上級者用
	sl6 := make([]int, 9, 15) // 要素数9, 容量15のスライスを作成
	fmt.Println(sl6)
	fmt.Println(cap(sl6))
	// ディープコピー
	copy(sl6, sl5)
	fmt.Println(sl5, sl6)
	// for（配列と一緒）
	for i, v := range sl6 {
		fmt.Println(i, v)
	}
	for i := 0; i < len(sl6); i++ {
		fmt.Println(i, sl6[i])
	}
	// 可変長引数
	fmt.Println(Sum(1, 2, 3))
	// スライスを展開して渡すことができる
	sl7 := []int{1, 2, 3}
	fmt.Println(Sum(sl7...))

	// マップ
	// 明示的宣言
	var m = map[string]int{"A": 100, "B": 200}
	fmt.Println(m)
	// 暗黙的宣言
	m2 := map[string]int{"A": 100, "B": 200}
	fmt.Println(m2)
	// 改行して宣言した場合(最後にカンマが必要になる)
	m3 := map[int]string{
		1: "a",
		2: "b",
	}
	fmt.Println(m3)
	// make関数
	m4 := make(map[int]string)
	fmt.Println(m4)
	// 値の追加
	m4[1] = "Japan"
	m4[2] = "USA"
	fmt.Println(m4)
	// 値の取り出し
	fmt.Println(m["A"])
	fmt.Println(m4[3]) // 登録されていないキーを取り出した場合は初期値0が取得される
	// 初期値が返ってきて意図せず処理が進んでしまうことがあるので、その場合はエラーハンドリングする
	m4s, success := m4[1]
	fmt.Println(m4s, success)
	_, success = m4[3]
	if !success {
		fmt.Println("error")
	}
	// 値の削除
	delete(m4, 2)
	fmt.Println(m4)
	// 要素数
	fmt.Println(len(m4))
	// for
	m5 := map[string]int{
		"Apple":  100,
		"Banana": 200,
	}
	for k, v := range m5 {
		fmt.Println(k, v)
	}

	// チャネル（データの送受信を行う）
	// 双方向チャネル宣言
	var ch1 chan int
	//// 受信専用チャネル
	//var ch2 <-chan int
	//// 送信専用チャネル
	//var ch3 chan<- int
	// 宣言のみでは書き込み、読み込みはできないnilのチャネル
	// make関数でチャネルの初期化、作成を行うことでチャネルとしての機能を持つことができる
	ch1 = make(chan int)
	// 直接make関数で作成することもできる
	ch2 := make(chan int)
	fmt.Println(cap(ch1))
	fmt.Println(cap(ch2))
	ch3 := make(chan int, 5) //容量5のチャネルを作成
	fmt.Println(cap(ch3))
	// チャネルへデータを送信
	ch3 <- 100
	ch3 <- 200
	fmt.Println("len", len(ch3))
	// チャネルからデータを受信
	// 受信するたびに要素が1つずつ減っていく (FIFO bufferとかQueueみたい)
	ch3i := <-ch3
	fmt.Println(ch3i)
	fmt.Println("len", len(ch3))
	ch3i2 := <-ch3
	fmt.Println(ch3i2)
	fmt.Println("len", len(ch3))
	// バッファサイズを超えた場合
	ch3 <- 1
	ch3 <- 2
	ch3 <- 3
	ch3 <- 4
	ch3 <- 5
	//ch3 <- 6 // deadlock(fatal error)になる
	// ゴルーチンとチャネル
	go receiver(ch1)
	go receiver(ch2)
	for i := 0; i < 20; {
		ch1 <- i
		ch2 <- i
		time.Sleep(50 * time.Millisecond)
		i++
	}
	// チャネルのクローズ
	ch4 := make(chan int, 2)
	close(ch4)
	// クローズされたチャネルへデータを送信するとRuntimeError
	//ch4 <- 1
	// クローズされたチャネルからデータを受信することはできる
	//fmt.Println(<-ch4)
	ch4i, success := <-ch4 // 第二返り値はチャネルのバッファ内が空でかつクローズされた状態の場合にfalseが入る
	fmt.Println(ch4i, success)
	// ゴルーチンとの連携
	ch5 := make(chan int, 2)
	// どのgoroutinが処理するかはタイミングによって異なる
	go receiver2("1.goroutin", ch5)
	go receiver2("2.goroutin", ch5)
	go receiver2("3.goroutin", ch5)
	for i := 0; i < 30; {
		ch5 <- i
		i++
	}
	close(ch5)
	time.Sleep(3 * time.Second)
	// チャネルのfor
	ch6 := make(chan int, 3)
	ch6 <- 1
	ch6 <- 2
	ch6 <- 3
	close(ch6) // for文で値を全て取り出した後にdeadlockで死なないためにclose処理が必要。空のチャネルから値を取り出そうとしてしまう
	for i := range ch6 {
		fmt.Println(i)
	}
	// チャネルのselect
	ch7 := make(chan int, 2)
	ch8 := make(chan string, 2)
	ch8 <- "A"
	//v1 := <- ch7 // v1に値が入っていないのでdeadlock
	//v2 := <-ch8
	//fmt.Println(v1)
	//fmt.Println(v2)
	// ↑のようにv2には値が入っているのにv1が原因でgoroutin全体を停止する問題を解決するためにselectがある
	// selectはチャネルに対する処理じゃないとエラーになる
	select {
	case v1 := <-ch7:
		fmt.Println(v1 + 1000)
		break
	case v2 := <-ch8:
		fmt.Println(v2 + "!!")
		break
	default:
		fmt.Println("どちらでもない")
		break
	}
	// select活用例
	ch9 := make(chan int, 2)
	ch10 := make(chan int, 2)
	ch11 := make(chan int, 2)
	go func() {
		for {
			i := <-ch9    // 2
			ch10 <- i * 2 // 3
		}
	}()
	go func() {
		for {
			i2 := <-ch10   // 4
			ch11 <- i2 - 1 // 5
		}
	}()
	nn := 0
L:
	for {
		select {
		case ch9 <- nn: // 1
			nn++
		case i3 := <-ch11: // 6
			fmt.Println("recieved", i3)
		default: // デフォルトでループを抜ける条件を定義することもできる
			if nn > 20 {
				break L
			}
		}
		if nn > 20 {
			break
		}
	}

	// ポインタ
	var pointa int = 100
	fmt.Println(pointa)
	fmt.Println(&pointa)
	DoublePointa(pointa)
	// 値を渡しても更新されない
	fmt.Println(pointa)
	// ポインタ型の宣言
	var p *int = &pointa
	// pの実態を表示(dereference)
	fmt.Println(*p)
	DoublePointaV2(&pointa)
	// ポインタを渡したので更新されている
	fmt.Println(pointa)

	// 構造体
	// 明示的な変数定義
	var user1 User
	// 初期値入ってる
	fmt.Println(user1)
	user1.Name = "user1"
	user1.Age = 10
	fmt.Println(user1)
	// 暗黙的な変数定義
	user2 := User{}
	fmt.Println(user2)
	user2.Name = "user2"
	fmt.Println(user2)
	// 初期値を設定
	user3 := User{Name: "user3", Age: 30}
	fmt.Println(user3)
	// フィールドを指定しないで初期値を設定。ただしフィールド宣言順に値を入れる必要がある
	user4 := User{"user4", 40}
	fmt.Println(user4)
	// newで宣言するとUser型のポインタを返すようになる
	user5 := new(User)
	fmt.Println(user5)
	// &（アドレス）演算子をつけたのと同じ意味
	user6 := &User{}
	fmt.Println(user6)
	// 関数の引数に渡す場合にポインタを使う
	updateUser(user1)  // 更新されない
	updateUser2(user6) // 更新される
	fmt.Println(user1)
	fmt.Println(user6)
	// 構造体のメソッド
	user1.sayName()
	// データを更新する場合の構造体メソッド
	user1.SetName("A")
	user1.sayName()
	// ポインタでも実行できる
	user6.SetName("B")
	user6.sayName()
	// 構造体の中に構造体を埋め込み
	tstruct := T{User: User{Name: "user1", Age: 10}}
	fmt.Println(tstruct)
	fmt.Println(tstruct.User)
	fmt.Println(tstruct.User.Name)
	// 構造体の中に構造体埋め込み＋フィールド省略
	tstruct2 := T2{User: User{Name: "user2", Age: 20}}
	fmt.Println(tstruct2)
	fmt.Println(tstruct2.User)
	fmt.Println(tstruct2.User.Name)
	// フィールド名を省略した場合、省略したフィールドを省略して出力することができる(User省略)
	fmt.Println(tstruct2.Name)
	// 構造体の中の構造体のメソッドも使える
	tstruct.User.SetName("C")
	tstruct.User.sayName()
	// 構造体のコンストラクタ
	user7 := NewUser("user7", 70)
	fmt.Println(user7)
	fmt.Println(*user7)
	// 構造体とスライス
	users := Users{}
	// ポインタを入れる
	users = append(users, &user1)
	users = append(users, &user2)
	users = append(users, &user3, &user4)
	for _, u := range users {
		fmt.Println(*u)
	}
	// makeで構造体のスライスを作成することもできる
	users2 := make([]*User, 0)
	users2 = append(users2, &user1, &user2)
	for _, u := range users2 {
		fmt.Println(*u)
	}
	// 構造体とマップ
	userMap := map[int]User{
		1: user1,
		2: user2,
		3: user3,
	}
	fmt.Println(userMap)
	userMap2 := map[User]string{
		{Name: "user1", Age: 10}: "Tokyo",
		{Name: "user2", Age: 20}: "LA",
	}
	fmt.Println(userMap2)
	// makeで構造体のマップを作成することもできる
	userMap3 := make(map[int]User)
	userMap3[1] = User{Name: "user3"}
	fmt.Println(userMap3)
	for _, v := range userMap3 {
		fmt.Println(v)
	}
	// 独自型
	var mi MyInt
	fmt.Println(mi)
	fmt.Printf("%T\n", mi)
	// 型が異なるので計算はできない（そういった処理を意図的にさせないような作りができる）
	//integer := 100
	//fmt.Println(mi * integer)
	// メソッドを追加することもできる
	mi.Print()

	// interface
	vs := []Stringfy{
		&Person{Name: "Taro", Age: 21},
		&Car{Number: "123- 456", Model: "AB-1234"},
	}
	for _, v := range vs {
		fmt.Println(v.ToString())
	}
	// カスタムエラー
	err2 := RaiseError()
	fmt.Println(err2.Error())
	// MyErrorのフィールドにアクセスしたい場合は型アサーションする必要がある
	e, ok := err2.(*MyError)
	if ok {
		fmt.Println(e.ErrCode)
	}
	// fmt.Stringer
	pp := &Point{100, "ABC"}
	fmt.Println(pp)

	// スコープ
	fmt.Println(foo.Max)
	//fmt.Println(foo.min) // privateの変数なので呼び出せない
	fmt.Println(foo.ReturnMin())
	// パッケージの省略
	fm.Println("短縮")
	Println("完全省略")
}

func outer() {
	var s4 string = "outer"
	fmt.Println(s4)
}

func Plus(x int, y int) int {
	return x + y
}

// 複数返り値
func Div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

// 返り値の変数を宣言することができる
func Double(price int) (result int) {
	result = price * 2
	return
}

// 引数返り値なし
func NoReturn() {
	fmt.Println("No Return")
}

// 関数を返す関数
func ReturnFunc() func() {
	return func() {
		fmt.Println("I'm a function")
	}
}

// 関数を引数にとる関数
func CallFunction(f func()) {
	f()
}

// クロージャ
func Later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

// ジェネレータ
func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// defer
func TestDefer() {
	defer fmt.Println("END") // TestDefer関数が完了した後に実行される
	defer fmt.Println(111)   // 複数のdefer文を書いた際には後から書いた（登録した）defer文の方が後に実行される
	fmt.Println("START")
}

// 並行処理
func sub() {
	for {
		fmt.Println("sub loop")
		time.Sleep(100 * time.Millisecond)
	}
}

// init
func init() {
	// どこにも呼び出しを書いてないけどパッケージ実行の最初に実行される
	fmt.Println("init")
}

// スライス可変長引数
func Sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

// ゴルーチンとチャネル
func receiver(c chan int) {
	// チャネルにデータが入るのを待つ
	for {
		i := <-c
		fmt.Println(i)
	}
}

// チャネルのクローズ
func receiver2(name string, c chan int) {
	for {
		i, ok := <-c
		if !ok {
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name, "END")
}

// ポインタ
func DoublePointa(i int) {
	i = i * 2
}
func DoublePointaV2(i *int) {
	*i = *i * 2
}

// 構造体
func updateUser(user User) {
	user.Name = "A"
	user.Age = 1000
}
func updateUser2(user *User) {
	user.Name = "A"
	user.Age = 1000
}
