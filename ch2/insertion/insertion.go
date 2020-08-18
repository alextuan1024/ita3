package main

import "fmt"

func main() {
	input := []int{5, 2, 4, 6, 1, 3}
	Insertion(input)
	fmt.Println(input)
}

func Insertion(array []int) {
	length := len(array)
	for j := 1; j < length; j++ {
		key := array[j]
		var i int
		for i = j - 1; i > -1 && array[i] > key; i-- {
			array[i+1] = array[i]
		}
		array[i+1] = key
	}
}
