package trinums

func At(position int) int {
	sum := 0
	for i := 1; i <= position; i++ {
		sum += i
	}
	return sum
}
