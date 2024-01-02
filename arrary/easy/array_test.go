package easy

import "testing"

//存在重复元素2
//https://leetcode.cn/problems/contains-duplicate-ii/solutions/1218075/cun-zai-zhong-fu-yuan-su-ii-by-leetcode-kluvk/

func TestContainsNearbyDuplicate(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	k := 3
	t.Log(containsNearbyDuplicate(nums, k))
}

func containsNearbyDuplicate(nums []int, k int) bool {
	//辅助hash
	m := make(map[int]int)
	//遍历数组
	for i, v := range nums {
		//判断值是否存在
		if _, ok := m[v]; ok {
			//判断索引差值是否小于等于k
			if i-m[v] <= k {
				return true
			}
		}
		//不存在则记录索引
		m[v] = i
	}
	return false
}
