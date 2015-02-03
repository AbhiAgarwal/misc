package search

type KMP struct {
	table []int
}

func (k *KMP) BuildKMPTable(pattern string) {
	k.table = make([]int, len(pattern)+1)
	for i := 2; i < len(k.table); i++ {
		j := k.table[i-1]
		for {
			if pattern[j] == pattern[i-1] {
				k.table[i] = j + 1
				break
			} else if j == 0 {
				break
			} else {
				j = k.table[j]
			}
		}
	}
}

// Returns the final state when simulating the DFA built using pattern on the string text
func (k *KMP) Simulate(text, pattern string) int {
	state := 0
	for i := 0; i < len(text); i++ {
		for {
			if text[i] == pattern[state] {
				state++
				break
			} else if state == 0 {
				break
			} else {
				state = k.table[state]
			}
		}
		if state == len(k.table)-1 {
			break
		}
	}
	return state
}
