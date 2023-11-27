package main 

import "fmt"

// keyword iota can just used in const
const (
	//keyword  iota can used in const() ,every line`s iota will add 1, default iota in line 1 is 0
	BEIJING = 10*iota
	SHANGHAI
	SHENZHEN
)

const (
	a,b=iota+1,iota+2
	c,d
	e,f

	g,h=iota*2,iota*3
	i,k
)
func main(){
	const length int = 10
	fmt.Println("length = ",length)

	fmt.Println("BEIJING=",BEIJING)
	fmt.Println("SHANGHAI=",SHANGHAI)
	fmt.Println("SHENZHEN=",SHENZHEN)
	fmt.Println("a,b=",a,b)
	fmt.Println("c,d=",c,d)
	fmt.Println("e,f=",e,f)
	fmt.Println("g,h=",g,h)
	fmt.Println("i,k=",i,k)
}

