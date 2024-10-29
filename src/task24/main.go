package main

import (
	"fmt"
	"math"
)


type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func (p *Point) GetCords() (float64, float64) {
	return p.x, p.y
}


func GetDistance(p1, p2 *Point) float64 {
	x1, y1 := p1.GetCords()
	x2, y2 := p2.GetCords()
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}
func main() {
	p1 := NewPoint(2, 2)
	p2 := NewPoint(4, 4)

	fmt.Println(GetDistance(p1, p2))
}