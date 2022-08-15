package main

import (
	"fmt"
	"sort"
)

type Pair struct {
	Pos int
	Size int
}
func groupThePeople(groupSizes []int) [][]int {
	pairs := make([]Pair, len(groupSizes))
	for i,size := range groupSizes {
		pairs[i] = Pair{
			Pos:  i,
			Size: size,
		}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Size < pairs[j].Size
	})
	var buckets [][]int
	bucket := []int{}
	for _, v := range pairs {
		if len(bucket)<= v.Size {
			bucket = append(bucket, v.Pos)
			if len(bucket)== v.Size {
				buckets = append(buckets, bucket)
				bucket = nil
			}
		}
	}
	return buckets
}

func main()  {
	fmt.Println(groupThePeople([]int{3,3,3,3,3,1,3}))
}