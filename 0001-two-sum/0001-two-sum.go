func twoSum(nums []int, target int) []int {
    var dict = make(map[int]int)
    
    for indexCurr, valCurr := range nums{
        indexDict, isPresent := dict[target - valCurr]

        if isPresent {
            return []int{indexDict, indexCurr}
        }
        dict[valCurr] = indexCurr
    }
    return []int{}
}