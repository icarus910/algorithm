package array

/** 189
* 给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
* 输入: [-1,-100,3,99] 和 k = 2
* 输出: [3,99,-1,-100]
*
* 知识点
* append会生成新的数组 所以使用copy
**/

func Rotate(nums []int, k int)  {
	length := len(nums)
	k = k %length
	if k <= length  {
		copy(nums, append(nums[length-k:], nums[0:length-k]...))
	}
}