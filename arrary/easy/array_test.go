package easy

import (
	"fmt"
	"strconv"
	"testing"
)

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

//汇总区间
//https://leetcode.cn/problems/summary-ranges/
func TestSummaryRanges(t *testing.T) {
	fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))
	fmt.Println(summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
}

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	//遍历数组 如果i等于i-1的值+1,插入数组, 否则创建新数组
	var ret []string
	head := nums[0]
	tail := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1]+1 {
			tail = nums[i]
		} else {
			ret = append(ret, getStrForSummaryRanges(head, tail))
			head = nums[i]
			tail = nums[i]
		}
	}
	ret = append(ret, getStrForSummaryRanges(head, tail))
	return ret
}

func getStrForSummaryRanges(head, tail int) string {
	if head == tail {
		return strconv.Itoa(head)
	}
	return fmt.Sprintf("%d->%d", head, tail)
}
