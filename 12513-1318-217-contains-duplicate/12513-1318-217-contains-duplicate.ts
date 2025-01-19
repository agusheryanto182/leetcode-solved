function containsDuplicate(nums: number[]): boolean {
    let temp = new Map<number, number>();
    
    for (let i : number = 0; i < nums.length; i++) {
        let currentTotal = temp.get(nums[i]) || 0
        temp.set(nums[i], currentTotal + 1)
        
        if (temp.get(nums[i]) > 1) {
            return true
        }   
    }
    
    return false
};