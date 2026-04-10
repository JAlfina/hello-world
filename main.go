package main

import "fmt"

func main() {
	var a int64
	var b int64

	fmt.Println("Оператор а = ")
	fmt.Scan(&a)

	fmt.Println("Оператор b = ")
	fmt.Scan(&b)

	fmt.Println(a, "+", b, "=", a+b)
	fmt.Println(a, "-", b, "=", a-b)
	fmt.Println(a, "/", b, "=", a/b)
	fmt.Println(a, "*", b, "=", a*b)

	fmt.Println(a, "==", b, "=", a == b)
	fmt.Println(a, "!=", b, "=", a != b)
	fmt.Println(a, "<", b, "=", a < b)
	fmt.Println(a, ">", b, "=", a > b)
	fmt.Println(a, "<=", b, "=", a <= b)
	fmt.Println(a, ">=", b, "=", a >= b)

}
