package main

import (
	"fmt"
)

func main() {

	/*	var number1 int8 = 8
		var number2 int64 = 8
		fmt.Println(unsafe.Sizeof(number1))
		fmt.Println(unsafe.Sizeof(number2)) //  -  diference between int8 and int 64
		//int8 = 1 bait
		//int64 =  8 bait */
	//		    |||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
	/*						   var number1 int = 31
						   	number2 := 1

						   	fmt.Println(number1 + number2)// - plus var

	//				      |||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||

	var number1 int = 31
	number2 := 1

	fmt.Println(number1 + number2) // - willn't working because difirent int (int64"number1" and int"number2")

	//	    |||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
	*/
	name := "arsen"
	hello := fmt.Sprintf("hello %s", name)
	fmt.Println(hello)

	fmt.Printf("Type: %T Value: %s\n", name, name) // - plus var

}
