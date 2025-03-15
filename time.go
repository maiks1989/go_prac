package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	t1 := time.Now().UnixNano()
	fmt.Println(reflect.TypeOf(t1))
	t2 := time.Now().UnixNano()
	fmt.Println(t2 - t1)

}
