package main

// my solution
/*
	func getConcatenation(nums []int) []int {
		length := len(nums)
		arr := make([]int, 2 * length)

		for i := 0; i < length; i++ {
			arr[i] = nums[i]
			arr[i + length] = nums[i]
		}

		return arr
	}
*/

// optimal solution
func getConcatenation(nums []int) []int {
	ans := make([]int, 0, 2*len(nums))
	ans = append(ans, nums...)
	ans = append(ans, nums...)

	return ans
}

// some bad examples
/*
	func main() {
		original := []int{1, 2, 3}

		// Assume original has a capacity greater than its length,
		// and we append to it, it may modify the underlying array,
		// which can lead to unexpected results.
		// Ex: original = make([]int, 3) -> cap=10, len=3

		result := getConcatenation(original)

		fmt.Println(result)   // [1, 2, 3, 1, 2, 3]
		fmt.Println(original) // [1, 2, 3, 1, 2, 3] <- BUG! Data is modified due to shared underlying array
	}
*/
