package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(fmt.Sprintf("%v", now.Format("Mon Jan 2 15:04:05 MST 2006")))
	fmt.Println(fmt.Sprintf("%v", now.Format("2006-01-02")))
}
