package main

import "fmt"

func main() {
    fmt.Println(binaryGap(343))
}

func binaryGap(n int) int {
	maxCnt := 0
	for cnt := 0; n != 0; n = n >> 1 {
		if cnt > 0 {
			if n&1 == 0 {
				cnt++
			} else {
				if cnt > maxCnt {
					maxCnt = cnt
				}
				n = n << 1
				cnt = 0
			}
		} else {
			if n&1 == 1 {
				cnt++
			}
		}
	}
	return maxCnt
}