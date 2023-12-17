package mapStudy

import "fmt"

func MapTest() {

	cityMap := map[string]string{
		"China": "BEIJING",
		"USA":   "WASHINGTON",
		"UK":    "LONDON",
	}
	fmt.Println(cityMap)
	printMap(cityMap)

	fmt.Println("————————————————————————————")
	addMap(cityMap)
	cityMap["USA"] = "DC"
	delete(cityMap, "UK")
	printMap(cityMap)

}

func addMap(myMap map[string]string) {
	myMap["test1"] = "TEST"
}

func printMap(myMap map[string]string) {
	for key, value := range myMap {
		fmt.Println("key=", key)
		fmt.Println("value=", value)
	}
}
func init() {
	fmt.Println("mapStudy.init()")
}
