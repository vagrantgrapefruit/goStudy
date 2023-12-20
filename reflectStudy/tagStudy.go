package reflectStudy

import (
	"fmt"
	"reflect"
)

type Resume struct {
	Name   string `info:"name" doc:"我的名字"`
	Age    int    `info:"age"`
	Gender string `info:"gender"`
}

func readTag(arg interface{}) {
	t := reflect.TypeOf(arg).Elem()

	for i := 0; i < t.NumField(); i++ {
		tagInfo := t.Field(i).Tag.Get("info")
		tagDoc := t.Field(i).Tag.Get("doc")
		fmt.Println("info:", tagInfo, "doc:", tagDoc)
	}
}
func TagStudy() {
	fmt.Println("TagStudy Start")
	var re Resume
	readTag(&re)
	fmt.Println("TagStudy End")
}

func init() {
	fmt.Println("tagStudy init")
}
