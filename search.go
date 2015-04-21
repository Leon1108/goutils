package goutils

// 查找'地板值'
// 也就是数组nums中最大的，小于num的值。
// @param val 地板值
// @param idx 其在数组中的索引
func SearchFloor(num int64, nums []int64) (val, idx int64) {
	// TODO 先对nums排序
	start := 0
	end := len(nums) - 1
	for {
		mid := (start + end) / 2

		if nums[mid] == num {
			return nums[mid], int64(mid)
		}

		if end-start == 1 {
			if num >= nums[end] {
				return nums[end], int64(end)
			} else {
				return nums[start], int64(start)
			}
		}

		if num < nums[mid] {
			end = mid // 找前半段
		} else if num > nums[mid] {
			start = mid // 找后半段
		}
	}
}
