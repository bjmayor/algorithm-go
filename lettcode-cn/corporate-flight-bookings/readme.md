https://leetcode.cn/problems/corporate-flight-bookings/

理解差分数组就一个点。

已知原数组是 [1,2,2,4]
求其差分数组为 [1,1,0,2], 通过求前缀和，可以得到还原数组，得到[1,2,2,4]

那么需求在[1,3] 都增加5, 
对差分数组的影响 就是 
diff[l] 增5,因为origin[l] 增加了5， origin[l-1]没变，所以是diff[l]+5
diff[l+1] - diff[r] 都不变。 ------origin[l] - origin[r] 都增加了5，
diff[r+1] -= 5 , 因为orgin[r]增加了5， 而origin[r+1]没变。所以 diff[r+1] = origin[r+1]-origin[r] 需要减5.

通过求前缀和就可以得到原数组了。

总结：
1. 原数组求差分数组。
2. 原数组区间改变对差分数组的影响
3. 通过对差分数组求前缀和 可以得到原数组。