package main

func peakIndexInMountainArray(arr []int) int {
	left := 0
	right := len(arr)-1
	for left < right {
		mid := left + (right-left)/2
		if arr[mid-1] <arr[mid]  && arr[mid] < arr[mid+1] {
			left = mid
		} else if  arr[mid-1] <arr[mid]  && arr[mid] > arr[mid+1] {
			return mid
		} else {
			right = mid
		}
	}
	return left
}

func main()  {
	println(peakIndexInMountainArray([]int{24,69,100,99,79,78,67,36,26,19}))
}