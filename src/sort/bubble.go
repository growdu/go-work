package main

import "fmt"

func main() {
	examples := []int{64, 23, 54, 45, 56, 798, 903}
	fmt.Println("examples:")
	fmt.Println(examples)
	bubble_sort(examples)
	println("after sorted:")
	fmt.Println(examples)
}

/*
每次交换相邻位置的元素
比较的是第二层循环j和j+1位置的值
*/
func bubble_sort(values []int) {
	for i := 0; i < len(values)-1; i++ {
		for j := 0; j < len(values)-1-i; j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
			}
		}
	}
}
