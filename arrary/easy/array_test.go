package easy

import (
	"fmt"
	"sort"
	"strconv"
	"testing"
)

// 杨辉三角2
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

// 汇总区间
// https://leetcode.cn/problems/summary-ranges/
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

// 区域和检索 - 数组不可变
// https://leetcode.cn/problems/range-sum-query-immutable/description/
func TestSumRange(t *testing.T) {
	numArray := Constructor([]int{-2, 0, 3, -5, 2, -1})
	fmt.Println(numArray.SumRange(0, 2))
	fmt.Println(numArray.SumRange(2, 5))
	fmt.Println(numArray.SumRange(0, 5))
}

type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	return NumArray{
		nums: nums,
	}
}

func (n *NumArray) SumRange(left int, right int) int {
	sum := 0
	for i := left; i <= right; i++ {
		sum += n.nums[i]
	}
	return sum
}

// 两个数组的交集
// https://leetcode.cn/problems/intersection-of-two-arrays/description/
func TestIntersection(t *testing.T) {
	fmt.Println(intersection([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}

func intersection(nums1 []int, nums2 []int) []int {
	var ret []int

	m := make(map[int]struct{}, len(nums1))
	for _, v := range nums1 {
		m[v] = struct{}{}
	}
	for _, v := range nums2 {
		if _, ok := m[v]; ok {
			ret = append(ret, v)
			delete(m, v)
		}
	}
	return ret
}

// 两个数组的交集2
//https://leetcode.cn/problems/intersection-of-two-arrays-ii/
func TestIntersect(t *testing.T) {
	fmt.Println(intersect([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}

func intersect(nums1 []int, nums2 []int) []int {
	var ret []int

	m := make(map[int]int, len(nums1))
	for _, v := range nums1 {
		m[v] += 1
	}
	for _, v := range nums2 {
		if _, ok := m[v]; ok {
			ret = append(ret, v)
			m[v] -= 1
			if m[v] == 0 {
				delete(m, v)
			}
		}
	}
	return ret
}

// 第三大的数
//https://leetcode.cn/problems/third-maximum-number/
func TestThirdMax(t *testing.T) {
	fmt.Println(thirdMax([]int{3, 2, 1}))
	fmt.Println(thirdMax([]int{1, 2}))
	fmt.Println(thirdMax([]int{2, 2, 3, 1}))
}

func thirdMax(nums []int) int {
	//排序
	sort.Ints(nums)
	m := make(map[int]struct{}, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		if _, ok := m[nums[i]]; !ok {
			m[nums[i]] = struct{}{}
			if len(m) == 3 {
				return nums[i]
			}
		}
	}
	return nums[len(nums)-1]
}
