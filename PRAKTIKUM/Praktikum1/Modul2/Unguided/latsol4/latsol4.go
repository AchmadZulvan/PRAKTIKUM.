package main

import "fmt"

func main() {
	var farenheit float64
	fmt.Scan(&farenheit)
	celcius := (5.0 / 9.0) * (farenheit - 32)
	fmt.Println(celcius)

}
