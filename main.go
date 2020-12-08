package main

import (
	"algorithm/array"
	"fmt"
)

func main()  {
	strSlice :=[]string{"abcwqwda111","abcdafwad222","abbdf333"}

	fmt.Println(array.LongestCommonPrefix(strSlice))
}
