package main

import (
	"fmt"
	"math"
)

var satLocations = [3][2]float32{
	{-500.0, -200.0},
	{100.0, -100.0},
	{500.0, 100.0},
}

func main() {
	var x float32
	var y float32
	var distance1 float32
	var distance2 float32
	var distance3 float32

	fmt.Println("Insert point location (x)(y)")
	fmt.Scanln(&x)
	fmt.Scanln(&y)

	distance1 = float32(math.Sqrt(math.Pow(float64(satLocations[0][0]-x), 2) + math.Pow(float64(satLocations[0][1]-y), 2)))
	distance2 = float32(math.Sqrt(math.Pow(float64(satLocations[1][0]-x), 2) + math.Pow(float64(satLocations[1][1]-y), 2)))
	distance3 = float32(math.Sqrt(math.Pow(float64(satLocations[2][0]-x), 2) + math.Pow(float64(satLocations[2][1]-y), 2)))

	fmt.Println("Distances:", distance1, distance2, distance3)
}
