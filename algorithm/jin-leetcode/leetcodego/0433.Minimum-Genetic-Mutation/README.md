# [433. Minimum Genetic Mutation](https://leetcode.com/problems/minimum-genetic-mutation/)


## 题目

A gene string can be represented by an 8-character long string, with choices from `"A"`, `"C"`, `"G"`, `"T"`.

Suppose we need to investigate about a mutation (mutation from "start" to "end"), where ONE mutation is defined as ONE single character changed in the gene string.

For example, `"AACCGGTT"` -> `"AACCGGTA"` is 1 mutation.

Also, there is a given gene "bank", which records all the valid gene mutations. A gene must be in the bank to make it a valid gene string.

Now, given 3 things - start, end, bank, your task is to determine what is the minimum number of mutations needed to mutate from "start" to "end". If there is no such a mutation, return -1.

**Note:**

1. Starting point is assumed to be valid, so it might not be included in the bank.
2. If multiple mutations are needed, all mutations during in the sequence must be valid.
3. You may assume start and end string is not the same.

**Example 1:**

    start: "AACCGGTT"
    end:   "AACCGGTA"
    bank: ["AACCGGTA"]
    
    return: 1

**Example 2:**

    start: "AACCGGTT"
    end:   "AAACGGTA"
    bank: ["AACCGGTA", "AACCGCTA", "AAACGGTA"]
    
    return: 2

**Example 3:**

    start: "AAAAACCC"
    end:   "AACCCCCC"
    bank: ["AAAACCCC", "AAACCCCC", "AACCCCCC"]
    
    return: 3


## 题目大意

现在给定3个参数 — start, end, bank，分别代表起始基因序列，目标基因序列及基因库，请找出能够使起始基因序列变化为目标基因序列所需的最少变化次数。如果无法实现目标变化，请返回 -1。

注意:

1. 起始基因序列默认是合法的，但是它并不一定会出现在基因库中。
2. 所有的目标基因序列必须是合法的。
3. 假定起始基因序列与目标基因序列是不一样的。


## 解题思路


- 给出 start 和 end 两个字符串和一个 bank 字符串数组，问从 start 字符串经过多少次最少变换能变换成 end 字符串。每次变换必须使用 bank 字符串数组中的值。
- 这一题完全就是第 127 题的翻版题，解题思路和代码 99% 是一样的。相似的题目也包括第 126 题。这一题比他们都要简单。有 2 种解法，BFS 和 DFS。具体思路可以见第 127 题的题解。



方法一：广度优先搜索
思路与算法

经过分析可知，题目要求将一个基因序列 
�
A 变化至另一个基因序列 
�
B，需要满足一下条件：

序列 
�
A 与 序列 
�
B 之间只有一个字符不同；
变化字符只能从 
‘A’,
 
‘C’,
 
‘G’,
 
‘T’
‘A’, ‘C’, ‘G’, ‘T’ 中进行选择；
变换后的序列 
�
B 一定要在字符串数组 
bank
bank 中。
根据以上变换规则，我们可以进行尝试所有合法的基因变化，并找到最小的变换次数即可。步骤如下：

如果 
start
start 与 
end
end 相等，此时直接返回 
0
0；如果最终的基因序列不在 
bank
bank 中，则此时按照题意要求，无法生成，直接返回 
−
1
−1；

首先我们将可能变换的基因 
�
s 从队列中取出，按照上述的变换规则，尝试所有可能的变化后的基因，比如一个 
AACCGGTA
AACCGGTA，我们依次尝试改变基因 
�
s 的一个字符，并尝试所有可能的基因变化序列 
�
0
,
�
1
,
�
2
,
⋯
 
,
�
�
,
⋯
 
,
�
23
s 
0
​
 ,s 
1
​
 ,s 
2
​
 ,⋯,s 
i
​
 ,⋯,s 
23
​
 ，变化一次最多可能会生成 
3
×
8
=
24
3×8=24 种不同的基因序列。

我们需要检测当前生成的基因序列的合法性 
�
�
s 
i
​
 ，首先利用哈希表检测 
�
�
s 
i
​
  是否在数组 
bank
bank 中，如果是则认为该基因合法，否则改变化非法直接丢弃；其次我们还需要用哈希表记录已经遍历过的基因序列，如果该基因序列已经遍历过，则此时直接跳过；如果合法且未遍历过的基因序列，则我们将其加入到队列中。

如果当前变换后的基因序列与 
end
end 相等，则此时我们直接返回最小的变化次数即可；如果队列中所有的元素都已经遍历完成还无法变成 
end
end，则此时无法实现目标变化，返回 
−
1
−1。

代码

Python3C++JavaC#CJavaScriptGolang

