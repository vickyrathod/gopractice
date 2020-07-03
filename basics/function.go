package main

import (
    "fmt"
)

func adder(a int, b int) int {
    return a + b;
}

func divMod(a int, b int) (int , int) {
    return a/b, a%b;
}

func main(){
    div, mod := divMod(45, 3);
    fmt.Printf("add :%d, div: %d, mod: %d", adder(45, 56), div, mod)
}