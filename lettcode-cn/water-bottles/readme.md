https://leetcode-cn.com/problems/water-bottles/

简单难度。试下笨方法。直接for循环，可解。

更优解:
 b-n(e-1) >= e 
b 为numBottles,
e 为 numExchange
等号左边 表示 空瓶子数量，因为 e瓶兑换完再喝掉 余1个空瓶。所以兑换一次消耗的是e-1 个空瓶子。
n 表示能兑换的次数。。

找最下的n, 使得
b-n(e-1) < e 即可
