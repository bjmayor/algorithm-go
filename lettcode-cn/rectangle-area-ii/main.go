package main

import "sort"

type seg []struct{ cover, len, maxLen int }

func (t seg) init(hBound []int, idx, l, r int) {
	if l == r {
		t[idx].maxLen = hBound[l] - hBound[l-1]
		return
	}
	mid := (l + r) / 2
	t.init(hBound, idx*2, l, mid)
	t.init(hBound, idx*2+1, mid+1, r)
	t[idx].maxLen = t[idx*2].maxLen + t[idx*2+1].maxLen
}

func (t seg) update(idx, l, r, ul, ur, diff int) {
	if l > ur || r < ul {
		return
	}
	if ul <= l && r <= ur {
		t[idx].cover += diff
		t.pushUp(idx, l, r)
		return
	}
	mid := (l + r) / 2
	t.update(idx*2, l, mid, ul, ur, diff)
	t.update(idx*2+1, mid+1, r, ul, ur, diff)
	t.pushUp(idx, l, r)
}

func (t seg) pushUp(idx, l, r int) {
	if t[idx].cover > 0 {
		t[idx].len = t[idx].maxLen
	} else if l == r {
		t[idx].len = 0
	} else {
		t[idx].len = t[idx*2].len + t[idx*2+1].len
	}
}

func rectangleArea(rectangles [][]int) (ans int) {
	n := len(rectangles) * 2
	hBound := make([]int, 0, n)
	for _, r := range rectangles {
		hBound = append(hBound, r[1], r[3])
	}
	// 排序，方便下面去重
	sort.Ints(hBound)
	m := 0
	for _, b := range hBound[1:] {
		if hBound[m] != b {
			m++
			hBound[m] = b
		}
	}
	hBound = hBound[:m+1]
	t := make(seg, m*4)
	t.init(hBound, 1, 1, m)

	type tuple struct{ x, i, d int }
	sweep := make([]tuple, 0, n)
	for i, r := range rectangles {
		sweep = append(sweep, tuple{r[0], i, 1}, tuple{r[2], i, -1})
	}
	sort.Slice(sweep, func(i, j int) bool { return sweep[i].x < sweep[j].x })

	for i := 0; i < n; i++ {
		j := i
		for j+1 < n && sweep[j+1].x == sweep[i].x {
			j++
		}
		if j+1 == n {
			break
		}
		// 一次性地处理掉一批横坐标相同的左右边界
		for k := i; k <= j; k++ {
			idx, diff := sweep[k].i, sweep[k].d
			// 使用二分查找得到完整覆盖的线段的编号范围
			left := sort.SearchInts(hBound, rectangles[idx][1]) + 1
			right := sort.SearchInts(hBound, rectangles[idx][3])
			t.update(1, 1, m, left, right, diff)
		}
		ans += t[1].len * (sweep[j+1].x - sweep[j].x)
		i = j
	}
	return ans % (1e9 + 7)
}

func main() {
	println(rectangleArea([][]int{{0, 0, 2, 2}, {1, 0, 2, 3}, {1, 0, 3, 1}})) // 6
}
