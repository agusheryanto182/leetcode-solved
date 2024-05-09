// 3, 6, 7, 8, 10
// left : 0
// right : 4

// mid : 2
// right = 1
// mid : 0
// left = 1
// mid : 1 
// right = 0
// break
// return
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}