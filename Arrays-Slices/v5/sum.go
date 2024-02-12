package main

func Sum(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, number := range numbersToSum {
		tail := number[1:]
		sums = append(sums, Sum(tail))

	}
	return sums
}
