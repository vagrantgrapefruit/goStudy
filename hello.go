package main

import (
	"fmt"

	// "_" means drop the value
	_ "github.com/vagrantgrapefruit/goStudy/test3"
	// "." in import means the func import into this file
	"github.com/vagrantgrapefruit/goStudy/deferLearn"
	_ "github.com/vagrantgrapefruit/goStudy/point"
	_ "github.com/vagrantgrapefruit/goStudy/test4"
)

func main() {
	fmt.Println("start")
	//If you want to export func you need to upcase the first character
	/*test3.TestConst()
	test4.Foo4("1", 1)
	m := 100
	n := 99
	fmt.Println("befor m=", m, "n=", n)
	point.Swap(&m, &n)
	fmt.Println("Later m=", m, "n=", n)*/
	a := deferLearn.Defer_call()
	fmt.Println("a=", a)
}
