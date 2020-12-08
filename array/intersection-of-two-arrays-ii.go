package array

/** 350
* 给定两个数组，编写一个函数来计算它们的交集。
* 输入: nums1 = [1,2,2,1], nums2 = [2,2]
* 输出: [2,2]
*
* 分析
* 典型映射题 <元素,出现次数>
*
* 进阶
* 给定的数组已排序
* 双指针
**/

func Intersect(sliceNum1 []int, sliceNum2 []int) []int {
	tmpMap := map[int]int{}
	for _, v := range sliceNum1 {
		tmpMap[v] ++
	}
	k := 0
	for _, v := range sliceNum2 {
		if tmpMap[v] > 0 {
			tmpMap[v] --
			sliceNum2[k] = v
			k++
		}
	}
	return sliceNum2[0:k]
}

func IntersectPlus(sliceNum1 []int, sliceNum2 []int) []int {
	i, j,k := 0, 0,0
	for i < len(sliceNum1) && j < len(sliceNum2) {
		if sliceNum1[i] == sliceNum2[j] {
			sliceNum1[k] = sliceNum1[i]
			i++
			j++
			k++
		} else if sliceNum1[i] > sliceNum2[j] {
			j++
		} else {
			i++
		}
	}
	return sliceNum1[0:k]
}
