package main

import (
	"fmt"
)

type Pointx struct{
	x int
	y int
}

//Pass by reference 
func shift(p *Pointx) {
	p.x += 2
	p.y += 2
}

// Pass by value and return
func shiftVal(p Pointx) (Pointx) {
	p.x += 2;
	p.y += 2;

	return p;
}

func main(){
	p1 := Pointx{1,2}
	shift(&p1);
	fmt.Println(p1)

	p1 = shiftVal(p1)

	fmt.Println(p1)
}