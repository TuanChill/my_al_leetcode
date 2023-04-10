/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function (nums, target) {
  for (let i = 0; i <= nums.length - 2; i++) {
    const otheri = nums.indexOf(target - nums[i], i + 1);
    if (otheri !== -1) return [i, otheri];
  }
};

//case 1:
console.log(twoSum([2, 7, 11, 15], 9));

//case 2:
console.log(twoSum([3, 2, 4], 6));

//case 3:
console.log(twoSum([3, 3], 6));
