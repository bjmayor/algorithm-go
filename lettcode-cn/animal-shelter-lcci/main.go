package main
type AnimalShelf struct {
	animals [][]int
}


func Constructor() AnimalShelf {
	obj := AnimalShelf{}
	obj.animals = make([][]int, 0)
	return obj
}


func (this *AnimalShelf) Enqueue(animal []int)  {
	this.animals = append(this.animals, animal)
}


func (this *AnimalShelf) DequeueAny() []int {
	if len(this.animals) == 0 {
		return []int{-1,-1}
	} else {
		head := this.animals[0]
		this.animals = this.animals[1:]
		return head
	}
}


func (this *AnimalShelf) DequeueDog() []int {
	if len(this.animals) == 0 {
		return []int{-1,-1}
	} else {
		for i:=0;i<len(this.animals);i++ {
			if this.animals[i][1] == 1 {
				result := this.animals[i]
				this.animals = append(this.animals[:i], this.animals[i+1:]...)
				return result
			}
		}
		return []int{-1,-1}
	}
}


func (this *AnimalShelf) DequeueCat() []int {
	if len(this.animals) == 0 {
		return []int{-1,-1}
	} else {
		for i:=0;i<len(this.animals);i++ {
			if this.animals[i][1] == 0 {
				result := this.animals[i]
				this.animals = append(this.animals[:i], this.animals[i+1:]...)
				return result
			}
		}
		return []int{-1,-1}
	}
}



