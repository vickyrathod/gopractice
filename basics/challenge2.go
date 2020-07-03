package main

import (
   "fmt"
   "strings"
)

func main(){
    str := "hi hello how are you, have a nice day. i have many changes to do, Whats your status."
    arr := strings.Split(str, " ")

    mp := make(map[string]int)

    for _, value := range arr {
        mp[value] = mp[value] + 1
    }

    for key, value := range mp {
            fmt.Printf("%v : %d", key, value)
        }
}