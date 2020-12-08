package main

import (
	"algorithm/array"
	"fmt"
)

func main()  {
	strSlice :=[]string{"abcwqwda1","abcdafwad2","abbdf233"}

	fmt.Println(array.LongestCommonPrefix(strSlice))
}
