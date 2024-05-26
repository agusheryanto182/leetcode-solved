// func shuffle(nums []int, n int) []int {
// 	result := []int{}
// 	numsx := []int{}
// 	numsy := []int{}
// 	numsx = append(numsx, nums[:len(nums)/2]...)
// 	numsy = append(numsy, nums[len(nums)/2:]...)

// 	for i := 0; i < len(nums)/2; i++ {
// 		result = append(result, numsx[i])
// 		result = append(result, numsy[i])
// 	}
// 	return result
// }

func shuffle(nums []int, n int) []int {
    result := make([]int,2*n)
    for i:=0;i<n;i++{
        result[i*2] = nums[i]
        result[i*2+1] = nums[i+n]
    }
    return result
}