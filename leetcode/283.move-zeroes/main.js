var moveZeroos = function (nums) {
    let pos = 0;
    for (let i = 0; i < nums.length; i++) {
        if (nums[i] != 0) {
            if (pos != i) {
                tmp = nums[i];
                nums[i] = nums[pos];
                nums[pos] = tmp;
            }
            pos++;
        }
    }
};

nums = [1, 2, 3, 0, 5];
moveZeroos(nums);
console.log(nums);
