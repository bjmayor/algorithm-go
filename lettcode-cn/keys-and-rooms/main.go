package main

import "fmt"

func canVisitAllRooms(rooms [][]int) bool {
	visited := make([]bool, len(rooms))
	visited[0] = true
	q := []int{0}

	for len(q) > 0 {
		head := q[0]
		q = q[1:]
		for _, key := range rooms[head] {
			if !visited[key] {
				visited[key] = true
				q = append(q, key)
			}
		}
	}

	for _, v := range visited {
		if !v {
			return false
		}
	}
	return true
}

func main() {
	// rooms := [][]int{{1}, {2}, {3}, {}}
	// fmt.Println(canVisitAllRooms(rooms))

	rooms2 := [][]int{{1, 3}, {3, 0, 1}, {2}, {0}}
	fmt.Println(canVisitAllRooms(rooms2))
}
