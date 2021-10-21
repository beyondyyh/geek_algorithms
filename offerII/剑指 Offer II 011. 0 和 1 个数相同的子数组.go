package offerII

// 剑指 Offer II 011. 0 和 1 个数相同的子数组
// 给定一个二进制数组 nums , 找到含有相同数量的 0 和 1 的最长连续子数组，并返回该子数组的长度。
// 示例 1:
// 输入: nums = [0,1]
// 输出: 2
// 说明: [0, 1] 是具有相同数量 0 和 1 的最长连续子数组。
// 示例 2:
// 输入: nums = [0,1,0]
// 输出: 2
// 说明: [0, 1] (或 [1, 0]) 是具有相同数量 0 和 1 的最长连续子数组。

func findMaxLength(nums []int) int {
	// var mp := make(map[int]int, len(nums))

	return 0
}

// class Solution {
// public:
// 	int findMaxLength(vector<int>& nums) {
// 		unordered_map<int, int> mp;
// 		int preSum = 0, maxLen = 0;
// 		mp[0] = -1;
// 		for (int i = 0; i < nums.size(); ++i){
// 			preSum += nums[i] ? 1 : -1;
// 			if(mp.find(preSum) != mp.end()) maxLen = max(maxLen, i - mp[preSum]);
// 			else mp[preSum] = i;
// 		}
// 		return maxLen;
// 	}
// };
