package main

import "fmt"

func main() {
	// Two different ways of declaring an array
	var cars = [...]string{"Ford","Cheverolet"}
	cars2 := [...]string{"Ferrari","Renault","Mercedes"}
	fmt.Println("Emp:", cars)

	var cars2len = len(cars2)
	// Arrays can't grow....what was I thinking?!?!?
	cars2[cars2len + 1] = "Rover"


	//cars = [...]string{"Ford",2:,"Chevrolet"}
	fmt.Println("Cars:", cars)
	fmt.Println("Cars2:", cars2)
	fmt.Println("Length of array is ", len(cars))
	fmt.Println("Length of cars2 is ", len(cars2))
}