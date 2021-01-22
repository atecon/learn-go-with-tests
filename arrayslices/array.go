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
