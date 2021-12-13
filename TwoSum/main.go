package main

import "fmt"

func main() {
	//驗證
	nums := []int{
		1, 2, 5, 7,
	}
	target := 12
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		c := target - nums[i]
		if _, ok := m[c]; ok {
			return []int{m[c], i}
		}
		m[nums[i]] = i
	}
	return nil
}
