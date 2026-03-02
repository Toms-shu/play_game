package homework01

import (
	"fmt"
	"sort"
	"strconv"
)

// 1. 只出现一次的数字
// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
func SingleNumber(nums []int) int {
	// TODO: implement

	/*
		查找非空整数数组中只出现一次的元素
	*/
	fmt.Println("原始数组：", nums)
	om := make(map[int]int)
	var os []int
	for _, v := range nums {
		fmt.Println(v)
		_, exist := om[v]
		if exist {
			delete(om, v)
		} else {
			om[v] = 1
		}
	}
	fmt.Println(om)
	for v, _ := range om {
		os = append(os, v)
	}
	fmt.Println("结果：", os)
	return os[0]
}

// 2. 回文数
// 判断一个整数是否是回文数
func IsPalindrome(x int) bool {
	// TODO: implement
	/*
		回文数判断
		1.长度，奇数位，
		2.长度，偶数位，完全对称
	*/
	boolValue := false

	str := strconv.Itoa(x)
	var sliceStr []string
	for _, v := range str {
		sliceStr = append(sliceStr, string(v))
	}
	fmt.Printf("str:%v, %T\n", sliceStr, sliceStr)
	l := len(sliceStr)
	s := l / 2

	for i := 0; i < s; i++ {
		fmt.Println(i, sliceStr[i], ";", l-1-i, ":", sliceStr[l-1-i])
		if sliceStr[l-1-i] == sliceStr[i] {
			boolValue = true
		}
	}
	fmt.Println(boolValue)
	return boolValue
}

// 3. 有效的括号
// 给定一个只包括 '(', ')', '{', '}', '[', ']' 的字符串，判断字符串是否有效
func IsValid(s string) bool {
	// TODO: implement
	m1 := make(map[string]int)
	//kh := []string{"(", ")", "[", "]", "{", "}"}
	m2 := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}
	var s1 []string
	for _, v := range s {
		switch string(v) {
		case "(":
			s1 = append(s1, string(v))
		case ")":
			s1 = append(s1, string(v))
		case "[":
			s1 = append(s1, string(v))
		case "]":
			s1 = append(s1, string(v))
		case "{":
			s1 = append(s1, string(v))
		case "}":
			s1 = append(s1, string(v))
		}

		r, exist := m1[string(v)]
		if exist {
			m1[string(v)] = r + 1
		} else {
			m1[string(v)] = 0
		}
	}

	l := len(s1)
	if l%2 != 0 {
		return false
	}
	fmt.Println("括号数组：", s1)
	fmt.Println("长度：", l/2)

	var left []string
	for _, v := range s1 {
		if leftValue, isRight := m2[v]; isRight {
			fmt.Println(leftValue, left[len(left)-1])
			if len(left) == 0 || left[len(left)-1] != leftValue {
				return false
			}
			//匹配上了，需要出栈
			left = left[:len(left)-1]
		} else {
			// 左括号，入栈
			left = append(left, v)
		}
	}
	//栈为空表示全部匹配
	return len(left) == 0
}

// 4. 最长公共前缀
// 查找字符串数组中的最长公共前缀
func LongestCommonPrefix(strs []string) string {
	// TODO: implement
	if len(strs) == 0 {
		fmt.Println("空数组！")
		return ""
	}
	if len(strs) == 1 {
		fmt.Println("单元素数组！")
		return strs[0]
	}
	v1 := strs[0]
	lStart := len(v1)
	prefix := ""
	for _, v := range strs[1:] {
		fmt.Println("起始元素：", v1, "对比元素：", v)
		lCompare := len(v)
		var l int
		if lStart >= lCompare {
			l = lCompare
		} else {
			l = lStart
		}

		fmt.Println("较小长度", l)

		prefix = ""
		for index := 0; index < l; index++ {
			//fmt.Println(index)
			vStart := v1[index]
			vCompare := v[index]
			fmt.Println(string(vStart), string(vCompare))
			if vStart != vCompare {
				//prefix = prefix[:index]
				fmt.Println("退出", prefix)
				break
			} else {
				prefix = prefix + string(vStart)
				fmt.Println("继续比较:", prefix)
			}
		}
		v1 = prefix

		lStart = len(prefix)

	}
	fmt.Println(prefix)

	return prefix
}

// 5. 加一
// 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
func PlusOne(digits []int) []int {
	// TODO: implement
	for index, v := range digits {
		digits[index] = v + 1
		fmt.Println(index, v)
	}
	fmt.Println(digits)
	return digits
}

// 6. 删除有序数组中的重复项
// 给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。
// 不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
func RemoveDuplicates(nums []int) int {
	// TODO: implement
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}

	fmt.Println(nums[:i+1])
	return len(nums[:i+1])
}

// 7. 合并区间
// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
// 请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
func Merge(intervals [][]int) [][]int {
	// TODO: implement
	if len(intervals) == 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	fmt.Println("sorted:", intervals)

	i := 0
	for j := 1; j < len(intervals); j++ {
		if intervals[j][0] > intervals[i][1] {
			i++
			intervals[i] = intervals[j]
		} else {
			intervals[i][1] = intervals[j][1]
		}
	}

	fmt.Println(intervals[:i+1])
	return intervals[:i+1]
}

// 8. 两数之和
// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
func TwoSum(nums []int, target int) []int {
	// TODO: implement
	fmt.Println(nums)
	var targetSlice [][]int
	match := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		v := nums[i]
		r := target - v

		if _, exist := match[v]; exist {
			targetSlice = append(targetSlice, []int{r, v})
		}
		match[r] = v
	}
	fmt.Println(targetSlice)
	return targetSlice[0]
}
