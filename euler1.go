package euler1

func SumNaturalBelow(n int) int {

	return 0
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
