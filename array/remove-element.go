package array

import "fmt"

/** 27
* 给定一个数组 nums 和一个值 val，你需要原地移除所有数值等于 val 的元素，返回移除后数组的新长度。
* 输入: nums = [3,2,2,3], val = 3,
* 输出: 函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。
*
* 说明
* 不能使用额外的数组空间
**/

func RemoveElement(nums []int, val int) int {
	var i int
	for _,v := range nums  {
		if v != val {
			nums[i] = v
			i++
		}
	}
	nums = nums[:i]
	fmt.Println(nums)
	return i
}

/** 26
* 给定一个排序数组，你需要在 原地 删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。
* 输入: nums = [0,0,1,1,1,2,2,3,3,4]
* 输出: 函数应该返回新的长度 5, 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4。
*
* 说明
* 不能使用额外的数组空间
* 善用双指针
**/

func RemoveDuplicates(nums []int) int {
	i,j := 0,0
	for j = 1 ;j< len(nums);j++ {
		if nums[i] != nums[j]{
			i++
			nums[i] = nums[j]
		}

		fmt.Println(nums)
	}

	nums = nums[:i+1]
	fmt.Println(nums)
	return i+1
}