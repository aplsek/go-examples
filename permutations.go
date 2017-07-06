package main

import (
	"fmt"
	"strings"
	"sort"
	
	
)

func permut(arr string ) []string {
	if len(arr) == 1 {
		return []string{arr}
	}
	c := arr[0:1]
	permss := permut(arr[1:])
	res := make([]string,0)
	for _, ss:= range permss {
		for i := 0; i < len(ss); i++ {
			res = append(res, fmt.Sprintf("%v%v%v", ss[0:i],c, ss[i:]))
		}
		//res = append(res, fmt.Sprintf("%v%v", c, ss))
		res = append(res, fmt.Sprintf("%v%v", ss, c))
	}
	return res
}

func permutate(str string) {
	perms := permut(str)
	sort.Strings(perms)
	s := strings.Join(perms, ",")
	fmt.Printf("%v\n", s)
}

func main() {
	permutate("hat")
}
