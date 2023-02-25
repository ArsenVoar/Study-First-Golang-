package main

import "fmt"

func main() {
	age := 26

	if age == 16 || age == 20 || age == 26 {
		fmt.Println("Yes")
	}
	/*
	 	age := 16
	   	if age >= 6 && age <= 18 {
	   		fmt.Println("You are young(third)")
	   	} else {
	   		fmt.Println("You are an adult")
	   	}//   - how work &&
	*/
	//	   ||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||

	age = 20
	if age < 18 {
		fmt.Println("You are young(third)")
	} else {
		fmt.Println("You are an adult")
	} //   - if else how work

	//	   				    ||||||||||||||||||||||||||||||||||||||||||||||||||||||||||

	if age < 18 {
		fmt.Println("You are young(second)")
	}

	if Child := Children(age); Child == true {
		fmt.Println("You are young(first)")
	} //  --- if - how work

	//	   			         ||||||||||||||||||||||||||||||||||||||||||||||||

	/*   					age := 16

	if age < 18 {
		fmt.Println("You are young")
	}//   ----  if - how work*/
}

func Children(age int) bool {
	return age < 18
}
