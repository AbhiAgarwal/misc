// generate and filter from: https://golang.org/doc/play/sieve.go
package math

// Send the sequence 2, 3, 4, ... to channel 'ch'.
func generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in // Receive value from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
	}
}

func GetPrimeNumbers(value int) []int {
	var primeNumbers []int
	ch := make(chan int)
	go generate(ch)
	for i := 0; i < value; i++ {
		prime := <-ch
		primeNumbers = append(primeNumbers, prime)
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
		if prime > value {
			break
		}
	}
	return primeNumbers
}

// Method to shorten time in searching for IsPrime function.
func getPrimeNumberHash(value int) map[int]bool {
	var primeNumbers map[int]bool
	primeNumbers = make(map[int]bool)
	ch := make(chan int)
	go generate(ch)
	for {
		prime := <-ch
		primeNumbers[prime] = true
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
		if prime > value {
			break
		}
	}
	return primeNumbers
}

// The prime sieve: Daisy-chain Filter processes.
func IsPrime(value int) bool {
	var primeNumbers map[int]bool
	primeNumbers = getPrimeNumberHash(value)
	if primeNumbers[value] {
		return true
	}
	return false
}
