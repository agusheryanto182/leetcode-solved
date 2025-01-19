function containsNearbyDuplicate(nums: number[], k: number): boolean {
    let temp = new Map<number, number[]>();
    let total : number = 0;
    
    if (nums.length == 1) {
        return false
    }

    for  (let i : number = 0; i < nums.length; i ++) {
        if (temp.has(nums[i])) {
            let arr = temp.get(nums[i])!;
            arr.push(i)
            total += 1
        } else {
            temp.set(nums[i], [i])
        }
    }


    for (let i : number = 0; i < nums.length; i++) {
        let currTemp =  temp.get(nums[i])
        if (currTemp.length > 1) {
            for (let j : number = 0; j < currTemp.length - 1; j++) {
             let result = Math.abs(currTemp[j] - currTemp[j+1])   
             if (result <= k) {
                    total -= 1
                    break
                }
            }
            if (total == 0) {
            return true
            }
        }
    }

    return false
};