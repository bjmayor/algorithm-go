https://leetcode.cn/problems/implement-magic-dictionary

1. 老规矩，暴力法来一波。提交居然通过。
击败72%,65%

2. 假定i字符不匹配，那么就得0,i-1和i+1,n 字符串都需要匹配，这个可应用字典树。
字典树的实现也比较简单，就是空间换 时间。
   回溯
    保持状态+进入下一步
   

    