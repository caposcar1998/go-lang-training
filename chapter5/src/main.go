package main

import "fmt"

func main() {
	oxford := CreateCity(51.45, 1.15)
	turin := CreateCity(45.04, 7.42)
	athens := CreateCity(37.983972, 23.727806)
	amsterdam := CreateCity(52.366667, 4.9)
	berlin := CreateCity(52.516667, 13.388889)

	fmt.Println("Oxford-turin: ")
	CalculateHarvesineDistance(oxford, turin)
	fmt.Println()

	fmt.Println("Oxford-turin-manual: ")
	oxford.CalculateHarvesineDistanceManual(turin)
	fmt.Println()

	fmt.Println("Turin-oxford-manual: ")
	turin.CalculateHarvesineDistanceManual(oxford)
	fmt.Println()

	fmt.Println("Athens-amsterdam-manual: ")
	athens.CalculateHarvesineDistanceManual(amsterdam)
	fmt.Println()

	fmt.Println("Amsterdam-berlin-manual: ")
	amsterdam.CalculateHarvesineDistanceManual(berlin)
	fmt.Println()

	fmt.Println("Berlin-athens-manual: ")
	berlin.CalculateHarvesineDistanceManual(athens)
}
