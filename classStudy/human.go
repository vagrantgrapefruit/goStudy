package classStudy

import "fmt"

type Human struct {
	Name   string
	Gender string
}

func (huamn *Human) SayHello() {
	fmt.Println("hello")
}

func (huamn *Human) Run() {
	fmt.Println("Human Run...")
}

type SuperMan struct {
	Human      //SuperMan Extented Human`s proptype and function
	SuperPower string
}

func (sm *SuperMan) SayHello() {
	fmt.Println("SuperMan SayHello")
}

func (sm *SuperMan) ChangePower() {
	sm.SuperPower = "SuperPower"
}

func HumanTest() {
	h := Human{"zhangsan", "male"}
	h.SayHello()
	h.Run()

	sm := SuperMan{h, "Powfer"}
	sm.SayHello()
	sm.Run()
	sm.ChangePower()
	fmt.Println(sm.SuperPower, sm.Name)
}
func init() {
	fmt.Println("classhuman init")
}
