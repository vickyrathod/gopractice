package main

import (
    "fmt"
)

//Pass by reference using pointer
func double(val *int){
    *val = 2 * *val;
}

func main(){
    val := 5;
    double(&val);
    fmt.Println(val);
}