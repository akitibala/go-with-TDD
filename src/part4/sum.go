package part4

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}
	return sum
}

// func SumAll(numbersToSum ...[]int) []int {
// 	lenOfNumbers := len(numbersToSum)
// 	sums := make([]int, lenOfNumbers)
// 	for i, num := range numbersToSum {
// 		sums[i] = Sum(num)
// 	}
// 	return sums
// }

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		tail := numbers[1:]
		sums = append(sums, Sum(tail))
	}

	return sums
}
