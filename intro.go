package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

//最初に実行される
// func init() {
// 	fmt.Println("First")
// }

//const (型なし)
const (
	Pi       = 3.14
	Password = "Test"
)

func main() {

	//-----------変数宣言------------------
	var (
		//数値入れないと初期値入っている
		i   int
		s   string  = "test"
		f64 float64 = 2.6
		f   bool    = true
	)
	//関数内ではこのように宣言しても良い
	//xi := 1
	//xf := "test"
	fmt.Println(i, s, f64, f)
	fmt.Println(Pi, Password)
	//strings にいろんな関数ある
	fmt.Println(strings.Contains(s, "st"))

	//-----------キャスト-----------------
	ii := int(f64)
	fmt.Printf("%T, %v, %d\n", ii, ii, ii)
	var ss string = "124a"
	iii, err := strconv.Atoi(ss)
	if err != nil {
		fmt.Printf("ERROR : %s\n", err)
	} else {
		fmt.Printf("%T, %d, %v\n", iii, iii, iii)
	}

	//-----------配列・スライス------------------
	var a [2]int = [2]int{2, 4}
	fmt.Println(a)
	//スライス
	var sr []int = []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(sr)
	//要素の追加
	var srr = append(sr, 100, 200, 300)
	fmt.Println(srr)
	//make, cap, len, value
	new1 := make([]int, 3, 5)
	fmt.Printf("len=%d cap=%d, value=%v\n", len(new1), cap(new1), new1)

	//------------map-------------------------
	m := map[string]int{"apple": 200, "banana": 150}
	fmt.Println(m["apple"])
	m["new"] = 600
	fmt.Println(m)
	//第二引数は、戻り値が存在するか否かの真偽値を返す。
	v, ok := m["apple"]
	fmt.Println(v, ok)
	v1, ok1 := m["nothing"]
	fmt.Println(v1, ok1)
	//初期化は、リストと同様
	//var mm map[string]int は初期化できてない（nil）
	mm := make(map[string]int)
	mm["new"] = 200000000000000
	fmt.Println(mm)

	//-----------byte------------------
	b := []byte{72, 73}
	fmt.Println(b)
	fmt.Println(string(b))

	//-----------func------------------
	a1 := 1
	b2 := 2
	r1, r2 := add(a1, b2)
	fmt.Println("Add Function")
	fmt.Println(r1, r2)
	fmt.Println("----------------")

	//-----------クロージャー-------------
	fmt.Println("------クロージャー------")
	c1 := circleArea(3.14)
	fmt.Println(c1(9))
	c2 := circleArea(3)
	fmt.Println(c2(9))
	fmt.Println("------------------------")

	//-------------可変長引数-------------
	fmt.Println("------可変長引数------")
	foo()
	foo(1, 2)
	foo(1, 2, 3)
	list := []int{11, 22, 33}
	foo(list...)
	fmt.Println("--------------------")

	//-------------range-------------
	//for i:=0; i < len(x); i++ {処理}　と同じことができる
	fmt.Println("------range------")
	tlist := []int{10, 20, 30, 40}
	for c, p := range tlist {
		fmt.Println(c, p)
	}
	tmap := map[string]int{
		"banana": 200,
		"apple":  150,
		"test":   200000,
	}
	for c, p := range tmap {
		fmt.Println(c, p)
	}
	fmt.Println("-----------------")

	//-------------defer-------------
	//関数内の全ての処理が終了してから実行される
	//複数存在する場合は、最後尾から実行される。
	// fmt.Println("--------defer---------")
	// defer fmt.Println("----------------------")
	// defer fmt.Println("1")
	// defer fmt.Println("2")
	// defer fmt.Println("3")
	// fmt.Println("start")

	//-------------logging-------------
	LoggingSettings("practice.log")
	_, err = os.Open("nothing")
	if err != nil {
		//log.Fatalln("Exit", err)
		log.Println("Exit", err)
	}

	//---------------new と make ---------------
	fmt.Println("------new & make------")
	list1 := make([]int, 10)
	fmt.Printf("%T\n", list1)
	map1 := make(map[string]int)
	fmt.Printf("%T\n", map1)
	ch := make(chan int)
	fmt.Printf("%T\n", ch)

	var p *int = new(int)
	fmt.Printf("%T\n", p)
	var st = new(struct{})
	fmt.Printf("%T\n", st)
	fmt.Println("-------------------")

	//-------------struct----------------
	// fmt.Println("------struct------")
	// Ver := Vertex{X: 1, Y: 2, S: "TEST"}
	// fmt.Println(Ver)
	// fmt.Println("------------------")
	// //メソッド、ポインタレシーバ、値レシーバ
	// fmt.Println("------メソッド(ポインタレシーバ、値レシーバ)------")
	// Ver.Scale(10)
	// fmt.Println(Ver.Area())
	// fmt.Println("-------------------------------------------")
	//コンストラクタ
	vc := New(3, 5, "constructor")
	fmt.Println(vc)

	//インラーフェース
	var mike Human = &Person{"Mike"}
	DriveCar(mike)

	//タイプアサーション, switch type
	fmt.Println("------type assertion && switch type------")
	do(10)
	do("type")
	do(true)
	fmt.Println("--------------------------")

	//------------Goroutine(並列処理)-------------
	fmt.Println("------Goroutine------")
	var wg sync.WaitGroup
	wg.Add(1)
	go goroutine("World", &wg)
	normal("Hello")
	wg.Wait()
	fmt.Println("----------------------")
	//------------------chanel----------------
	fmt.Println("------Channel------")
	lis := []int{1, 2, 3, 4, 5}
	c := make(chan int)
	go routine1(lis, c)
	go routine2(lis, c)
	x := <-c
	fmt.Println(x)
	y := <-c
	fmt.Println(y)
	//buffered channels
	cc := make(chan int, 2)
	cc <- 100
	fmt.Println("First", len(cc))
	cc <- 200
	fmt.Println("Second", len(cc))
	close(cc)
	for c := range cc {
		fmt.Println(c)
	}
	fmt.Println("-------------------")
	//fan-out, fun-in
	first := make(chan int)
	second := make(chan int)
	third := make(chan int)

	go producer(first)
	go multi2(second, first)
	go multi4(third, second)
	for result := range third {
		fmt.Println(result)
	}
	//selection, for break
	// tick := time.Tick(100 * time.Millisecond)
	// boom := time.After(500 * time.Millisecond)
	// AnithigIsOK:
	// 	for {
	// 		select {
	// 		case t := <-tick:
	// 			fmt.Println("tick : ", t)
	// 		case <-boom:
	// 			fmt.Println("Boom!")
	// 			break AnithigIsOK
	// 		default:
	// 			fmt.Println("     .")
	// 			time.Sleep(50 * time.Millisecond)
	// 		}
	// 	}
	// fmt.Println("#######Fin#######")
}

