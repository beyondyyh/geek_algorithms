package offerII

// 剑指 Offer II 010. 和为 k 的子数组
// 给定一个整数数组和一个整数 k ，请找到该数组中和为 k 的连续子数组的个数。
// 示例 1 :
// 输入:nums = [1,1,1], k = 2
// 输出: 2
// 解释: 此题 [1,1] 与 [1,1] 为两种不同的情况
// 示例 2 :
// 输入:nums = [1,2,3], k = 3
// 输出: 2
// 提示:
// 1 <= nums.length <= 2 * 104
// -1000 <= nums[i] <= 1000
// -107 <= k <= 107

// 这道题目非常简洁，就是求数组中何为整数k的连续子数组个数。
// 如果这道题的取值没有负数，那就是标准的滑窗问题，但因为有了负数，滑窗思想不能用了。
// 通过分析，这道题应该属于我们上面列举四种情况的最后一种。具体思路如下：

// 初始化一个空的哈希表和pre_sum=0的前缀和变量
// 设置返回值ret = 0，用于记录满足题意的子数组数量
// 循环数组的过程中，通过原地修改数组的方式，计算数组的累加和
// 将当前累加和减去整数K的结果，在哈希表中查找是否存在
// 如果存在该key值，证明以数组某一点为起点到当前位置满足题意，ret加等于将该key值对应的value
// 判断当前的累加和是否在哈希表中，若存在value+1，若不存在value=1
// 最终返回ret即可

// class Solution {
//     public int subarraySum(int[] nums, int k) {
//         int pre_sum = 0;
//         int ret = 0;
//         HashMap<Integer, Integer> map = new HashMap<>();
//         map.put(0, 1);
//         for (int i : nums) {
//             pre_sum += i;
//             ret += map.getOrDefault(pre_sum - k, 0);
//             map.put(pre_sum, map.getOrDefault(pre_sum, 0) + 1);
//         }
//         return ret;
//     }
// }

func subarraySum(nums []int, k int) int {
	return 0
}
