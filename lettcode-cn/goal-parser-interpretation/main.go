package main

func interpret(command string) string {
	var ans string
	for i := 0; i < len(command); {
		switch command[i] {
		case 'G':
			ans += "G"
			i += 1
		case '(':
			if command[i+1] == ')' {
				i += 2
				ans += "o"
			} else {
				i += 4
				ans += "al"
			}
		}
	}
	return ans
}
