/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {number[]} nums
 * @return {TreeNode}
 */
var sortedArrayToBST = function (nums) {
    if (nums.length === 0) {
        return null;
    }

    const buildBST = (left, right) => {
        if (left > right) {
            return null;
        }

        const mid = Math.floor((left + right) / 2);
        const node = new TreeNode(nums[mid]);
        node.left = buildBST(left, mid - 1);
        node.right = buildBST(mid + 1, right);

        return node;
    };

    return buildBST(0, nums.length - 1);
};