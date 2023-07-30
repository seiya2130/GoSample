package main

import "fmt"

func main() {

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
		fmt.Println("fuga")
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

	// func
	fmt.Println(add(1, 1))

	// struct
	p1 := Person{"John", 20}
	fmt.Println(p1.hello())
	b1 := Book{"title"}

	// interface
	PrintOut(p1)
	PrintOut(b1)

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
