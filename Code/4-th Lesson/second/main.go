package main

import "fmt"

func main() {
	/*level2:
	for j := 1; j <= 20; j++ {
		for y := 1; y <= 8; y++ {
			fmt.Println("J:", j, "Y:", y)
			if j >= 10 {
				break level2
			}
		}
	}

		Level:
			for f := 1; f <= 20; f++ {
				for p := 1; p <= 10; p++ {
					fmt.Println("f:", f, "P:", p)
					if f >= 10 {
						break Level
					}
				}
			}// - break
	*/
	//		|||||||||||||||||||||||||||||||||||||||||||||

Level:
	for f := 1; f <= 20; f++ {
		for p := 1; p <= 10; p++ {
			fmt.Println("f:", f, "P:", p)
			if f >= 10 {
				continue Level
			}
		}
	} //  - continue

	//		 ||||||||||||||||||||||||||||||

	for i := 0; i <= 20; i++ {
		if i%2 == 1 {
			continue
		}
		fmt.Println(i)
	} //- continue how work

}
