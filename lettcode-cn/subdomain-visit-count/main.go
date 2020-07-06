package main

import (
	"fmt"
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	m := make(map[string]int)
	for _, s := range cpdomains {
		parts := strings.Split(s," ")
		times, _ := strconv.Atoi(parts[0])
		domains := strings.Split(parts[1],".")
		for i:=0;i<len(domains);i++{
			m[strings.Join(domains[i:],".")]+=times
		}
	}

	var res []string
	for k,v := range m {
		res = append(res, strconv.Itoa(v)+ " " + k)
	}
	return res
}

func main()  {
	fmt.Println(subdomainVisits([]string{
		"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org",
	}))
}
