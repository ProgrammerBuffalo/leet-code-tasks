package main

import "fmt"

type RecentCounter struct {
	queue []int
}

func Constructor() RecentCounter {
	return RecentCounter{queue: make([]int, 0)}
}

func (this *RecentCounter) Ping(t int) int {
	diff := t - 3000
	for len(this.queue) > 0 {
		if this.queue[0] >= diff {
			break
		}
		this.queue = this.queue[1:]
	}

	this.queue = append(this.queue, t)

	return len(this.queue)
}

func main() {
	recentCounter := Constructor()
	fmt.Println(recentCounter.Ping(1))    // requests = [1], range is [-2999,1], return 1
	fmt.Println(recentCounter.Ping(100))  // requests = [1, 100], range is [-2900,100], return 2
	fmt.Println(recentCounter.Ping(3001)) // requests = [1, 100, 3001], range is [1,3001], return 3
	fmt.Println(recentCounter.Ping(3002)) // requests = [1, 100, 3001, 3002], range is [2,3002], return 3
}
