import "math"
func thirdMax(nums []int) int {
    max, second, third := math.MinInt64, math.MinInt64, math.MinInt64

    for _, v := range(nums){
        if v == max || v == second || v == third{
                continue
        }
        switch{
        case v >  max:
            max, second, third = v, max, second
            break
        case v > second:
            second, third = v, second
            break
        case v > third:
            third = v        
        }

    }

    if third == math.MinInt64 {
		return max
	}else{
        return third
    }


}