package summultiples

//SumMultiples returns the total sum of all multiples of one or more given divisor(s).
func SumMultiples(limit int, divisors ...int) int {
	var sum int
	for i := 1; i < limit; i++ {
		for _, d := range divisors {
			if d != 0 && i%d == 0 {
				sum += i
				break
			}
		}
	}
	return sum
}
