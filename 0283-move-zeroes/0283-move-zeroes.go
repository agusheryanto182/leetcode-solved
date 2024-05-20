func moveZeroes(nums []int) {
    lastNonZeroFoundAt := 0
    
    // Move all non-zero elements to the front of the array
    for i := 0; i < len(nums); i++ {
        if nums[i] != 0 {
            nums[lastNonZeroFoundAt] = nums[i]
            lastNonZeroFoundAt++
        }
    }
    
    // Fill the remaining positions with zeroes
    for i := lastNonZeroFoundAt; i < len(nums); i++ {
        nums[i] = 0
    }
}