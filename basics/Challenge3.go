package main

import (
    "fmt"
    "net/http"
)

func getContentType(url string) (string, error) {
    // make http call
    resp, err := http.Get(url)
    if(err != nil){
        return "", err;
    }
    // Make sure we close the Body at end
    defer resp.Body.Close()

    ctype := resp.Header.Get("Content-Type")
    // If content tyoe is empty
    if ctype == "" {
        return "", fmt.Errorf("No content type available")
    }

    return ctype, nil
}

func main(){
   ctype, err := getContentType("http://google.com/")
   if(err != nil){
        fmt.Printf("NO content type as error: %s\n", err);
   } else {
        fmt.Println(" content type is : ", ctype)
   }
}