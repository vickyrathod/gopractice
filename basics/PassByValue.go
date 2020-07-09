package main

import (
"fmt"
)

// Pass By values
func double(a int) (int) {
    return 2 * a
}

// Pass by reference
func addValue(values []int, val int) ([]int) {
    for i := range values {
        values[i] += val;
    }

    return values;
}

func main() {
    values := []int {34, 637, 53, 78, 325}
    val := 10

    double(val)
    addValue(values, val);

    fmt.Println(values)
}