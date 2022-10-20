package main

func countStudents(students []int, sandwiches []int) int {
	for len(sandwiches) > 0 {
		try := 0
		for sandwiches[0] != students[0] && try < len(students) {
			try++
			students = append(students[1:], students[0])
		}
		if sandwiches[0] == students[0] {
			students = students[1:]
			sandwiches = sandwiches[1:]
		} else {
			return len(students)
		}
	}
	return 0
}
