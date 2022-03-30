package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

func main() {
	//Insert Code Here
	var height, weight string
	var bmi float64
	fmt.Println("Please enter your weight(kg): ")
	fmt.Scanln(&weight)
	fmt.Println("Please enter your height(m): ")
	fmt.Scanln(&height)

	h, errh := strconv.ParseFloat(height, 64)
	w, errw := strconv.ParseFloat(weight, 64)

	if errh != nil && errw != nil {
		log.Fatal(errh, errw)
	} else {

		bmi = float64(w) / math.Pow(float64(h), 2)

		switch {
		case bmi < 18.5:
			fmt.Println("You are underweight!")
		case bmi > 18.5 && bmi < 25:
			fmt.Println("Your weight is healthy.")
		case bmi >= 25 && bmi < 30:
			fmt.Println("You are overweight!")
		case bmi >= 30 && bmi < 35:
			fmt.Println("You are obese!")
		case bmi > 35 && bmi < 40:
			fmt.Println(("You are severly obese"))
		case bmi >= 40:
			fmt.Println("You are morbidly obese")
		default:
			fmt.Println("Other weight categories")
		}
	}

}
