package main

import "fmt"

func main() {
	num := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10} //make integer slice
	for _, n := range num {
		if n%2 == 0 { //if remainder is 0 it is even number except 0 i think
			fmt.Println(n, " is even") //print out result
		} else {
			fmt.Println(n, " is odd") //print out result
		}
	}
}
