package sliceStudy

import "fmt"

func printArray(myArray [4]int) {
	//Value Copy  Can`t change original data
	for index, value := range myArray {
		fmt.Println("index=", index, ",value=", value)
	}
	myArray[0] = -11
}

func TestArray() {
	var arry1 [10]int
	arry2 := [10]int{1, 2, 3, 4}
	arry3 := [4]int{1, 2, 3, 4}

	for index, value := range arry2 {
		fmt.Println("index=", index, ",value=", value)
	}
	//Show ArrayType
	fmt.Printf("arry1 types= %T\n", arry1)
	fmt.Printf("arry2 types= %T\n", arry2)
	fmt.Printf("arry3 types= %T\n", arry3)
	printArray(arry3)
}

func TestSlice() {
	myArray := []int{1, 2, 3, 4} //myArray is a Slice Not Array
	fmt.Printf("myArray types= %T\n", myArray)

	slice1 := []int{1, 2, 3}
	fmt.Printf("len = %d slice1 = %v\n", len(slice1), slice1)

	var slice2 []int
	fmt.Printf("len = %d slice2 = %v\n", len(slice2), slice2)
	//  ```slice2[0] = 1``` will panic cause the length of slice2 is zero
	slice2 = make([]int, 3) //make slice2
	slice2[0] = 1
	//After Make ,```slice2[0] = 1``` can run correctly
	fmt.Printf("After Make ,Slice2 len = %d slice2 = %v\n", len(slice2), slice2)

	var slice3 []int = make([]int, 3) //equals slice3:=make([]int, 3)
	fmt.Printf("len = %d slice3 = %v\n", len(slice3), slice3)

	var slice []int
	if slice == nil {
		fmt.Println("slice is Empty")
	} else {
		fmt.Println("slice is not Empty")
	}
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
func SliceAdd() {
	//make a slice which len=3 and cap=5
	//slice=[0,0,0]
	slice := make([]int, 3, 5)
	fmt.Print("Firstly ")
	printSlice(slice)

	//add a value to slice,then slice = [0,0,0,1]  cap=5
	slice = append(slice, 1)
	fmt.Printf("After Add one value ")
	printSlice(slice)

	//add a value to slice,then slice = [0,0,0,1,2]  cap=5
	slice = append(slice, 2)
	fmt.Printf("After Add another value ")
	printSlice(slice)

	//add a value to a full slice ,the cap will double the value of cap
	// slice = [0,0,0,1,2,3]  cap=10
	slice = append(slice, 3)
	fmt.Printf("After add a value to a full slice ")
	printSlice(slice)

	slice2 := make([]int, 3) //cap(slice2)=3
	fmt.Printf("Firstly slice2 ")
	printSlice(slice2)
	slice2 = append(slice2, 1) //cap(slice2)=6
	fmt.Printf("add one value slice2 ")
	printSlice(slice2)
	slice2 = append(slice2, 1) //cap(slice2)=6
	slice2 = append(slice2, 1) //cap(slice2)=6
	slice2 = append(slice2, 1) //cap(slice2)=12
	fmt.Printf("add another 3 value slice2 ")
	printSlice(slice2)
	copy(slice, slice2) //copy just copy the value ,don`t change len and cap
	fmt.Printf("After copy slice: ")
	printSlice(slice)
}

func init() {
	fmt.Println("slicestudy.init()")
}
