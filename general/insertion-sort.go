package main

import "fmt"

func insertionSort(myArray []int) {

	n := len(myArray)

	for i := 1; i < n; i++ {

		insertIndex := i
		currentValue := myArray[i]
		fmt.Println("Outer Loop", i)
		for j := i - 1; j >= 0; j-- {
			fmt.Println("Inner Loop", j)

			if myArray[j] > currentValue {
				myArray[j+1] = myArray[j]
				insertIndex = j
			}
		}

		myArray[insertIndex] = currentValue
	}
	fmt.Println("Sorted Array:", myArray)
}

func main() {
	myArray := []int{12, 2, 64, 34, 25, 12, 22, 11, 90, 5}
	insertionSort(myArray)
}
