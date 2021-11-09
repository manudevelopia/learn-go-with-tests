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

func SumAll(numberToSum ...[]int) (sum []int) {
	result := make([]int, len(numberToSum))
	for i, numbers := range numberToSum {
		result[i] = Sum(numbers)
	}
	return result
}