//fan-out, fun-in
func producer(first chan int) {
	defer close(first)
	for i := 0; i < 10; i++ {
		first <- i
	}
}
func multi2(second chan<- int, first <-chan int) {
	defer close(second)
	for c := range first {
		second <- c * 2
	}
}
func multi4(third chan int, second chan int) {
	defer close(third)
	for c := range second {
		third <- c * 4
	}
}

//------------------chanel----------------
func routine1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}
func routine2(s []int, c chan int) {
	sum := 1
	for _, v := range s {
		sum = sum * v
	}
	c <- sum
}

//------------Goroutine(並列処理)-------------
func normal(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}
func goroutine(s string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
	wg.Done()
}

//タイプアサーション, switch type
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("int : ", v*10)
	case string:
		fmt.Println("string : " + v)
	default:
		fmt.Println("I don't know : ", v)
	}
}

//インターフェース
type Human interface {
	Say() string
}
type Person struct {
	Name string
}

func (p *Person) Say() string {
	p.Name = "Mr." + p.Name
	fmt.Println(p.Name)
	return p.Name
}
func DriveCar(human Human) {
	if human.Say() == "Mr.Mike" {
		fmt.Println("OK! Drive")
	} else {
		fmt.Println("No!!")
	}
}

//Embedded (継承)
type Vertex3D struct {
	//Embedded
	Vertex
	z int
}

//コンストラクタ
func New(x, y int, s string) *Vertex {
	return &Vertex{x, y, s}
}

//メソッド(ポインタレシーバ、値レシーバ)
// func (v Vertex) Area() int {
// 	return v.X * v.Y
// }
// func (vp *Vertex) Scale(i int) {
// 	vp.X = i * vp.X
// 	vp.Y = i * vp.Y
// }
//-------------struct----------------
type Vertex struct {
	//大文字だとprivate扱いになるから注意
	x int
	y int
	s string
}

//-------------log-------------
func LoggingSettings(logFile string) {
	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.SetOutput(multiLogFile)
}

//-------------可変長引数-------------
func foo(params ...int) {
	fmt.Println(params)
	for _, param := range params {
		fmt.Println(param)
	}
}

//-----------クロージャー-------------
func circleArea(pi float64) func(radius float64) float64 {
	return func(radius float64) float64 {
		return pi * radius * radius
	}
}

//-----------func------------------
//戻り値、型など基本的に後
//func add(a int, b int) (int, int) {
func add(a int, b int) (result int, ans int) {
	result = a + b
	ans = a - b

	//関数内に関数の定義も可能
	func() {
		fmt.Println("------func------")
		fmt.Println("First Inner Func")
	}()
	f := func() {
		fmt.Println("Second Inner Func")
	}
	f()

	return result, ans
}
