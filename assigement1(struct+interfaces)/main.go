package main

import "fmt"

//In aceasta problema se cere a se defini 2 structuri una pentru triungi si una pentru patrat
//si o interfata care sa permita utilizarea funtiei print pentru ambele structuri

//definirea interfetei (poate folosi tipul shape orice tip care utilizeaza funtia getArea si returneaza un float64)
type shape interface {
	getArea() float64
}

//definirea unui struct de tip triangle
type triangle struct {
	base   float64
	height float64
}

//definirea unui struct de tip square
type square struct {
	side float64
}

func main() {

	//alimentarea structurilor
	tr := triangle{
		base:   10,
		height: 3,
	}

	sq := square{side: 10}

	//utilizarea funcitilor de print
	fmt.Println("Triangle")
	printArea(tr)
	fmt.Println("Square")
	printArea(sq)

}

//calcularea ariei tringhiului
func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

//calcularea ariei patratului
func (s square) getArea() float64 {
	return s.side * s.side
}

//functia de printare
func printArea(s shape) {
	fmt.Println("area is:", s.getArea())
}
