package main

import (
	"fmt"
	"strings"
)

func printVertically(s string) []string {
	words := strings.Split(s, " ")
	var res []string
	i:=0
	for {
		finished := true
		data := []byte{}
		for _, word := range words {

			if i<len(word) {
				finished = false
				data = append(data,word[i])
			} else {
				data = append(data,' ')
			}
		}
		if finished {
			break
		}
		res = append(res, strings.TrimRight(string(data)," "))
		i++
	}
	return res
}

func main()  {
	fmt.Println(printVertically("HOW ARE YOU"))
	fmt.Println(printVertically("TO BE OR NOT TO BE"))
}