function findDisappearedNumbers(nums: number[]): number[] {
    let result: number[] = [];
    let numSet = new Set(nums);
    
    for (let i = 1; i <= nums.length; i++) {
        if (!numSet.has(i)) {
            result.push(i);
        }
    }
    
    return result;
}
