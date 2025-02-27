package main

import (
	"container/list"
	"fmt"
)

type TextEditor struct {
	cursor *list.Element // 光标位置
	text   *list.List    // 文本内容
}

func Constructor() TextEditor {
	return TextEditor{
		cursor: nil,
		text:   list.New(),
	}
}

func (this *TextEditor) AddText(text string) {
	for _, ch := range text {
		// 在光标位置后插入新字符
		if this.cursor == nil {
			this.cursor = this.text.PushBack(ch)
		} else {
			this.cursor = this.text.InsertAfter(ch, this.cursor)
		}
	}
}

func (this *TextEditor) DeleteText(k int) int {
	if this.cursor == nil {
		return 0
	}

	count := 0
	for count < k && this.cursor != nil {
		prev := this.cursor.Prev()
		this.text.Remove(this.cursor)
		this.cursor = prev
		count++
	}
	return count
}

func (this *TextEditor) CursorLeft(k int) string {
	// 向左移动光标
	for i := 0; i < k && this.cursor != nil; i++ {
		this.cursor = this.cursor.Prev()
	}
	return this.getLeftString()
}

func (this *TextEditor) CursorRight(k int) string {
	// 向右移动光标
	if this.cursor == nil {
		if this.text.Len() > 0 {
			this.cursor = this.text.Front()
			k--
		}
	}

	for i := 0; i < k && this.cursor != nil && this.cursor.Next() != nil; i++ {
		this.cursor = this.cursor.Next()
	}
	return this.getLeftString()
}

// 获取光标左侧最多10个字符
func (this *TextEditor) getLeftString() string {
	result := make([]rune, 0, 10)
	curr := this.cursor

	for i := 0; i < 10 && curr != nil; i++ {
		result = append([]rune{curr.Value.(rune)}, result...)
		curr = curr.Prev()
	}
	return string(result)
}

func main() {
	obj := Constructor()
	obj.AddText("bxyackuncqzcqo")
	fmt.Println(obj.CursorLeft(12)) // bx
	fmt.Println(obj.DeleteText(3))  // 2
	fmt.Println(obj.CursorLeft(5))
	obj.AddText("osdhyvqxf")
	fmt.Println(obj.CursorRight(10)) //yackuncqzc
}
