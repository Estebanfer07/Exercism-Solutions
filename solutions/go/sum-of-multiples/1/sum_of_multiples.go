package sumofmultiples

func SumMultiples(limit int, divisors ...int) int {
	set := make(map[int]struct{})
	sum := 0

	for _, divisor := range divisors {
		if divisor == 0 {
			continue
		}

		for val := divisor; val < limit; val += divisor {
			if _, ok := set[val]; !ok {
				set[val] = struct{}{}
				sum += val
			}
		}
	}

	return sum
}