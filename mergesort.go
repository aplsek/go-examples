package main

import "fmt"


func merge(a, b []int) []int {
	size :=  len(a) + len(b)
	res := make([]int, size)

	k:=0
	l:=0

	for i := 0; i < size; i++ {
		if k >= len(a) {
			y := b[l]
			res[i] = y
			l++
			continue
		} else if l >= len(b) {
			x := a[k]
			res[i] = x
			k++
			continue
		}
		if a[k] > b[l] {
			res[i] = b[l]
			l++
		} else {
			res[i] = a[k]
			k++
		}
	}
	return res
}

func mergesort(arr []int)[]int {
	if len(arr) <= 1 {
		return arr 
	} 
	idx := len(arr)/2
	
	a := arr[0:idx]
	b := arr[idx:]
	fmt.Printf(" a:%v, b:%v\n", a, b)
	
	a1 := mergesort(a)
	b1 := mergesort(b)
	
	return merge(a1, b1)
}

func main() {
	arr := []int{4,5,7,2,3,1,8,9}
	fmt.Printf("res: %v\n", mergesort(arr))
	
}
