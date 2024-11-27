package main

import "fmt"

func test() {
	fmt.Println("test")
	num := 0
	fmt.Println(num)

	type Test string
	var t1 Test = "1"
	fmt.Println(t1)

	arry := [3]int{0, 1, 2}
	for i := 0; i < len(arry); i++ {
		fmt.Println(arry[i])
	}

	s1 := []int{}
	s1 = append(s1, 1)
	fmt.Println(s1)

	if 1 > 0 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

	x := 2
	switch x {
	case 1:
		fmt.Println(true)
	default:
		fmt.Println(false)
	}

	h1 := Human{"John", 20}
	h1.hello()
}

func main() {
	test()
	return

	// 出力
	fmt.Println("Hello World!")
	num := 123
	str := "abc"
	fmt.Print(num, str)
	fmt.Println(num, str)
	fmt.Printf("num=%d, str=%s", num, str)

	// var
	a1 := 1
	a2 := "test"
	fmt.Println(a1, a2)

	// type
	type Test string
	var t1 Test = "1"
	fmt.Println(t1)
	t2 := Test(t1)
	fmt.Println(t2)

	// array
	ary1 := [3]int{}
	ary1[0] = 1
	ary1[1] = 2
	ary1[2] = 3
	fmt.Println(ary1[0], ary1[1], ary1[2])

	ary2 := [...]string{"1", "2", "3"}
	fmt.Println(ary2)

	// slice
	s1 := []int{}
	s1 = append(s1, 1)
	s1 = append(s1, 2)
	s1 = append(s1, 3)
	fmt.Println(s1)
	fmt.Println(len(s1), cap(s1))

	// map
	m1 := map[string]int{
		"a": 100,
		"b": 200,
	}
	fmt.Println(m1["a"])
	m1["c"] = 300
	fmt.Println(m1["c"])

	// if
	x := 2
	y := 1
	if x > y {
		fmt.Println(true)
	}
	x = 0
	if x > y {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

	// switch
	str1 := "hoge"
	switch str1 {
	case "hoge":
		fmt.Println("hoge")
		fallthrough
	case "fuga":
	}

	// goto
	gt := '1'
	if gt == '1' {
		fmt.Println("goto1")
		goto Done
		fmt.Println("goto2")
	}

Done:
	fmt.Println("done")

	// for文
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// func
	fmt.Println(add(1, 1))

	// struct
	p1 := Person{"John", 20}
	fmt.Println(p1.hello())
	b1 := Book{"title"}

	// interface
	PrintOut(p1)
	PrintOut(b1)

	// generics
	fmt.Println(GenericAdd(1, 2))
	fmt.Println(GenericAdd(1.5, 2.5))

	// pointer
	po1 := 1
	po2 := 2
	fn(po1, &po2)
	fmt.Println(po1, po2)
}

func add(x int, y int) int {
	return x + y
}

type Person struct {
	name string
	age  int
}

type Human struct {
	name string
	age  int
}

func (h *Human) hello() {
	fmt.Println("Hello, I'm " + h.name)
}

func (p *Person) hello() string {
	return "Hello, I'm " + p.name
}

func (p Person) ToString() string {
	return p.name
}

type Book struct {
	title string
}

func (b Book) ToString() string {
	return b.title
}

type Printable interface {
	ToString() string
}

func PrintOut(p Printable) {
	fmt.Println(p.ToString())
}

func fn(a int, b *int) {
	a = 100
	*b = 200
}

// generics
func GenericAdd[T int | float64](a T, b T) T {
	return a + b
}
