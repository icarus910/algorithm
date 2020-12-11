package array

import (
	"fmt"
	"sort"
	"strconv"
)

/** 15
* 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a+b+c=0 ？请你找出所有满足条件且不重复的三元组。注意：答案中不可以包含重复的三元组。
*
* 输入:  nums = [-1, 0, 1, 2, -1, -4]
* 输出:  [
*			[-1, 0, 1],
*			[-1, -1, 2]
* 		 ]
*
* 分析
* 双指针，先排序，三数之和 <0 则值小的指针向后移动
**/

func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	var threeInt [][]int
	length := len(nums)
	for k, v := range nums {
		i, j := k+1, length-1
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}
		for i < j {
			fmt.Println("i:" + strconv.Itoa(i))
			fmt.Println("j:" + strconv.Itoa(j))
			fmt.Println("------------")
			result := v + nums[i] + nums[j]
			if result == 0 {
				threeInt = append(threeInt, []int{v, nums[i], nums[j]})
				for i < j && nums[i] == nums[i+1] {
					i += 1
				}
				for i < j && nums[j] == nums[j-1] {
					j -= 1
				}
				i += 1
				j -= 1
			} else if result < 0 {
				i++
			} else {
				j--
			}
		}

	}
	return threeInt
	//// 先从小到大排序
	//sort.Ints(nums)
	//// 接收结果
	//var res [][]int
	//// 获取数组长度
	//length := len(nums)
	//// 边界处理，数字不足三个直接返回空
	//if len(nums) < 3 {
	//	return res
	//}
	//// 开始循环第一个固定值
	//for index, value := range nums {
	//	// 如果固定位的值已经大于0，因为已经排好序了，后面的两个指针对应的值也肯定大于0，则和不可能为0，所以返回
	//	if nums[index] > 0 {
	//		return res
	//	}
	//	// 排除值重复的固定位
	//	if index > 0 && nums[index] == nums[index-1] {
	//		continue
	//	}
	//	// 指针初始位置，固定位右边第一个和数组最后一个
	//	l := index + 1
	//	r := length - 1
	//	// 开始移动两个指针
	//	for l < r {
	//		// 判断三个数字之和的三种情况
	//		sum := value + nums[l] + nums[r]
	//		switch {
	//		case sum == 0:
	//			// 将结果加入二元组
	//			res = append(res, []int{nums[index], nums[l], nums[r]})
	//			// 去重，如果l < r且下一个数字一样，则继续挪动
	//			for l < r && nums[l] == nums[l+1] {
	//				l += 1
	//			}
	//			// 同理
	//			for l < r && nums[r] == nums[r-1] {
	//				r -= 1
	//			}
	//			l += 1
	//			r -= 1
	//		case sum > 0:
	//			// 如果和大于 0，那就说明 right 的值太大，需要左移
	//			r -= 1
	//			// 如果和小于 0，那就说明 left 的值太小，需要右移
	//		case sum < 0:
	//			l += 1
	//		}
	//	}
	//}
	//return res
}

//三指针 error 思路错误
func ThreePointer(nums []int) [][]int {
	var threeInt [][]int
	length := len(nums)
	if length < 3 {
		return threeInt
	}
	x, y, z := 0, 1, 2
	a, b, c := 0, 1, 2
	for !(x == length-3 && y == length-2 && z == length-1) {
		if x == length-3 && y == length-2 && z == length-1 {
			break
		}
		result := nums[x] + nums[y] + nums[z]

		if result == 0 {
			tmpInt := []int{nums[x], nums[y], nums[z]}
			sort.Ints(tmpInt)

			if unique(tmpInt,threeInt) {
				threeInt = append(threeInt, tmpInt)
			}

		}
		if z < length-1 {
			z++
			if y == length-2 && z == length-1 {
				a++
				x = a
				b = x + 1
				y = b
				c = y + 1
				z = c
			}
			continue
		} else {
			c++
			z = c
		}
		if y != length-2 {
			y++
			if y == length-2 && z == length-1 {
				a++
				x = a
				b = x + 1
				y = b
				c = y + 1
				z = c
			}
			continue
		} else {
			b++
			y = b
		}
	}
	if nums[length-3]+nums[length-2]+nums[length-1] == 0 && unique([]int{nums[length-3], nums[length-2], nums[length-1]},threeInt) {
		threeInt = append(threeInt, []int{nums[length-3], nums[length-2], nums[length-1]})
	}
	return threeInt
}

func unique(tmpInt []int, threeInt [][]int) bool {
	check := true
	for _, v := range threeInt {
		fmt.Println("v[0]:" + strconv.Itoa(v[0]))
		fmt.Println("v[1]:" + strconv.Itoa(v[1]))
		fmt.Println("v[2]:" + strconv.Itoa(v[2]))
		fmt.Println("tmpInt[0]:" + strconv.Itoa(tmpInt[0]))
		fmt.Println("tmpInt[1]:" + strconv.Itoa(tmpInt[1]))
		fmt.Println("tmpInt[2]:" + strconv.Itoa(tmpInt[2]))
		if v[0] == tmpInt[0] && v[1] == tmpInt[1] && v[2] == tmpInt[2] {
			check = false
		}
		fmt.Println(check)
		fmt.Println(threeInt)
		fmt.Println("------------")
	}
	return check
}
