package main

import "fmt"

//function to check if number is even or odd.
func evenOdd(num int) string {
	if num%2 == 0 {
		return "Even"
	}
	return "Odd"
}

//calling the function
func main() {
	result := evenOdd(6)
	fmt.Println(result)
}
