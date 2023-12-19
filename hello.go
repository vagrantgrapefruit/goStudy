package main

import (
	"fmt"

	"github.com/vagrantgrapefruit/goStudy/classStudy"
	// "_" means import but drop the value
	// "." in import means the func import into this file
	//_ "github.com/vagrantgrapefruit/goStudy/deferLearn"
	//_ "github.com/vagrantgrapefruit/goStudy/point"
	//_ "github.com/vagrantgrapefruit/goStudy/test3"
	//_ "github.com/vagrantgrapefruit/goStudy/test4"
	//_ "github.com/vagrantgrapefruit/goStudy/sliceStudy"
	//_ "github.com/vagrantgrapefruit/goStudy/mapStudy"
	//_ "github.com/vagrantgrapefruit/goStudy/structStudy"
)

func main() {
	fmt.Println("______START______")
	//If you want to export func you need to upcase the first character
	/*test3.TestConst()
	test4.Foo4("1", 1)
	m := 100
	n := 99
	fmt.Println("befor m=", m, "n=", n)
	point.Swap(&m, &n)
	fmt.Println("Later m=", m, "n=", n)

	//fmt.Println("a=", a)
	sliceStudy.TestSlice()
	sliceStudy.SliceAdd()
	mapStudy.MapTest()
	structStudy.StructStudy()
	classStudy.ClassStudy()*/
	classStudy.HumanTest()
	fmt.Println("______END______")
	countProfit(1000, 10, 180000, 18, 28)
}

func countProfit(monthin int, yearIn int, getMoney int, yearOutStart int, yearOutEnd int) float64 {
	var profit float64

	println("Don`t finish")
	return profit
}
