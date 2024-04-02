package _1_TwoSum

func twoSum(nums []int, target int) []int {
	hash := map[int]int{}
	for i, x := range nums {
		if _, ok := hash[target-x]; ok {
			return []int{i, hash[target-x]}
		}
		hash[x] = i
	}
	return []int{}
}
