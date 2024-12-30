package main

import (
	"fmt"
)

func merge(left, right []int) []int {
	res := make([]int, 0, len(left)+len(right))

	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			res = append(res, left[i])
			i++
		} else {
			res = append(res, right[j])
			j++
		}
	}

	// adding remaining elements
	res = append(res, left[i:]...)
	res = append(res, right[j:]...)

	return res
}

func mergeSort(arr []int, ch chan []int) {
	if len(arr) <= 1 {
		ch <- arr
		return
	}
	mid := len(arr) / 2

	leftCh := make(chan []int)
	rightCh := make(chan []int)

	go mergeSort(arr[:mid], leftCh)
	go mergeSort(arr[mid:], rightCh)

	left, right := <-leftCh, <-rightCh

	close(leftCh)
	close(rightCh)

	ch <- merge(left, right)
}

func main() {
	arr := []int{9, 4, 3, 6, 1, 2, 10, 5, 7, 8}
	ch := make(chan []int)

	go mergeSort(arr, ch)

	sortedArr := <-ch
	fmt.Println("Original Array:", arr)
	fmt.Println("Sorted Array:", sortedArr)
}
