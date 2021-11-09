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
