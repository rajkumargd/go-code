/*
*
myArray to sort
Outer For each to countrol inner loop count
Inner Loop to Swaps Value
*/
package main

import "fmt"

func main() {
	myAry := []int{64, 34, 25, 12, 22, 11, 90, 5}
	n := len(myAry)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if myAry[j] > myAry[j+1] {
				myAry[j], myAry[j+1] = myAry[j+1], myAry[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}

	}
	fmt.Println("Sorted ary:", myAry)
}
