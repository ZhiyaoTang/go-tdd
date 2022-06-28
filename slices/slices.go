package slices

func Sum(numbers []int) int {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}

	return sum
}

func SumAll(numbersLists ...[]int) []int {
	var sums []int
	for _, numbers := range numbersLists {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails()
