package arrays

func SumFixedSize(numbers [5]int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}

func Sum(numbers []int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}

func SumAll(numberToSum ...[]int) []int {
	var result []int
	for i, numbers := range numberToSum {
		result[i] = Sum(numbers)
	}
	return result
}

func SumAllTails(numberToSum ...[]int) (sum []int) {
	for _, numbers := range numberToSum {
		tail := numbers[1:]
		sum = append(sum, Sum(tail))
	}
	return sum
}
