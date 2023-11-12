/*

№ 24 Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.

*/

package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func main() {
	point1 := create_point(0.0, 8.0)
	point2 := create_point(8.0, 0.0)

	fmt.Println(find_distance(point1, point2))
}

func create_point(x float64, y float64) *Point {
	return &Point{x, y}
}

func find_distance(point1, point2 *Point) float64 {
	return math.Sqrt(math.Pow((point1.x-point2.x), 2) + math.Pow((point1.y-point2.y), 2))
}
