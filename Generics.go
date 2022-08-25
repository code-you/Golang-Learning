package main

import "fmt"

type Number interface {
	int | int8 | int16 | int32 | int64 | float64 | float32
}

//func newGenerics[age any](myAge age) {
//	fmt.Println(myAge)
//}

//func newGenerics[age int | float64](myAge age) {
//	ages := int(myAge) + 1
//	fmt.Println(ages)
//}

func newGenerics[age Number](myAge age) {
	ages := int(myAge) + 1
	fmt.Println(ages)
}

func main() {
	testAge := 8
	testAge2 := 9.8
	//testAge3 := "New Age"

	newGenerics(testAge)
	newGenerics(testAge2)
	//newGenerics(testAge3)
}
