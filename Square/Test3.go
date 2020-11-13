package main

import (
	"fmt"
	"log"
)

type Point struct {
	X int
	Y int
}

type Square struct {
	Center Point
	Length int
}

func NewSquare(x int, y int, length int) (*Square, error) {
	// null check
	if length <= 0 {
		return nil, fmt.Errorf("length is negative")
	}
	center := Point{ // new center obj
		X: x,
		Y: y,
	}
	square := &Square{Center: center, Length: length} // new sq object
	/**
	square := &Square{Center: Point{ // new center obj
		X: x,
		Y: y,
	}, Length: length} // new sq object
	**/
	return square, nil
}


// move point inside Square object
func (sq *Square) Move(dx int, dy int) {
	sq.Center.X = sq.Center.X + dx
	sq.Center.Y += dy
}

// move point
func (p *Point) Move(dx int, dy int)  {

	p.X += dx
	p.Y += dy
}

func (sq *Square) InterfaceFun() int {
	area := sq.Length * sq.Length
	return area
}

func main() {
	newSquare, err := NewSquare(1, 3, 0)
	if err != nil {
		log.Fatalf("Error : Can't create square..")
	}

	fmt.Printf("before %+v \n", newSquare)
	newSquare.Move(10, 10)
	fmt.Printf("after %+v \n", newSquare)

	area := newSquare.InterfaceFun()
	fmt.Printf("%+v \n", area)
}
