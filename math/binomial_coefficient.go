package math

func BinomialCoefficient(n, k uint64) uint64 {
	var val uint64 = 1
	// Multiplicative formula - individual binomial coefficients
	for i := uint64(1); i <= k; i++ {
		val *= (((n + 1) - i) / i)
	}
	return val
}
