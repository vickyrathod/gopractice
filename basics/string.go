package main

import (
    "fmt"
)

func main() {
    str := "hello go"
    fmt.Println(str);
    //utf8 index
    fmt.Println(str[0]);

    // int to string
    x := 25;
    str1 := fmt.Sprintf("%d", x);

    // Integer way printing
    fmt.Printf("str: %v (Type : %T) ", str1, str1);

    // String way printing
    fmt.Printf("str: %d (Type : %T) ", str1, str1);
}