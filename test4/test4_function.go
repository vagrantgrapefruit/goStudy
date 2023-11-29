package test4

import "fmt"

func foo1(a string, b int) int {
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	c := 100
	return c
}

//returns multiple results anonymously
func foo2(a string, b int) (int, int) {
	fmt.Println("-----foo2-----")
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	return 123, 321
}

func foo3(a string, b int) (r1 int, r2 int) {
	fmt.Println("------foo3-----")
	fmt.Println("a=", a)
	fmt.Println("b=", b)

	r1 = b + 100
	r2 = b + 200

	return
}

func foo4(a string, b int) (r1, r2 int) {
	fmt.Println("------foo4-----")
	fmt.Println("a=", a)
	fmt.Println("b=", b)

	r1 = b + 100
	r2 = b + 200

	return
}

func init() {
	fmt.Println("test4_function.init()")
	c := foo1("abc", 1)
	foo2("abc", 1)
	foo3("abc", 1)
	foo4("abc", 1)
	fmt.Println("c=", c)

}
