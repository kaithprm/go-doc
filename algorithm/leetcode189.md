# 189. 轮转数组
## 题目:给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。
#### 暴力解法
```
func rotate(nums []int, k int)  {

l := len(nums)
// 1.将nums看作一个循环滚动的栈, k = k % len(nums)
k = k % l
// 外层基于k次数
for i := 1 ; i <= k; i++{
    // 2. 找到每次右移时的末尾数
    last := nums[l - 1]
    for j := l - 1; j > 0; j--{
        nums[j] = nums [j - 1] 
    }
    nums[0] = last
}
}
```
* 超时了
### 将一次一次右移变成统计后一次性右移，这样就少一层嵌套循环
* q:该如何确定一次性右移的末尾数呢？
#### 想法1:设置一个map,其实最多只有len(nums)种组合
* 也就是来动态建立一个长度为len(nums)的二维数组 其中包含了该数组的所有组合
* 二维数组切片每一个都要初始化 非常麻烦，而二维数组限制数还只能是常量 不能将len()作为限制
## 旋转数组方法
```
func rotate(nums []int, k int)  {


    k %= len(nums)
    reverse(nums)
    reverse(nums[:k])
    reverse(nums[k:])
}
* 先

func reverse( a []int){
    for i , n := 0,len(a);i < n/2 ; i++{
        a[i], a[n-1-i] = a[n-1-i], a[i]
    }
}
```
