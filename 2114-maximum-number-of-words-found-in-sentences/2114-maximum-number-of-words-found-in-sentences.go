func mostWordsFound(sentences []string) int {
    max := 0
    for _, val := range sentences{
        res := strings.Split(val, " ")
        if len(res) > max {
            max = len(res)
        }
    }
    return max
}