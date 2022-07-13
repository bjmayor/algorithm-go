package main

//type MagicDictionary struct {
// 	dict map[string]bool
//}
//
//
//func Constructor() MagicDictionary {
//	return MagicDictionary{dict: map[string]bool{}}
//}
//
//
//func (this *MagicDictionary) BuildDict(dictionary []string)  {
//	for _, v := range dictionary {
//		this.dict[v] =true
//	}
//}
//
//
//func (this *MagicDictionary) Search(searchWord string) bool {
//	for k, _ := range this.dict {
//		if len(k) == len(searchWord) {
//			n := 0
//			for i, ch := range  []byte(searchWord) {
//				if ch != k[i] {
//					n++
//				}
//			}
//			if n ==1 {
//				return true
//			}
//		}
//	}
//	return false
//}

type Node struct {
	ch       uint8           //当前存的字符
	children map[uint8]*Node // 后续可能的字符
	end      bool            // 是否是一个结果。譬如单词
}
type MagicDictionary struct {
	head *Node
}

func Constructor() MagicDictionary {
	return MagicDictionary{head: new(Node)}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	this.head.children = make(map[uint8]*Node)
	for _, v := range dictionary {
		cur := this.head
		for i := 0; i < len(v); i++ {
			ch := v[i]
			if n, ok := cur.children[ch]; ok {
				cur = n
			} else {
				node := new(Node)
				node.ch = ch
				node.children = make(map[uint8]*Node)
				cur.children[ch] = node
				cur = node
			}
		}
		cur.end = true
	}
}

func (this *MagicDictionary) Search(searchWord string) bool {
		return dfs(this.head, searchWord, false)
}
func dfs(node *Node, searchWord string, modified bool) bool {
	if searchWord == "" {
		return modified && node.end
	}
	c := searchWord[0]
	if node.children[c] != nil && dfs(node.children[c], searchWord[1:], modified) {
		return true
	}
	if !modified {
		for _, child := range node.children {
			if child.ch != c && dfs(child, searchWord[1:], true) {
				return true
			}
		}
	}
	return false
}

/*
["MagicDictionary", "buildDict", "search", "search", "search", "search", "search", "search"]
[[], [["hello","hallo","leetcode","judge", "judgg"]], ["hello"], ["hallo"], ["hell"], ["leetcodd"], ["judge"], ["juggg"]]
 */
func main() {
	obj := Constructor()
	obj.BuildDict([]string{"hello", "leetcode"})
	println(obj.Search("hello"))     // false
	println(obj.Search("hhllo"))     // true
	println(obj.Search("hell"))     // false
	println(obj.Search("leetcoded"))     // false
}
