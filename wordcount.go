package main


import (
	"fmt"
	"strings"
	//"math/rand"
)

var m map[string]int 


func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	count := map[string]int{}
	for _,word := range words {
		count[word]++
	}
	return count
}

func add (word string) {
	elem, ok := m[word]
	if ok == true {
		m[word] = elem + 1
	} else {
		m[word] = 1
	}
}


func main() {
    m = make(map[string]int)

    add("aa")
    add("aa")
    add("aa")
    add("cc")
    add("bb")
    

    fmt.Printf("%#v\n",m)
    var wc = WordCount("ss ss aa ss aa jj")

    fmt.Printf("%#v\n",wc)
}
