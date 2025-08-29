package main

import "fmt"

// 只出现一次的数字
func singleNumber(nums []int) int {
	length := len(nums) / 2
	m := make(map[int]int, length+1)
	res := 0
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			delete(m, nums[i])
		} else {
			m[nums[i]] = 1
		}
	}
	for k, _ := range m {
		res = k
	}
	return res
}

// 回文数
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	xArr := make([]int, 0)
	for x > 0 {
		xArr = append(xArr, x%10)
		x = x / 10
	}
	length := len(xArr)
	for i := 0; i < length; i++ {
		if xArr[i] != xArr[length-i-1] {
			return false
		}
	}
	return true
}

// 字符串有效的括号
func isValid(s string) bool {

	sByte := []byte(s)
	sLen := len(sByte)
	list := make([]byte, 0, sLen/2+1)
	if sLen <= 1 {
		return false
	}

	isTrue := true
	for i := 0; i < sLen; i++ {
		// fmt.Println(string(sByte[i]))
		if string(sByte[i]) == "(" || string(sByte[i]) == "{" || string(sByte[i]) == "[" {
			list = append(list, sByte[i])
			continue
		}
		if (string(sByte[i]) == "}" || string(sByte[i]) == ")" || string(sByte[i]) == "]") && len(list) == 0 {
			isTrue = false
			break
		}
		sright := string(list[len(list)-1])
		if string(sByte[i]) == ")" && sright != "(" || string(sByte[i]) == "}" && sright != "{" || string(sByte[i]) == "]" && sright != "[" {
			isTrue = false
			break
		} else {
			list = list[:len(list)-1]
		}
	}
	if len(list) != 0 {
		isTrue = false
	}

	return isTrue
}

// 最长公共前缀
// func longestCommonPrefix(strs []string) string {
// 	if len(strs) == 0 {
// 		return ""
// 	}

// 	prefix := strs[0]
// 	for i := 1; i < len(strs); i++ {
// 		prefix = comparePrefix(prefix, strs[i])
// 		if prefix == "" {
// 			break
// 		}
// 	}
// 	return prefix
// }

// func comparePrefix(a, b string) string {
// 	minLength := len(a)
// 	if minLength > len(b) {
// 		minLength = len(b)
// 	}
// 	i := 0
// 	for i < minLength && a[i] == b[i] {
// 		i++
// 	}
// 	return a[:i]
// }

/**
* 加一
 */
// func plusOne(digits []int) []int {
// 	length := len(digits)
// 	for i := length - 1; i >= 0; i-- {
// 		if digits[i] < 9 {
// 			digits[i] = digits[i] + 1
// 			return digits
// 		}
// 		digits[i] = 0
// 	}
// 	return append([]int{1}, digits...)
// }

/*
* 删除数组重复项
 */
// func removeDuplicates(nums []int) int {
// 	if len(nums) == 0 {
// 		return 0
// 	}
// 	i := 0
// 	for j := 1; j < len(nums); j++ {
// 		if nums[i] != nums[j] {
// 			i++
// 			nums[i] = nums[j]
// 		}
// 	}
// 	return i + 1
// }

/*
* 合并区间
*
 */
// func merge(intervals [][]int) [][]int {
// 	if len(intervals) == 0 {
// 		return [][]int{}
// 	}

// 	sort.Slice(intervals, func(i, j int) bool {
// 		return intervals[i][0] < intervals[j][0]
// 	})

// 	res := [][]int{intervals[0]}
// 	j := 0
// 	for i := 1; i < len(intervals); i++ {

// 		if intervals[i][0] <= res[j][1] && intervals[i][1] > res[j][1] {
// 			res[j][1] = intervals[i][1]
// 		}
// 		if intervals[i][0] > res[j][1] {
// 			res = append(res, intervals[i])
// 			j++
// 		}
// 	}
// 	return res
// }

func twoSum(nums []int, target int) []int {
	if len(nums) <= 1 {
		return []int{}
	}
	res := []int{}
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if v, ok := m[target-nums[i]]; ok {
			return append(res, v, i)
		}
		m[nums[i]] = i
	}
	return res
}

func main() {
	// 只出现一次的数字
	// nums := []int{4, 1, 2, 1, 2}
	// nu := singleNumber(nums)
	// fmt.Println(nu)

	// 回文数
	// x := 10
	// fmt.Println(isPalindrome(x))

	// 字符串，有效的括号
	// s := "{}"
	// fmt.Println(isValid(s))

	//最长公共前缀
	// strs := []string{"flower", "flow", "flight"}
	// fmt.Println(longestCommonPrefix(strs))

	//加一
	// digits := []int{9, 9, 9, 8}
	// fmt.Println(plusOne(digits))

	//删除有序数组中的重复项
	// nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	// fmt.Println(removeDuplicates(nums))

	//合并区间
	// intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	// fmt.Println(merge(intervals))

	//两数之和
	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, 9))
}
