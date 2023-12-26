# 给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。

## 暴力递归
```
func lengthOfLIS(nums []int) int {
    result := 0
    for i := 0; i < len(nums); i++ {
        len := a(nums, i)
        if result < len {
            result = len
        }
    }
    return result+1
}

func a(nums []int, i int) int {
    if i == len(nums) {
        return 0
    }
    maxLen := 0
    for j := i + 1; j < len(nums); j++ {
        if nums[j] > nums[i] {
            len := 1 + a(nums, j)
            if len > maxLen {
                maxLen = len
            }
        }
    }
    return maxLen
}

```
* 超时了
