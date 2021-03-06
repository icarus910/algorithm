package array

/** 66
* 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
* 输入: [4,3,2,1]
* 输出: [4,3,2,2]
*
* 分析
* 9加一需进位
**/
func PlusOne(digits []int) []int {
	length := len(digits)
	if digits[length-1] != 9 {
		digits[length-1] ++
	} else {
		for length != 0 && digits[length-1] == 9 {
			digits[length-1] = 0
			length--
		}
		if length-1 >= 0 {
			digits[length-1] ++
		}

		if digits[0] == 10 || digits[0] == 0 {
			digits[0] = 0
			digits = append([]int{1}, digits...)
		}
	}
	return digits
}
