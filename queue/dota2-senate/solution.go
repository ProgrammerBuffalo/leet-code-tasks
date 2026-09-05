package main

import "fmt"

func main() {
	fmt.Println(predictPartyVictory("DDRRR"))
}

func predictPartyVictory(senate string) string {
	q := make([]byte, 0, len(senate))
	qNext := make([]byte, 0, len(senate))
	for i := 0; i < len(senate); i++ {
		q = append(q, senate[i])
	}

	for d, r := 0, 0; ; d, r = 0, 0 {
		for i := 0; i < len(q); i++ {
			if q[i] == 'D' {
				j := i
				for ; j < len(q) && q[j] != 'R'; j++ {
				}
				if j < len(q) {
					q[j] = 'X'
				} else {
					for k := 0; k < len(qNext); k++ {
						if qNext[k] == 'R' {
							qNext[k] = 'X'
							r--
							break
						}
					}
				}
				d++
				qNext = append(qNext, q[i])
			} else if q[i] == 'R' {
				j := i
				for ; j < len(q) && q[j] != 'D'; j++ {
				}
				if j < len(q) {
					q[j] = 'X'
				} else {
					for k := 0; k < len(qNext); k++ {
						if qNext[k] == 'D' {
							qNext[k] = 'X'
							d--
							break
						}
					}
				}
				r++
				qNext = append(qNext, q[i])
			}
		}
		if d <= 0 {
			return "Radiant"
		} else if r <= 0 {
			return "Dire"
		}
		q = qNext
		qNext = qNext[:0]
	}

}
