package goutils

import (
	"errors"
	"fmt"
)

// 查找'地板值'
// 也就是数组nums中最大的，小于num的值。
// @param val 地板值，如果没有找到则返回-1
// @param idx 其在数组中的索引, 如果没有找到则返回-1
// @param err 当没有找到时，一般是由于需要查找的值小于数组中的最小值，返回err
func SearchFloor(num int64, nums []int64) (val, idx int64, err error) {
	// TODO 先对nums排序
	start := 0
	end := len(nums) - 1
	for {
		mid := (start + end) / 2
		fmt.Println("start: ", start, "; end: ", end, "; mid : ", mid)

		if nums[mid] == num {
			return nums[mid], int64(mid), nil
		}

		if end-start == 1 {
			if num >= nums[end] {
				return nums[end], int64(end), nil
			} else if num >= nums[start] {
				return nums[start], int64(start), nil
			} else {
				return -1, -1, errors.New("Not Found!")
			}
		}

		if num < nums[mid] {
			end = mid // 找前半段
		} else if num > nums[mid] {
			start = mid // 找后半段
		}
	}
}
