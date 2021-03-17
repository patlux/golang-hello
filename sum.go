package main

func Sum(arr []int) (sum int) {

	for _, nr := range arr {
		sum += nr
	}

	return
}

func SumAll(arr, arr2 []int) (sumarr []int) {
	all := [][]int{arr, arr2}

	for _, arr := range all {
		result := Sum(arr)
		sumarr = append(sumarr, result)
	}
	
	return
}
