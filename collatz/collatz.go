package collatz

type Collatz struct {
	StartingPoint int
	Length        int
}

func From(from int) Collatz {
	collatz := Collatz{
		StartingPoint: from,
		Length:        1,
	}

	return calculate(collatz)
}

func calculate(collatz Collatz) Collatz {
	if collatz.StartingPoint == 1 {
		// We're done!
		return collatz
	}

	if collatz.StartingPoint%2 == 0 {
		collatz.StartingPoint = collatz.StartingPoint / 2
	} else {
		collatz.StartingPoint = collatz.StartingPoint*3 + 1
	}

	collatz.Length = collatz.Length + 1

	return calculate(collatz)
}
