package easy

//给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target 的那 两个 整数，并返回它们的数组下标。
//
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
//
// 你可以按任意顺序返回答案。
//
//
//
// 示例 1：
//
//
//输入：nums = [2,7,11,15], target = 9
//输出：[0,1]
//解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
//
//
// 示例 2：
//
//
//输入：nums = [3,2,4], target = 6
//输出：[1,2]
//
//
// 示例 3：
//
//
//输入：nums = [3,3], target = 6
//输出：[0,1]
//
//
//
//
// 提示：
//
//
// 2 <= nums.length <= 104
// -109 <= nums[i] <= 109
// -109 <= target <= 109
// 只会存在一个有效答案
//
//
// 进阶：你可以想出一个时间复杂度小于 O(n2) 的算法吗？
// Related Topics 数组 哈希表
// 👍 11555 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func twoSum(num []int, target int) []int {
	return twoSumHashBetterSolution(num, target)
}

func twoSumDirectSolution(nums []int, target int) []int {
	// 两个指针遍历寻找
	for idx1, currentPointer := range nums {
		startIdx := idx1 + 1
		for idx2, nextPointer := range nums[startIdx:] {
			if (currentPointer + nextPointer) == target {
				return []int{idx1, startIdx + idx2}
			}
		}
	}
	return nil
}

func twoSumHashSolution(nums []int, target int) []int {
	hashTable := make(map[int]int)
	for idx, v := range nums {
		hashTable[v] = idx
	}

	for idx, v := range nums {
		if nextItem, ok := hashTable[target-v]; ok {
			if idx == hashTable[nextItem] {
				continue
			}
			return []int{idx, hashTable[nextItem]}
		}
	}

	return nil
}

func twoSumHashBetterSolution(nums []int, target int) []int {
	hashTable := make(map[int]int)
	for idx, v := range nums {
		if secondIdx, ok := hashTable[target-v]; ok {
			return []int{idx, secondIdx}
		}
		hashTable[v] = idx
	}
	return nil
}

//leetcode submit region end(Prohibit modification and deletion)
