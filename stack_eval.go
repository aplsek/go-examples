//Sample code to read in test cases:
package main

import "fmt"
import "log"
import "bufio"
import "os"



import (
	
	"strconv"
	"strings"
)

func isOperator(s string) bool {
	if s == "*" || s == "+" || s == "/" {
		return true
	}
	return false
}

func eval(s string, x, y int) int {
	switch s {
		case "*":
			return x * y
		case "+":
			return x + y
		case "/":
			return x / y
	}
	return 0
}

type stack struct {
	val []string
}

func newstack() *stack {
	return &stack{
		val: make([]string,0),
	}
}

func (a *stack) push(ss string) {
	a.val = append(a.val, ss)
}

func (a *stack) pop() string {
	ss := a.val[len(a.val)-1]
	a.val = a.val[:len(a.val)-1]
	return ss
}


func (a *stack) peak() string {
	ss := a.val[len(a.val)-1]
	return ss
}

func (a *stack) len() int {
	return len(a.val)
}



func evalAll(arr []string) string {
	if len(arr) == 0 {
	    return ""
	}
	ss := newstack()
	for _, s := range arr {
		if ss.len() <= 1 {
			ss.push(s)
			continue
		}
		if isOperator(ss.peak()) {
			ss.push(s)
			continue
		}
		nn := ss.pop()
		op := ss.pop()
		if !isOperator(op) {
			fmt.Println("error: operator expected")
			return ""
		}
		r := eval(op, toInt(nn), toInt(s))
		ss.push(fmt.Sprintf("%v",r))
	}
	return ss.pop()
}

func toInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
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
        arr := strings.Split(scanner.Text(), " ")
        fmt.Println( evalAll(arr))
    }   
}
