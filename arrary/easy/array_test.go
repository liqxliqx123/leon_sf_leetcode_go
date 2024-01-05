package easy

import (
	"fmt"
	"sort"
	"testing"
)

// https://leetcode.cn/problems/pascals-triangle-ii/
func TestPascals(t *testing.T) {
	fmt.Println(getRow(0))
	fmt.Println(getRow(1))
	fmt.Println(getRow(2))
	fmt.Println(getRow(3))
	fmt.Println(getRow(4))
}

func getRow(rowIndex int) []int {
	ret := []int{1}
	//循环rowIndex次
	for i := 1; i <= rowIndex; i++ {
		//开头的1
		res := []int{1}
		if i > 1 {
			//循环上一个数组得到要插入的值
			for j := 0; j < len(ret)-1; j++ {
				res = append(res, ret[j]+ret[j+1])
			}
		}

		//结尾的1
		res = append(res, 1)
		//更新ret
		ret = res
	}

	return ret
}

// 股票最大利润
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/
func TestMaxProfit(t *testing.T) {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

func maxProfit(prices []int) int {
	//最大利润
	ret := 0
	//最小值
	minVal := prices[0]
	//遍历找到最小值
	for i := 1; i < len(prices); i++ {
		if prices[i] < minVal {
			minVal = prices[i]
		}
		//最大利润
		if prices[i]-minVal > ret {
			ret = prices[i] - minVal
		}
	}

	return ret
}

// 只出现一次的数字
// https://leetcode.cn/problems/single-number/
func TestSingleNumber(t *testing.T) {
	fmt.Println(singleNumber([]int{1}))
	fmt.Println(singleNumber([]int{2, 2, 1}))
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))
}

//思路1： stack 需排序和定义栈函数
//思路2： map 需要额外的存储空间

func singleNumber(nums []int) int {
	//排序
	sort.Ints(nums)
	ret := []int{nums[0]}
	//遍历数组，元素和前面元素一样则从结果集中删除最后一个元素；否则添加到数组
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			//删除最后一个元素
			ret = ret[0 : len(ret)-1]
		} else {
			ret = append(ret, nums[i])
		}
	}
	return ret[0]
}

//存在重复元素2 - map存值
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

//存在重复元素2 - 滑动窗口
//https://leetcode.cn/problems/contains-duplicate-ii/solutions/1218075/cun-zai-zhong-fu-yuan-su-ii-by-leetcode-kluvk/

func TestContainsNearbyDuplicateV2(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	k := 3
	t.Log(containsNearbyDuplicateV2(nums, k))
}

func containsNearbyDuplicateV2(nums []int, k int) bool {
	//辅助hash, 记录滑动窗口中的值
	m := make(map[int]int)
	//遍历数组
	for i, v := range nums {
		if i > k {
			//删除滑动窗口外的值
			delete(m, nums[i-k-1])
		}
		//判断值是否存在于map
		if _, ok := m[v]; ok {
			return true
		}
		//不存在则记录值
		m[v] = i

	}
	return false
}

// 丢失的数字
// https://leetcode.cn/problems/missing-number/
func TestMissingNumber(t *testing.T) {
	fmt.Println(missingNumber([]int{3, 0, 1}))
	fmt.Println(missingNumber([]int{0, 1}))
	fmt.Println(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
	fmt.Println(missingNumber([]int{0}))
}

func missingNumber(nums []int) int {
	//排序
	sort.Ints(nums)
	//遍历数组
	for i := 0; i < len(nums); i++ {
		//判断是否连续
		if nums[i] != i {
			return i
		}
	}
	return len(nums)
}
