package main

import "fmt"
import "math"

// https://projecteuler.net/problem=1

func multiples(x float64, y float64, max int) int {
    sum := 0
    for i := 1; i < max; i++ {
       // var m = math.Mod(float64(i),x)
        //fmt.Println("mod is %d", m)
        if math.Mod(float64(i),x) == 0  || math.Mod(float64(i),y) == 0 {
        	sum += i
        }
    }
    return sum  
}

func pln (str string) {
    fmt.Println(str)
}

func main() {
    pln("start");
    var res = multiples(float64(3),float64(5),1000)
    fmt.Println("result = ", res)
}
