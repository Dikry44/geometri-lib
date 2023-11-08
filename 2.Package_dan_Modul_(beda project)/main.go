package main

import (
	"fmt"
	"geometri-lib/shape"
	//ngambil struct dan fungsi dari direktori shape
)

func main() {
	firstFloor := shape.Rectangle{Width: 7.5, Length: 6.5}
	secondFloor := shape.Rectangle{Width: 4.5, Length: 5.5}

	totalArea := firstFloor.Area() + secondFloor.Area()
	fmt.Println("Total Area :", totalArea)
}
