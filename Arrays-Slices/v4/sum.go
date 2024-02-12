package main

func Sum(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}
func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, number := range numbersToSum {
		sums = append(sums, Sum(number))
	}
	return sums
}
