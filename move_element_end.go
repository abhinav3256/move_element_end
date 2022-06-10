package main

import "fmt"

func main() {
	arr := []int{5, 1, 48, 52, 1, 5, 1}

	fmt.Println(arr)
	result := processArray(arr)

	fmt.Println(result)

}

func processArray(arr []int) []int {
	var i int
	var countOne int
	result_arr := []int{}
	for i = 0; i < len(arr); i++ {
		if arr[i] != 1 {
			result_arr = append(result_arr, arr[i])
		} else {
			countOne++
		}
	}

	//mt.Println(countOne)

	for j := 0; j < countOne; j++ {

		result_arr = append(result_arr, 1)

	}
	return result_arr
}
