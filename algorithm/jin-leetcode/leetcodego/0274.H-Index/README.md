# [274. H-Index](https://leetcode.com/problems/h-index/)

## 题目

Given an array of citations (each citation is a non-negative integer) of a researcher, write a function to compute the researcher's h-index.

According to the definition of h-index on Wikipedia: "A scientist has index h if h of his/her N papers have at least h citations each, and the other N − h papers have no more than h citations each."

Example 1:

```c
Input: citations = [3,0,6,1,5]
Output: 3 
Explanation: [3,0,6,1,5] means the researcher has 5 papers in total and each of them had 
             received 3, 0, 6, 1, 5 citations respectively. 
             Since the researcher has 3 papers with at least 3 citations each and the remaining 
             two with no more than 3 citations each, her h-index is 3.
```

Note: 

If there are several possible values for h, the maximum one is taken as the h-index.



## 题目大意

求 h-index。h-index 值的定义：如果他/她的 N 篇论文中至少有 h 引用，而其他 N-h 论文的引用数不超过 h 引用数。

## 解题思路

可以先将数组里面的数从小到大排序。因为要找最大的 h-index，所以从数组末尾开始往前找，找到第一个数组的值，小于，总长度减去下标的值，这个值就是 h-index。

<!-- h指数由J.E. Hirsch于2005年提出，发表在《美国国家科学院院刊》上。[i] h指数是一种定量指标，基于对出版物和引用的出版物数据分析，以提供“对科学家累积研究贡献的重要性，重要性和广泛影响的估计”。[ii] 根据赫希的说法，h指数被定义为：“如果科学家的Np论文中有h篇至少有h次引用，而其他（Np – h）论文每篇都有≤h次引用，那么他就有索引h -->