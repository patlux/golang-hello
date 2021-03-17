package main

func Sum(arr [5]int) (sum int) {

	for i := 0; i<len(arr); i++ {
		sum += arr[i]
	}

	return
}
