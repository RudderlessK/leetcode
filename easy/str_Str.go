package easy

import (
	"reflect"
	"strings"
)

//实现 strStr() 函数。
//
// 给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。如
//果不存在，则返回 -1 。
//
//
//
// 说明：
//
// 当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。
//
// 对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与 C 语言的 strstr() 以及 Java 的 indexOf() 定义相符。
//
//
//
// 示例 1：
//
//
//输入：haystack = "hello", needle = "ll"
//输出：2
//
//
// 示例 2：
//
//
//输入：haystack = "aaaaa", needle = "bba"
//输出：-1
//
//
// 示例 3：
//
//
//输入：haystack = "", needle = ""
//输出：0
//
//
//
//
// 提示：
//
//
// 0 <= haystack.length, needle.length <= 5 * 104
// haystack 和 needle 仅由小写英文字符组成
//
// Related Topics 双指针 字符串 字符串匹配
// 👍 959 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func strStr(haystack string, needle string) int {
	return SundaySolution2(haystack, needle)
}

func strStrSolution1(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	return strings.Index(haystack, needle)
}

// 字符串很大且重合度高时，暴力解法效率非常低，会超时
// 特别是用了reflect.DeepEqual
func strStrSolution2(haystack string, needle string) int {
	if needle == "" || haystack == needle {
		return 0
	}

	var (
		runesHaystack  = []rune(haystack)
		runesNeedle    = []rune(needle)
		lengthHaystack = len(runesHaystack)
		lengthNeedle   = len(runesNeedle)
	)

	if lengthHaystack < lengthNeedle {
		return -1
	}

	for i := 0; i <= lengthHaystack-lengthNeedle; i++ {
		if runesHaystack[i] == runesNeedle[0] &&
			reflect.DeepEqual(runesNeedle[1:], runesHaystack[i+1:i+lengthNeedle]) {
			return i
		}
	}
	return -1
}

// Sunday算法
func SundaySolution2(haystack string, needle string) int {
	if needle == "" || haystack == needle {
		return 0
	}

	var (
		lengthNeedle   = len(needle)
		lengthHaystack = len(haystack)

		offset int
	)

	if lengthHaystack < lengthNeedle {
		return -1
	}

	for i := 0; i <= lengthHaystack-lengthNeedle; {
		if needle == haystack[i:i+lengthNeedle] {
			return i
		}
		if i+lengthNeedle == lengthHaystack{
			break
		}

		// 寻找偏移量
		offset = i + lengthNeedle + 1
		for j := 0; j < lengthNeedle; j++ {
			if haystack[i+lengthNeedle] == needle[j] {
				offset = lengthNeedle - j
				break
			}
		}

		i += offset

	}

	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
