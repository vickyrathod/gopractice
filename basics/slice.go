package main

import (
    "fmt"
)

func main() {
    arr :=[]string{"vicky", "rathod", "poonam", "makwana"}

    // Interation by index
    for i:=0; i<len(arr); i ++ {
        fmt.Println(arr[i]);
    }

    // Interation by index value with range
    for i, value := range arr {
        fmt.Println(i, value);
    }

    // Interation by index with range
    for i := range arr {
        fmt.Println(arr[i]);
    }

    // Interation by value with range
    for _, value := range arr {
        fmt.Println(value);
    }
}