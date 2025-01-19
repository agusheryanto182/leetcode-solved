function generate(numRows: number): number[][] {
    let result: number[][] = [
        [1],
    ]

    if (numRows  == 1) {
        return result
    } else {
        result.push([1,1])
    }

    if (numRows === 2) {
        return result;
    }

    while (true) {
        let currLength : number = result.length;
        let prevRow : number[] = result[currLength - 1];
        let res : number[] = [];

        for (let i : number = 0; i < prevRow.length - 1; i++) {
            let resCalculate = calculate(prevRow[i], prevRow[i + 1])
            res.push(resCalculate)
        }
        result.push([1, ...res ,1])

        if (result.length == numRows) break;
    }
    return result
};

function calculate(numb1: number, numb2: number): number {
    return numb1 + numb2
}