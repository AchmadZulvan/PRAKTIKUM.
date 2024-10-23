package main

import (
	"fmt"
	"math"
)

func main() {
	var r , t, n int
	var volume float64
	fmt.Scan(&n)
	
	for i := 1; i <= n; i++ {
		
		fmt.Scan(&r, &t)
		volume = (1.0/3.0) * math.Pi * math.Pow(float64(r), 2) * float64(t)
		fmt.Print(volume)
	}
}
