package array

import (
	"strings"
)

/** 14
* 编写一个函数来查找字符串数组中的最长公共前缀。如果不存在公共前缀，则返回""。
* 输入: ["flower","flow","flight"]
* 输出: "fl"
*
* 分析
* 拆分元素依次对比
* strings.Index(s,str) -1:查找不到 0:str为空 >0:str第一次出现的位置
**/
func LongestCommonPrefix(str []string) string {
	if len(str) <= 0 {
		return ""
	}
	first := str[0]
	for _, v := range str {
		for strings.Index(v, first) != 0 {
			if len(first) == 0 {
				return ""
			}
			first = first[:len(first)-1]
		}
	}
	return first
}
