package easy

import (
	"bytes"
	"encoding/gob"
	"sort"
)

//编写一个函数来查找字符串数组中的最长公共前缀。
//
// 如果不存在公共前缀，返回空字符串 ""。
//
//
//
// 示例 1：
//
//
//输入：strs = ["flower","flow","flight"]
//输出："fl"
//
//
// 示例 2：
//
//
//输入：strs = ["dog","racecar","car"]
//输出：""
//解释：输入不存在公共前缀。
//
//
//
// 提示：
//
//
// 1 <= strs.length <= 200
// 0 <= strs[i].length <= 200
// strs[i] 仅由小写英文字母组成
//
// Related Topics 字符串 👍 1765 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func longestCommonPrefix(strs []string) string {
	length := len(strs)
	if length == 0 {
		return ""
	}

	//copyStrs := make([]string, length)
	//copyStrs = append(copyStrs, strs...)
	//sort.Strings(copyStrs)
	//min, max := copyStrs[0], copyStrs[length-1]

	sort.Strings(strs)
	min, max := []rune(strs[0]), []rune(strs[length-1])

	var prefixs []rune
	for idx, c := range min {
		if c == max[idx] {
			prefixs = append(prefixs, c)
			continue
		}
	}

	return string(prefixs)
}

func deepCopy(a interface{}, b interface{}) error {
	buff := new(bytes.Buffer)
	enc := gob.NewEncoder(buff)
	dec := gob.NewDecoder(buff)

	if err := enc.Encode(a); err != nil {
		return err
	}

	return dec.Decode(b)
}

//leetcode submit region end(Prohibit modification and deletion)
