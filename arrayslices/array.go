package main

// Sum does...
func Sum(numbers []int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return sum
}

// SumAll does...
func SumAll(numbersToSum ...[]int) (sums []int) {
	lenInput := len(numbersToSum)
	sums = make([]int, lenInput)

	for i, v := range numbersToSum {
		sums[i] = Sum(v)
	}

	return
}

// SumAllWithAppend makes use of the append()
func SumAllWithAppend(numbersToSum ...[]int) (sums []int) {
	for _, v := range numbersToSum {
		sums = append(sums, Sum(v))
	}

	return
}

// SumTails neglects the first array value and sums only the remaining ones
func SumTails(numbersToSum ...[]int) (sums []int) {
	for _, v := range numbersToSum {
		if len(v) == 0 {
			sums = append(sums, 0)
		} else {
			tail := v[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return
}
