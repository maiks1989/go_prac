package main

import (
	"fmt"
	"time"
)

func main() {
	ns := 5000000000 // 5秒分のナノ秒
	seconds := int64(time.Duration(ns) / time.Second)

	fmt.Println(seconds) // 5
}
