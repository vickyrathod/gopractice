package main

import (
    "fmt"
)

func main(){
    for i := 0; i < 20; i ++ {
        fmt.Println(i);
    }

    fmt.Println("______");

    x := 0;
    for x < 20 {
        fmt.Println(x);
        x ++;
    }
}