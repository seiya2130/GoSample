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

}
