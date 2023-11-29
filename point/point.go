package point

import "fmt"

//use point to swap
func Swap(a, b *int) {
	*a, *b = *b, *a
}

func init() {
	fmt.Println("point.init()")
}
