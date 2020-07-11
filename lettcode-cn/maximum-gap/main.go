package main

import (
	"fmt"
	"math"
)

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	var maxv int
	for i:=0;i<len(nums);i++ {
		if nums[i]	> maxv {
			maxv = nums[i]
		}
	}

	//基数排序(计数排序，桶排序），都是基于非比较排序，都有各自的限制场景
	//计数排序。要求数字大小范围比较小, 如果数据范围k 比要排序的n 大很多就不适合。 这个题不适合。k太大, n>=2
	//基数排序。 需要可以分割独立的"位"来比较，而且位之间有递进的关系。这个题可以。radix=10, 最多排10趟够了。
	//桶排序。分成k个桶，每个桶分别排序。容易划分成m个桶，并且划分比较均匀。
	exp := 1
	for maxv / exp > 0{
		tmp := make([][]int, 10)
		for _, v := range nums {
			d := v/exp%10
			tmp[v/exp%10] = append(tmp[d], v)
		}
		exp *= 10
		i:=0
		for _, link := range tmp {
			for _, v := range link {
				nums[i] = v
				i++
			}
		}
	}



	//遍历比较O(n)
	pre := nums[0]
	var result int
	for i:=1;i<len(nums);i++ {
		if nums[i]	- pre > result  {
			result = nums[i]-pre
		}
		pre = nums[i]
	}
	return result
}
//桶排序
func maximumGap2(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	max := nums[0]
	min := nums[0] //找到最大和最小
	for _, v := range nums {
		if max < v {
			max = v
		}
		if min > v {
			min = v
		}
	}
	if min == max {
		return 0
	}
	c_max := make([]int, len(nums)+1) //记录容器中的最大值
	c_min := make([]int, len(nums)+1) //记录容器中的最小值
	has := make([]bool, len(nums)+1)  //某个桶内是否已经装有数据,默认值false
	for _, v := range nums {
		cid := CountCid(v, max, min, len(nums)) //求容器编号
		if has[cid] == false {                  //该容器无数据则最大，最小都为v，更改标记
			c_max[cid] = v
			c_min[cid] = v
			has[cid] = true
		} else {					//该容器已有数据，则比较，更新
			if c_max[cid] < v {
				c_max[cid] = v
			}
			if c_min[cid] > v {
				c_min[cid] = v
			}
		}
	}
	res := 0
	lastMax := c_max[0]
	i := 1
	for ; i <= len(nums); i++ {
		if has[i] {
			res = int(math.Max(float64(res), float64(c_min[i]-lastMax)))
			lastMax = c_max[i]
		}
	}
	return res
}
func CountCid(num, max, min, len int) (cid int) {
	return (num - min) * (len + 1) / (max - min + 1)
}

func main()  {
	fmt.Println(maximumGap([]int{3,6,9,1}))//3
	fmt.Println(maximumGap([]int{10}))//0
	fmt.Println(maximumGap([]int{1, 10}))//999
}