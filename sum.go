package main

func Sum(arr []int) (sum int) {
	for _, nr := range arr {
		sum += nr
	}
	return
}

func SumAll(listOfArrNumbers ...[]int) (sumarr []int) {
	for _, arrNumbers := range listOfArrNumbers {
		sumarr = append(sumarr, Sum(arrNumbers))
	}
	return
}

func SumAllTails(listOfArrNumbers ...[]int) (sumarr []int) {
	for _, arrNumbers := range listOfArrNumbers {
		if len(arrNumbers) == 0 {
			sumarr = append(sumarr, 0)
			continue
		}

		lastNumbers := arrNumbers[1:]
		sumarr = append(sumarr, Sum(lastNumbers))
	}
	return
}
