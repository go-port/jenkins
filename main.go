package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i < 30; i++ {
		time.Sleep(time.Second)
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "Hello Jenkins!")
	}
}
