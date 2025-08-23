/*
*
Variable 'minVal' = array[0]
For each element in the array

	If current element < minVal
	    minVal = current element
*/
package main

import "fmt"

func main() {
	myAry := []int{7, 12, 9, 4, 11}
	minVal := myAry[0]
	for _, v := range myAry {
		if v < minVal {
			minVal = v
		}
	}
	fmt.Println("Lowest Value:", minVal)
}
