package main

import (
    "fmt"
)

func main() {
    x := 1

    switch {
        case x == 1:
            fmt.Println("i am one")
            break
        case x == 2:
            fmt.Println("i am two")
        default:
            fmt.Println("i am no one")
    }
}