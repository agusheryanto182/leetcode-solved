function intersection(nums1: number[], nums2: number[]): number[] {
    let dict = new Map<number, number>();
    let currResult = new Map<number, boolean>();
    let result: number[] = [];

    for (let val of nums1) {
        dict.set(val, (dict.get(val) || 0) + 1);
    }

    for (let val of nums2) {
        if (dict.has(val) && !currResult.has(val)) {
            currResult.set(val, true);
            result.push(val);
        }
    }

    return result;
}
