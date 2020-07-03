package main

import (
    "fmt"
)

func main(){
    // Map (dont forget last comma)
    mp := map [string]float64{
        "vicky": 56.89,
        "rathod": 45.938,
        "poonam": 8399.834845,
        "makwana": 37487.4884,
    }

    fmt.Println(mp["vicky"])
}