func minMutation(start, end string, bank []string) int {
    if start == end {
        return 0
    }
    bankSet := map[string]struct{}{}
    for _, s := range bank {
        bankSet[s] = struct{}{}
    }
    if _, ok := bankSet[end]; !ok {
        return -1
    }

    q := []string{start}
    for step := 0; q != nil; step++ {
        tmp := q
        q = nil
        for _, cur := range tmp {
            for i, x := range cur {
                for _, y := range "ACGT" {
                    if y != x {
                        nxt := cur[:i] + string(y) + cur[i+1:]
                        if _, ok := bankSet[nxt]; ok {
                            if nxt == end {
                                return step + 1
                            }
                            delete(bankSet, nxt)
                            q = append(q, nxt)
                        }
                    }
                }
            }
        }
    }
    return -1
}
复杂度分析

时间复杂度：
�
(
�
×
�
×
�
)
O(C×n×m)，其中 
�
n 为基因序列的长度，
�
m 为数组 
bank
bank 的长度。对于队列中的每个合法的基因序列每次都需要计算 
�
×
�
C×n 种变化，在这里 
�
=
4
C=4；队列中最多有 
�
m 个元素，因此时间复杂度为 
�
(
�
×
�
×
�
)
O(C×n×m)。

空间复杂度：
�
(
�
×
�
)
O(n×m)，其中 
�
n 为基因序列的长度，
�
m 为数组 
bank
bank 的长度。合法性的哈希表中一共存有 
�
m 个元素，队列中最多有 
�
m 个元素，每个元素的空间为 
�
(
�
)
O(n)；队列中最多有 
�
m 个元素，每个元素的空间为 
�
(
�
)
O(n)，因此空间复杂度为 
�
(
�
×
�
)
O(n×m)。

方法二：预处理优化
思路与算法

经过分析可知，题目要求将一个基因序列 
�
A 变化至另一个基因序列 
�
B，需要满足一下条件：

序列 
�
A 与 序列 
�
B 之间只有一个字符不同；
变化字符只能从 
‘A’,
 
‘C’,
 
‘G’,
 
‘T’
‘A’, ‘C’, ‘G’, ‘T’ 中进行选择；
变换后的序列 
�
B 一定要在字符串数组 
bank
bank 中。
已知方法一中广度优先搜索方法，我们可以对 
bank
bank 进行预处理，只在合法的基因变化进行搜索即可。由于题目中给定的 
bank
bank 基因库的长度较小，因此可以直接在对 
bank
bank 进行预处理，找到基因库中的每个基因的合法变换，而不需要像方法一中每次都需要去计算基因的变化序列，我们将每个基因的合法变化关系存储在邻接表 
adj
adj 中，每次基因变化搜索只在 
adj
adj 中进行即可。

Python3C++JavaC#CJavaScriptGolang

func diffOne(s, t string) (diff bool) {
    for i := range s {
        if s[i] != t[i] {
            if diff {
                return false
            }
            diff = true
        }
    }
    return
}

func minMutation(start, end string, bank []string) int {
    if start == end {
        return 0
    }

    m := len(bank)
    adj := make([][]int, m)
    endIndex := -1
    for i, s := range bank {
        if s == end {
            endIndex = i
        }
        for j := i + 1; j < m; j++ {
            if diffOne(s, bank[j]) {
                adj[i] = append(adj[i], j)
                adj[j] = append(adj[j], i)
            }
        }
    }
    if endIndex == -1 {
        return -1
    }

    var q []int
    vis := make([]bool, m)
    for i, s := range bank {
        if diffOne(start, s) {
            q = append(q, i)
            vis[i] = true
        }
    }
    for step := 1; q != nil; step++ {
        tmp := q
        q = nil
        for _, cur := range tmp {
            if cur == endIndex {
                return step
            }
            for _, nxt := range adj[cur] {
                if !vis[nxt] {
                    vis[nxt] = true
                    q = append(q, nxt)
                }
            }
        }
    }
    return -1
}
复杂度分析

时间复杂度：
�
(
�
×
�
2
)
O(m×n 
2
 )，其中 
�
m 为基因序列的长度，
�
n 为数组 
bank
bank 的长度。计算合法的基因变化 
adj
adj 需要的时间为 
�
(
�
×
�
2
)
O(m×n 
2
 )，广度优先搜索时，队列中最多有 
�
n 个元素，需要的时间为 
�
(
�
)
O(n)，因此时间复杂度为 
�
(
�
×
�
2
)
O(m×n 
2
 )。

空间复杂度：
�
(
�
2
)
O(n 
2
 )，其中 
�
n 为数组 
bank
bank 的长度。计算合法的基因变化 
adj
adj 需要的空间为 
�
(
�
2
)
O(n 
2
 )，队列中最多有 
�
n 个元素，因此空间复杂度为 
�
(
�
2
)
O(n 
2
 )。

作者：LeetCode-Solution
链接：https://leetcode.cn/problems/minimum-genetic-mutation/solution/zui-xiao-ji-yin-bian-hua-by-leetcode-sol-lhwy/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。