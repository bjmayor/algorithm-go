package main

import (
	"container/heap"
	"fmt"
)

// 食物信息结构体
type Food struct {
	name    string
	rating  int
	cuisine string
}

// 优先队列实现
type PriorityQueue []*Food

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// 按评分降序,评分相同按名字字典序
	return pq[i].rating > pq[j].rating || 
		(pq[i].rating == pq[j].rating && pq[i].name < pq[j].name)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Food)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

type FoodRatings struct {
	cuisineHeaps map[string]*PriorityQueue // 每个菜系对应一个大根堆
	foodMap      map[string]*Food          // 快速查找食物信息
}

func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	fr := FoodRatings{
		cuisineHeaps: make(map[string]*PriorityQueue),
		foodMap:      make(map[string]*Food),
	}
	
	for i := 0; i < len(foods); i++ {
		food := &Food{
			name:    foods[i],
			rating:  ratings[i],
			cuisine: cuisines[i],
		}
		fr.foodMap[foods[i]] = food
		
		if fr.cuisineHeaps[cuisines[i]] == nil {
			pq := make(PriorityQueue, 0)
			fr.cuisineHeaps[cuisines[i]] = &pq
		}
		heap.Push(fr.cuisineHeaps[cuisines[i]], food)
	}
	return fr
}

func (this *FoodRatings) ChangeRating(food string, newRating int) {
	// 更新食物评分
	foodInfo := this.foodMap[food]
	// 创建新的食物信息
	newFood := &Food{
		name:    food,
		rating:  newRating,
		cuisine: foodInfo.cuisine,
	}
	// 更新foodMap
	this.foodMap[food] = newFood
	// 直接将新食物加入堆中
	heap.Push(this.cuisineHeaps[foodInfo.cuisine], newFood)
}

func (this *FoodRatings) HighestRated(cuisine string) string {
	pq := this.cuisineHeaps[cuisine]
	// 移除评分不是最新的堆顶元素
	for pq.Len() > 0 {
		top := (*pq)[0]
		curFood := this.foodMap[top.name]
		if top.rating == curFood.rating {
			// 当前堆顶是最新评分
			return top.name
		}
		// 移除旧评分的堆顶元素
		heap.Pop(pq)
	}
	return "" // 不应该到达这里
}

func main() {
	obj := Constructor([]string{"kimchi", "miso", "sushi", "moussaka", "ramen", "bulgogi"}, []string{"korean", "japanese", "japanese", "greek", "japanese", "korean"}, []int{9, 12, 8, 15, 14, 7})
	fmt.Println(obj.HighestRated("korean"))   // kimchi
	fmt.Println(obj.HighestRated("japanese")) // ramen
	obj.ChangeRating("sushi", 16)
	fmt.Println(obj.HighestRated("japanese")) // sushi
	obj.ChangeRating("ramen", 16)
	fmt.Println(obj.HighestRated("japanese")) // ramen
}
