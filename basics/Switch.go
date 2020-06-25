package main

import (
    "fmt"
)

func main() {
    x := 3

    switch x {
        case 1:
            fmt.Println("i am one")
            break;
        case 2:
            fmt.Println("i am two")
        default:
            fmt.Println("i am no one")
    }
}