https://leetcode-cn.com/problems/throne-inheritance/

其实就是一课树。
按照长子先继位的规则。
长子那支都没了，再次子。

把题目中的curOrder理解成都死了就好理解了。

Successor(x, curOrder):
    如果 x 没有孩子或者x的孩子都挂了:
        如果 x 是国王，那么返回 null
        否则，返回 Successor(x 的父亲, curOrder)
    否则，返回 x 不在 curOrder 中最年长的孩子

这么理解实现就简单了呀。。 就是树的先序遍历。先遍历根，然后左边的节点，然后右边的节点。 只是子节点不止2个。

提交可解。但是只击败11%。
把递归改成for循环可以提高性能。也只是击败16.67%。

