https://leetcode.cn/problems/asteroid-collision/

负数只会和她左面的正数 发生碰撞。
那么维护一个栈就可以了。
遇到正数就入栈，遇到负数就出栈 发生碰撞，如果是个空栈，也直接入栈。
1次遍历就可以了。