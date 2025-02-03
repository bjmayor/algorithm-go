package main

import "fmt"

func getFolderNames(names []string) (ans []string) {
	uniq := map[string]int{}
	ans = make([]string, len(names))
	for i, name := range names {
		if v, ok := uniq[name]; !ok {
			ans[i] = name
			uniq[name] = 0
		} else {
			for idx := v + 1; ; idx++ {
				newName := fmt.Sprintf("%s(%d)", name, idx)
				if _, ok := uniq[newName]; !ok {
					ans[i] = newName
					uniq[name] = idx
					uniq[newName] = 0
					break
				}
			}
		}

	}
	return
}

func main() {
	fmt.Println(getFolderNames([]string{"gta", "gta(1)", "gta", "avalon"}))
}
