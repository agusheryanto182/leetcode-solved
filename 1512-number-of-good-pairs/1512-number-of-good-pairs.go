func numIdenticalPairs(nums []int) int {
    counter := 0
    for i := 0; i < len(nums); i++{
        for j:=i; j < len(nums); j++{
            if nums[i] == nums[j] && i < j{
                counter++
            }
        }
    }
    return counter
}