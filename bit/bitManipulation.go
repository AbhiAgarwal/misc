package bit

import "math"

func IsOn(S int, j uint64) bool {
	val := S & (1 << j)
	return (val > 0)
}

func SetBit(S int, j uint64) int {
	S |= (1 << j)
	return S
}

func ClearBit(S int, j uint64) int {
	S &= ^(1 << j)
	return S
}

func ToggleBit(S int, j uint64) int {
	S ^= (1 << j)
	return S
}

func LowBit(S int) int {
	return S & (-S)
}

func SetAll(S int, j uint64) int {
	S = (1 << j) - 1
	return S
}

func Modulo(S int, j int) int {
	return (S) & (j - 1)
}

func IsPowerOfTwo(S int) int {
	return ^(S & (S - 1))
}

func NearestPowerOfTwo(S int) float64 {
	return math.Pow(2.0, float64(((math.Log(float64(S)) / math.Log(2.0)) + 0.5)))
}

func TurnOffLastBit(S int) int {
	return (S) & (S - 1)
}

func TurnOnLastZero(S int) int {
	return (S) | (S + 1)
}

func TurnOffLastConsecutiveBits(S int) int {
	return (S) & (S + 1)
}

func TurnOnLastConsecutiveZeroes(S int) int {
	return (S) | (S - 1)
}
