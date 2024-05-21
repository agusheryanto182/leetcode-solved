// func countPairs(nums []int, target int) int {
// 	result := 0

// 	for i := 0; i < len(nums)-1; i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			if nums[i]+nums[j] < target {
// 				result++
// 			}
// 		}
// 	}
// 	return result
// }

func countPairs(nums []int, target int) int {
    count := 0
    sort.Ints(nums)

    left, right := 0, len(nums) - 1

    for left < right {
        if nums[left] + nums[right] < target {
            count += right - left
            left ++
        } else {
            right --
        }
    }

    return count
}