package euler1

func SumNaturalBelow(n int) int {
	result := 0
	for i := 1; i < n; i++ {
		if MultipleOf3Or5(i) {
			result += i
		}
	}
	return result
}

func MultipleOf3Or5(n int) bool {
	return MultipleOf3(n) || MultipleOf5(n)
}

func MultipleOf3(n int) bool {
	return (n % 3) == 0
}

func MultipleOf5(n int) bool {
	return (n % 5) == 0
}
