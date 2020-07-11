package main

import "fmt"

func isRobotBounded(instructions string) bool {
	x := 0
	y := 0
	d := 0 //0北,1西,2南,3东
	for i:=0;i<4;i++ {
		for _, ins := range instructions {
			switch ins {
			case 'G':
				{
					switch d {
					case 0:
						y++
					case 1:
						x--
					case 2:
						y--
					case 3:
						x++
					}
				}
			case 'L':
				{
					d = (d + 1) % 4
				}
			case 'R':
				{
					d = (d - 1 + 4) % 4
				}
			}
		}
	}

	return x==0 && y==0 && d == 0
}



func main() {
	tests := []struct {
		Nums     string
		Expected bool
	}{
		{"GGLLGG", true},
		{"GG", false},
		{"GL", true},
		{"GLGLGGLGL", false},
	}
	for i, t := range tests {
		var real = isRobotBounded(t.Nums)
		if real != t.Expected {
			fmt.Println(i, " expected:", t.Expected, " real:", real, t.Nums)
		}
	}
}
