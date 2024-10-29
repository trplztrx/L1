package main

import (
	"fmt"
	"math"
)

// Адаптер вычисляет наименьший радиус окружности,
// в которую можно вписать квадратный колышек,
// и представляет его как круглый колышек с этим радиусом.

type RoundHole struct {
	radius int
}

func NewRoundHole(radius int) *RoundHole {
	return &RoundHole{
		radius: radius,
	}
}

func (rc *RoundHole) GetRadius() int {
	return rc.radius
}

func (rc *RoundHole) Fits(peg *RoundPeg) bool {
	return rc.GetRadius() == peg.GetRadius()
}

type RoundPeg struct {
	radius int
}

func NewRoundPeg(radius int) *RoundPeg {
	return &RoundPeg{
		radius: radius,
	}
}
func (rp *RoundPeg) GetRadius() int {
	return rp.radius
}

type SquarePeg struct {
	width int
}

func NewSquarePeg(width int) *SquarePeg {
	return &SquarePeg{
		width: width,
	}
}

func (sq *SquarePeg) GetWidth() int {
	return sq.width
}

type SquarePegAdapter struct {
	peg *SquarePeg
}

func NewSquarePegAdapter(peg *SquarePeg) *SquarePegAdapter {
	return &SquarePegAdapter{
		peg: peg,
	}
}

func (spa *SquarePegAdapter) GetRadius() int {
	return int(float64(spa.peg.GetWidth()) * math.Sqrt(2) / 2)
}

func main() {
	rPeg := NewRoundPeg(1)
	hole := NewRoundHole(1)

	fmt.Println(hole.Fits(rPeg))

	sqPeg := NewSquarePeg(2)
	adapter := NewSquarePegAdapter(sqPeg)

	fmt.Println(hole.Fits(NewRoundPeg(adapter.GetRadius())))
}