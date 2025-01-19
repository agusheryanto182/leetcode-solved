func generate(numRows int) [][]int {
    result := [][]int{{1}}
    
    if numRows == 1 {
        return result
    } 
    
    result = append(result, []int{1,1})
    
    for len(result) < numRows {
        currLength := len(result)
        prevRow := result[currLength -1]
        res := []int{1}
        
        for i := 0; i < len(prevRow) -1; i++{
            res = append(res, calculate(prevRow[i], prevRow[i+1]))
        }
        
        res = append(res, 1)
        
        result = append(result, res)
    }
    
    return result
}

func calculate(n1, n2 int) int {
    return n1 + n2
}