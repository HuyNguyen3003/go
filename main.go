package main

import "fmt"


func main() {
	
	
	arr := []int{1, 2, 3, 4, 5}
	arr = append(arr, 6)
	
	
	fmt.Println(cap(arr))
	
}