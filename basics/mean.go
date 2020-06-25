// very first code
package main

import (
    "fmt"
)

func main() {
    x := 1
    y := 2
   fmt.Printf("val : %v, Type : %T\n", x, y)

   mean := (x + y) / 2

   fmt.Printf("mean : %v", mean)
}