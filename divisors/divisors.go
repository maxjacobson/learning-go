package divisors

type Divisor struct {
	Num int
}

func (divisor *Divisor) MoreThan(count int) bool {
	actualCount := 0
	for i := 1; i <= divisor.Num; i++ {
		if divisor.Num%i == 0 {
			actualCount += 1
		}
	}
	return actualCount > count
}

func New(num int) Divisor {
	return Divisor{Num: num}
}
