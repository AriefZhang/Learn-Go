package main

import "fmt"
import "strconv"

var num = [2]int{21, 31}

var x = len(num)

func main() {
    var count = 0
    
    for i := num[0]; i < num[1]; i++ {
        var check = 0
        var str = strconv.Itoa(i)

        for j := 0; j < len(str); j++ {
            if str[check] == str[len(str) - 1 - check] {
                check++
            }
        }
        if check == len(str) {
            count++
        }
	}

    fmt.Println(count)
}
