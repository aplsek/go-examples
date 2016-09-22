package main

import (
	"fmt"
	
	"time"
)

func main() {
	fmt.Println("Hello, playground")
	
	fmt.Println("ts:%v", ts)
	fmt.Println("nanos:%v", ts.UnixNano())
	fmt.Println("millis:%v", dropNanos(ts.UnixNano()))
	mm := dropNanos(ts.UnixNano())
	fmt.Println("millis:%v", time.Unix(0, mm*int64(time.Nanosecond)))
}
