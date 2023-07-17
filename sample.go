package main

import "fmt"

func main() {
	fmt.Println("Hello World!");
	num := 123
	str := "abc"
	fmt.Print(num, str)
	fmt.Println("=====")
	fmt.Println(num, str)
	fmt.Println("=====")
	fmt.Printf("num=%d, str=%s", num, str)
}