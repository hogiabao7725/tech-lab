package main

func hasDuplicate(nums []int) bool {
	myMap := make(map[int]struct{})

	for _, v := range nums {
		if _, ok := myMap[v]; ok {
			return true
		}
		myMap[v] = struct{}{}
	}

	return false
}
