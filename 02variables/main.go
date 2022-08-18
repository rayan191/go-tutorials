package main

import "fmt"

var MaxAge = 25

func main() {
	var username string = "Rayan"
	fmt.Println(username)
	fmt.Printf("value of the variable is of type %T \n", username)

	var isSingnin bool = true
	fmt.Println(isSingnin)
	fmt.Printf("value of the variable is of type %T \n", isSingnin)

	var age int = 24
	fmt.Println(age)
	fmt.Printf("value of the variable is of type %T \n", age)

	var ages bool
	fmt.Println(ages)

	numberOfUsers := 30000
	fmt.Println(numberOfUsers)

	fmt.Println(MaxAge)

}
