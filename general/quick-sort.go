package main

import "fmt"

func partition(arr []int, low, high int) int {
	fmt.Println("Partition called with:", low, high, arr)
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	fmt.Println("Swaped array:", arr)
	return i + 1
}

func quickSort(arr []int, low, high int) {
	fmt.Println("QuickSort called with:", low, high, arr)
	if low < high {
		pivotIndex := partition(arr, low, high)
		fmt.Println("Pivot found at:", pivotIndex, "with value:", arr[pivotIndex])
		quickSort(arr, low, pivotIndex-1)
		quickSort(arr, pivotIndex+1, high)
	}
}

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90, 5}
	n := len(arr)
	quickSort(arr, 0, n-1)
	fmt.Println("Sorted array:", arr)
}
