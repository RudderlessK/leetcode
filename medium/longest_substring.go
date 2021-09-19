package medium

//给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
//
//
//
// 示例 1:
//
//
//输入: s = "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//
//
// 示例 2:
//
//
//输入: s = "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//
//
// 示例 3:
//
//
//输入: s = "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
//
//
// 示例 4:
//
//
//输入: s = ""
//输出: 0
//
//
//
//
// 提示：
//
//
// 0 <= s.length <= 5 * 10⁴
// s 由英文字母、数字、符号和空格组成
//
// Related Topics 哈希表 字符串 滑动窗口 👍 6090 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubstring(s string) int {
	return SlideWindowSolution(s)
}

func SlideWindowSolution(s string) int {
	var (
		left, right   int
		longestLength int

		window = make(map[rune]bool)
	)

	runes := []rune(s)
	for right < len(runes) {
		rnew := runes[right]

		// 不存在right滑动
		if _, ok := window[rnew]; !ok {
			window[rnew] = true
			right++
			continue
		}
		longestLength = max(longestLength, right-left)

		// 已经存在left滑动
		for _, r := range runes[left:right] {
			left++
			delete(window, runes[left])
			if r == rnew {
				break
			}
		}
	}

	return longestLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
