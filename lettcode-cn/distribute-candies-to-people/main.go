package main

import "fmt"

func distributeCandies(candies int, num_people int) []int {
	total := (1+num_people)*num_people/2
	q := 1
	left := candies
	for left > total {
		q += 1
		left -= total
		total += num_people*num_people
	}

	result := make([]int, num_people)
	left  = candies
	if q > 1 {
		for i:=0;i<num_people;i++ {
			result[i] = (q-1)*(i+1) + (q-1)*(q-2)/2*num_people  //i,n+i,2n+i,...(q-2)*n+i
			left -= result[i]
		}
	}
	base := (q-1)*num_people

	for i:=0; i<num_people; i++ {
		if left > base+i+1 {
			result[i] += base+i+1
		} else {
			result[i] += left
			break
		}
		left -= i+1+base
	}

	return result
}

func main()  {
	fmt.Println(distributeCandies(7, 4)) //[1,2,3,1]
	fmt.Println(distributeCandies(10, 3)) //[5,2,3]
	fmt.Println(distributeCandies(60, 4)) //[15,18,15,12]
}