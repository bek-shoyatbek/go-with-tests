package arraysandslices

func SumNumbers(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}
