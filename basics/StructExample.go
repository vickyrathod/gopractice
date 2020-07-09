package main

import (
	"fmt"
)

type Trade struct {
	name string
	age int
	speed float64
}

func main(){
	t1 := Trade{"vicky", 35, 343.56};
	// Without formating
	fmt.Println(t1)

	// With formating
	fmt.Printf("formated printing: %+v", t1)

	// Creating empty struct
	t2 := Trade{}
	fmt.Printf("empty struct %+v", t2)
}