package main

import (
	"algorithm/array"
	"fmt"
	"sort"
)

func main()  {
	sliceNum1 :=[]int{1,2,2,1}

	sliceNum2 :=[]int{2,2}
	sort.Ints(sliceNum1)
	sort.Ints(sliceNum2)
	fmt.Println(array.IntersectPlus(sliceNum1,sliceNum2))
}
