package main
//递归法
func findTargetSumWays1(nums []int, target int) int {
	if len(nums)==0 {
		if target==0 {
			return 1
		} else {
			return 0
		}
	}
	last := len(nums)-1
	return findTargetSumWays1(nums[:last], target+nums[last]) + findTargetSumWays1(nums[:last], target-nums[last])
}
//直接遍历求和
func findTargetSumWays2(nums []int, target int) int {
	if len(nums)==0 {
		return 0
	}
	set := make(map[int]int) //k 为和, v为次数
	set[nums[0]]++
	set[-nums[0]]++
	for i:=1;i<len(nums);i++ {
		tmp := make(map[int]int)
		for k,_ := range set{
			tmp[k+nums[i]]+=max(set[k],1)
			tmp[k-nums[i]]+=max(set[k],1)
		}
		set = tmp
	}
	return set[target]
}

//转化问题后dp求解
func findTargetSumWays(nums []int, target int) int {
	var total int
	for i:=0;i<len(nums);i++ {
		total += nums[i]
	}
	if (total+target) % 2 ==1 || total < target {
		return 0
	}
	return findTargetSumPositiveWays(nums, (total+target)/2)
}
func findTargetSumPositiveWays(nums []int ,target int) int {
	var dp = make([]int, target+1, target+1)
	dp[0] = 1
	for i:=0;i<len(nums);i++ {
		for j:=target;j>=nums[i];j-- {
			dp[j] = dp[j]+dp[j-nums[i]]
		}
	}
	return dp[target]
}

func max(x,y int) int  {
	if x > y {
		return x
	}
	return y
}


func main()  {
	//nums:=[]int{1,1,1,1,1}
	//println(findTargetSumWays(nums,3))
	nums:=[]int{0,0,0,0,0,0,0,0,1}
	println(findTargetSumWays(nums,1))
}