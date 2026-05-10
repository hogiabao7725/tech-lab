package main

func twoSum(nums []int, target int) []int {
	myMap := make(map[int]int, len(nums)) // value / index

	for i, v := range nums {
		need := target - v

		if idx, ok := myMap[need]; ok {
			return []int{idx, i}
		} else {
			myMap[v] = i
		}
	}

	return []int{}
}
