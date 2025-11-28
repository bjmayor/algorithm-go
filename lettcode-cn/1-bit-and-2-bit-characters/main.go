package main

func isOneBitCharacter(bits []int) bool {
	n := len(bits)
	i := 0
	for i < n-1 {
		if bits[i] == 0 {
			i++
		} else {
			i += 2
		}
	}
	return i == n-1
}
