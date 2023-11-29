package deferLearn

import "fmt"

func foo1() {
	fmt.Println("foo1")
}

func foo2() {
	fmt.Println("foo2")
}

func foo3() {
	fmt.Println("foo3")
}

func foo4() int {
	fmt.Println("foo4")
	return 1
}

func Defer_call() int {
	defer foo1()
	defer foo2()
	defer foo3()
	fmt.Println("Start Defer_call")
	return foo4()
}

func init() {
	fmt.Println("defer.init()")
}
