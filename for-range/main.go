package main

import (
	"fmt"
	"time"
)

func main() {
	slice := []int{1, 2, 3}
	m := make(map[int]*int)
	var slice2 [3]int
	for index, value := range slice {
		slice = append(slice, value)
		go func() {
			fmt.Println("in goroutine: ", index, value)
		}()
		//time.Sleep(time.Second * 1)
		m[index] = &value
		if index == 0 {
			slice[1] = 11
			slice[2] = 22
		}
		slice2[index] = value
	}
	fmt.Println("slice: ", slice) // [1, 2, 3, 1, 2, 3]
	for key, value := range m {
		fmt.Println("in map: ", key, "->", *value) // 0, 1 1, 2, 2, 3
	}
	fmt.Println("slice2: ", slice2) // 1, 2, 3
	time.Sleep(time.Second * 10)
}
