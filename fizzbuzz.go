//Sample code to read in test cases:
package main

import "log"
import "bufio"
import "os"
import (
	"fmt"
	"strings"
	"strconv"
)


func bizbaz(str string) {
	arr := strings.Split(str, " ")
	f, _ := strconv.Atoi(arr[0])
	b, _ := strconv.Atoi(arr[1])
	n, _ := strconv.Atoi(arr[2])
	
	for i:= 1; i <= n ; i++ {
		match := false
		if i % f == 0 {
			match = true
		    fmt.Printf("F")
		} 
		if i % b == 0 {
		match = true
		    fmt.Printf("B")
		}
		if !match {
		  	fmt.Printf("%v", i)
		}	
		
		fmt.Printf(" ")
	}
	fmt.Printf("\n")
}

func main() {
    file, err := os.Open(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }   
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        //'scanner.Text()' represents the test case, do something with it
        //fmt.Println(scanner.Text())
        bizbaz(scanner.Text())
    }   
}
