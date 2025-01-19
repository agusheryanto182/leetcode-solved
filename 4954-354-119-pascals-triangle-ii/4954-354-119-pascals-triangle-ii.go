func getRow(rowIndex int) []int {
    resTemp  := [][]int{{1}}

    if (rowIndex == 0) {
        return  []int{1}
    } else {
        resTemp = append(resTemp, []int{1,1})
    }

    if (rowIndex == 1) {
        return []int{1,1}
    }

    for len(resTemp) <= rowIndex {
        currLength := len(resTemp)
        prevRow := resTemp[currLength -1]
        res := []int{1}

        for i:=0; i < len(prevRow) -1; i++ {
            res = append(res, calculate(prevRow[i], prevRow[i+1]))
        }

        res = append(res, 1)

        resTemp = append(resTemp, res)
    }
    return resTemp[rowIndex]
}

func calculate(n1, n2 int) int {
    return n1 + n2
}