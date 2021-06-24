package main

import (
	"fmt"
	"sort"
)
// 原始做法
func permutation1(s string) []string {
	var set = make(map[string]struct{})
	for i, _ := range s {
		char := s[i:i+1]
		if len(set)	== 0 {
			set[char] = struct{}{}
		} else {
			tmp := make(map[string]struct{})
			for k, _ := range set {
				vv := k+char
				if _, ok := tmp[vv]; !ok {
					tmp[vv] = struct{}{}
				}
				for j,_ := range k {
					nn := k[0:j]+char+k[j:]
					if _, ok := tmp[nn]; !ok {
						tmp[nn] = struct{}{}
					}
				}
			}
			set = tmp
		}
	}
	result := make([]string, len(set), len(set))
	i := 0
	for k,_ := range set {
		result[i] = k
		i++
	}
	return result
}

// 回溯法
func permutation2(s string) (ans []string) {
	t := []byte(s)
	sort.Slice(t, func(i, j int) bool { return t[i] < t[j] })
	n := len(t)
	perm := make([]byte, 0, n)
	vis := make([]bool, n)
	var backtrack func(int)
	backtrack = func(i int) {
		if i == n {
			ans = append(ans, string(perm))
			return
		}
		for j, b := range vis {
			if b || j > 0 && !vis[j-1] && t[j-1] == t[j] {
				continue
			}
			vis[j] = true
			perm = append(perm, t[j])
			backtrack(i + 1)
			perm = perm[:len(perm)-1]
			vis[j] = false
		}
	}
	backtrack(0)
	return
}

func reverse(a []byte) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

func nextPermutation(nums []byte) bool {
	n := len(nums)
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i < 0 {
		return false
	}
	j := n - 1
	for j >= 0 && nums[i] >= nums[j] {
		j--
	}
	nums[i], nums[j] = nums[j], nums[i]
	reverse(nums[i+1:])
	return true
}
// 下一个排列
func permutation(s string) (ans []string) {
	t := []byte(s)
	sort.Slice(t, func(i, j int) bool { return t[i] < t[j] })
	for {
		ans = append(ans, string(t))
		if !nextPermutation(t) {
			break
		}
	}
	return
}


func main()  {
	fmt.Println(permutation("abc"))
	fmt.Println(permutation("aaa"))
}