func twoSum(nums []int, target int) []int {
	set := make(map[int]int)

	for i, v := range nums {
        val := target - v
        if _, ok := set[val]; !ok {
            set[v] = i
        } else {
            return [] int {i, set[val]}
        }
    }

	return []int{}
}
