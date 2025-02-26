package main

import "fmt"

type BrowserHistory struct {
	current int
	links   []string
}

func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{
		current: 0,
		links:   append([]string{}, homepage),
	}
}

func (this *BrowserHistory) Visit(url string) {
	this.links = this.links[:this.current+1]
	this.links = append(this.links, url)
	this.current++
}

func (this *BrowserHistory) Back(steps int) string {
	if steps >= this.current {
		this.current = 0
	} else {
		this.current -= steps
	}
	return this.links[this.current]
}

func (this *BrowserHistory) Forward(steps int) string {
	if this.current+steps >= len(this.links) {
		this.current = len(this.links) - 1
	} else {
		this.current += steps
	}
	return this.links[this.current]
}

func main() {
	browserHistory := Constructor("leetcode.com")
	browserHistory.Visit("google.com")
	browserHistory.Visit("facebook.com")
	browserHistory.Visit("youtube.com")
	fmt.Println(browserHistory.Back(1))
	fmt.Println(browserHistory.Back(1))
	fmt.Println(browserHistory.Forward(1))
	browserHistory.Visit("linkedin.com")
	fmt.Println(browserHistory.Forward(2))
	fmt.Println(browserHistory.Back(2))
	fmt.Println(browserHistory.Back(7))
}
