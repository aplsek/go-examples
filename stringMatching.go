package main

import "fmt"



func pln (str string) {
    fmt.Println(str)
}

func stringmatch(str string, p string) bool {
    if str == p {
      return true
    }
    if len(str) < len(p) {
        return false
    }
    var len1 = len(str)
    //var len2 = len(p)
    //fmt.Println("len is %d", len1)
    
    var len2 = len(p)
    var idx = 0
    for i := 0 ; i < len1  ; i++ {
        //fmt.Println("sub is = %s", str[i])
        //fmt.Println("    p is = %s, idx=%d, len=%d", p[idx], idx, len2)
        if (str[i] == p[idx]) {
            if idx == len2-1 {
                fmt.Println("match found! i=%d ", i)
             	return true
            } else {
                idx++
            }
        } else {
          idx = 0
          
        }
    }
    
    return false   
}

func main() {
    pln("start");
    var res = stringmatch("hello ales xxx", "ales")
    res = stringmatch("abracadabra", "cad")
    fmt.Println("result = ", res)
    
    
}
