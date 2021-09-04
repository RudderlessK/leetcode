package medium

//给你一个字符串 s，找到 s 中最长的回文子串。
//
//
//
// 示例 1：
//
//
//输入：s = "babad"
//输出："bab"
//解释："aba" 同样是符合题意的答案。
//
//
// 示例 2：
//
//
//输入：s = "cbbd"
//输出："bb"
//
//
// 示例 3：
//
//
//输入：s = "a"
//输出："a"
//
//
// 示例 4：
//
//
//输入：s = "ac"
//输出："a"
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 1000
// s 仅由数字和英文字母（大写和/或小写）组成
//
// Related Topics 字符串 动态规划 👍 4038 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func longestPalindrome(s string) string {
	return simpleSolution(s)
}

// moreThanSimpleSolution
func moreThanSimpleSolution(s string) string{
	return ""
}

// simpleSolution 暴力拆解
func simpleSolution(s string) string {
	length := len(s)
	if length <= 1 {
		return s
	}

	var (
		maxPalindromeLen = 1
		maxPalindrome    = s[:1]
	)
	for idx, _ := range s {
		//// 要检查的最大长度小于当前回文长度值，结束判断
		//if length-idx <= maxPalindromeLen {
		//	break
		//}

		for j := idx + maxPalindromeLen; j <= length+1; j++ {
			substring := s[idx:j]
			if isPalindrome(substring) {
				maxPalindromeLen = j - idx
				maxPalindrome = s[idx:j]
			}
		}
	}
	return maxPalindrome
}

func isPalindrome(s string) bool {
	runes := []rune(s)
	length := len(runes)
	for i := 0; i < length/2; i++ {
		if runes[i] != runes[length-i-1] {
			return false
		}
	}
	return true
}
//leetcode submit region end(Prohibit modification and deletion)
