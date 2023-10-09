# [127. Word Ladder](https://leetcode.com/problems/word-ladder/)


## 题目

Given two words (*beginWord* and *endWord*), and a dictionary's word list, find the length of shortest transformation sequence from *beginWord* to *endWord*, such that:

1. Only one letter can be changed at a time.
2. Each transformed word must exist in the word list. Note that *beginWord* is *not* a transformed word.

**Note:**

- Return 0 if there is no such transformation sequence.
- All words have the same length.
- All words contain only lowercase alphabetic characters.
- You may assume no duplicates in the word list.
- You may assume *beginWord* and *endWord* are non-empty and are not the same.

**Example 1:**

    Input:
    beginWord = "hit",
    endWord = "cog",
    wordList = ["hot","dot","dog","lot","log","cog"]
    
    Output: 5
    
    Explanation: As one shortest transformation is "hit" -> "hot" -> "dot" -> "dog" -> "cog",
    return its length 5.

**Example 2:**

    Input:
    beginWord = "hit"
    endWord = "cog"
    wordList = ["hot","dot","dog","lot","log"]
    
    Output: 0
    
    Explanation: The endWord "cog" is not in wordList, therefore no possible transformation.


## 题目大意

给定两个单词（beginWord 和 endWord）和一个字典，找到从 beginWord 到 endWord 的最短转换序列的长度。转换需遵循如下规则：

1. 每次转换只能改变一个字母。
2. 转换过程中的中间单词必须是字典中的单词。

说明:

- 如果不存在这样的转换序列，返回 0。
- 所有单词具有相同的长度。
- 所有单词只由小写字母组成。
- 字典中不存在重复的单词。
- 你可以假设 beginWord 和 endWord 是非空的，且二者不相同。


## 解题思路

- 这一题要求输出从 `beginWord` 变换到 `endWord` 最短变换次数。可以用 BFS，从 `beginWord` 开始变换，把该单词的每个字母都用 `'a'~'z'` 变换一次，生成的数组到 `wordList` 中查找，这里用 Map 来记录查找。找得到就入队列，找不到就输出 0 。入队以后按照 BFS 的算法依次遍历完，当所有单词都 `len(queue)<=0` 出队以后，整个程序结束。
- 这一题题目中虽然说了要求找到一条最短的路径，但是实际上最短的路径的寻找方法已经告诉你了：
	1. 每次只变换一个字母
	2. 每次变换都必须在 `wordList` 中  
所以不需要单独考虑何种方式是最短的。 







基本分析
根据题意，每次只能替换一个字符，且每次产生的新单词必须在 wordList 出现过。

一个朴素的实现方法是，使用 BFS 的方式求解。

从 beginWord 出发，枚举所有替换一个字符的方案，如果方案存在于 wordList 中，则加入队列中，这样队列中就存在所有替换次数为 
1
1 的单词。然后从队列中取出元素，继续这个过程，直到遇到 endWord 或者队列为空为止。

同时为了「防止重复枚举到某个中间结果」和「记录每个中间结果是经过多少次转换而来」，我们需要建立一个「哈希表」进行记录。

哈希表的 KV 形式为 { 单词 : 由多少次转换得到 }。

当枚举到新单词 str 时，需要先检查是否已经存在与「哈希表」中，如果不存在则更新「哈希表」并将新单词放入队列中。

这样的做法可以确保「枚举到所有由 beginWord 到 endWord 的转换路径」，并且由 beginWord 到 endWord 的「最短转换路径」必然会最先被枚举到。

双向 BFS
经过分析，BFS 确实可以做，但本题的数据范围较大：1 <= beginWord.length <= 10

朴素的 BFS 可能会带来「搜索空间爆炸」的情况。

想象一下，如果我们的 wordList 足够丰富（包含了所有单词），对于一个长度为 
10
10 的 beginWord​ 替换一次字符可以产生 
10
∗
25
10∗25 个新单词（每个替换点可以替换另外 
25
25 个小写字母），第一层就会产生 
250
250 个单词；第二层会产生超过 
6
∗
1
0
4
6∗10 
4
  个新单词 ...

随着层数的加深，这个数字的增速越快，这就是「搜索空间爆炸」问题。



在朴素的 BFS 实现中，空间的瓶颈主要取决于搜索空间中的最大宽度。

那么有没有办法让我们不使用这么宽的搜索空间，同时又能保证搜索到目标结果呢？

「双向 BFS」 可以很好的解决这个问题：

同时从两个方向开始搜索，一旦搜索到相同的值，意味着找到了一条联通起点和终点的最短路径。



「双向 BFS」的基本实现思路如下：

创建「两个队列」分别用于两个方向的搜索；
创建「两个哈希表」用于「解决相同节点重复搜索」和「记录转换次数」；
为了尽可能让两个搜索方向“平均”，每次从队列中取值进行扩展时，先判断哪个队列容量较少；
如果在搜索过程中「搜索到对方搜索过的节点」，说明找到了最短路径。
「双向 BFS」基本思路对应的伪代码大致如下：

Java

d1、d2 为两个方向的队列
m1、m2 为两个方向的哈希表，记录每个节点距离起点的
    
// 只有两个队列都不空，才有必要继续往下搜索
// 如果其中一个队列空了，说明从某个方向搜到底都搜不到该方向的目标节点
while(!d1.isEmpty() && !d2.isEmpty()) {
    if (d1.size() < d2.size()) {
        update(d1, m1, m2);
    } else {
        update(d2, m2, m1);
    }
}

// update 为将当前队列 d 中包含的元素取出，进行「一次完整扩展」的逻辑（按层拓展）
void update(Deque d, Map cur, Map other) {}
回到本题，我们看看如何使用「双向 BFS」进行求解。

估计不少同学是第一次接触「双向 BFS」，因此这次我写了大量注释。

建议大家带着对「双向 BFS」的基本理解去阅读。

代码：

Java

class Solution {
    String s, e;
    Set<String> set = new HashSet<>();
    public int ladderLength(String _s, String _e, List<String> ws) {
        s = _s;
        e = _e;
        // 将所有 word 存入 set，如果目标单词不在 set 中，说明无解
        for (String w : ws) set.add(w);
        if (!set.contains(e)) return 0;
        int ans = bfs();
        return ans == -1 ? 0 : ans + 1;
    }

    int bfs() {
        // d1 代表从起点 beginWord 开始搜索（正向）
        // d2 代表从结尾 endWord 开始搜索（反向）
        Deque<String> d1 = new ArrayDeque<>(), d2 = new ArrayDeque(); 
        
        /*
         * m1 和 m2 分别记录两个方向出现的单词是经过多少次转换而来
         * e.g. 
         * m1 = {"abc":1} 代表 abc 由 beginWord 替换 1 次字符而来
         * m2 = {"xyz":3} 代表 xyz 由 endWord 替换 3 次字符而来
         */
        Map<String, Integer> m1 = new HashMap<>(), m2 = new HashMap<>();
        d1.add(s);
        m1.put(s, 0);
        d2.add(e);
        m2.put(e, 0);
        
        /*
         * 只有两个队列都不空，才有必要继续往下搜索
         * 如果其中一个队列空了，说明从某个方向搜到底都搜不到该方向的目标节点
         * e.g. 
         * 例如，如果 d1 为空了，说明从 beginWord 搜索到底都搜索不到 endWord，反向搜索也没必要进行了
         */
        while (!d1.isEmpty() && !d2.isEmpty()) {
            int t = -1;
            // 为了让两个方向的搜索尽可能平均，优先拓展队列内元素少的方向
            if (d1.size() <= d2.size()) {
                t = update(d1, m1, m2);
            } else {
                t = update(d2, m2, m1);
            }
            if (t != -1) return t;
        }
        return -1;
    }

    // update 代表从 deque 中取出一个单词进行扩展，
    // cur 为当前方向的距离字典；other 为另外一个方向的距离字典
    int update(Deque<String> deque, Map<String, Integer> cur, Map<String, Integer> other) {
        int m = deque.size();
        while (m-- > 0) {
            // 获取当前需要扩展的原字符串
            String poll = deque.pollFirst();
            int n = poll.length();

            // 枚举替换原字符串的哪个字符 i
            for (int i = 0; i < n; i++) {
                // 枚举将 i 替换成哪个小写字母
                for (int j = 0; j < 26; j++) {
                    // 替换后的字符串
                    String sub = poll.substring(0, i) + String.valueOf((char)('a' + j)) + poll.substring(i + 1);
                    if (set.contains(sub)) {
                        // 如果该字符串在「当前方向」被记录过（拓展过），跳过即可
                        if (cur.containsKey(sub) && cur.get(sub) <= cur.get(poll) + 1) continue;

                        // 如果该字符串在「另一方向」出现过，说明找到了联通两个方向的最短路
                        if (other.containsKey(sub)) {
                            return cur.get(poll) + 1 + other.get(sub);
                        } else {
                            // 否则加入 deque 队列
                            deque.addLast(sub);
                            cur.put(sub, cur.get(poll) + 1);
                        }
                    }
                }
            }
        }
        return -1;
    }
}
时间复杂度：令 wordList 长度为 
�
n，beginWord 字符串长度为 
�
m。由于所有的搜索结果必须都在 wordList 出现过，因此算上起点最多有 
�
+
1
n+1 节点，最坏情况下，所有节点都联通，搜索完整张图复杂度为 
�
(
�
2
)
O(n 
2
 ) ；从 beginWord 出发进行字符替换，替换时进行逐字符检查，复杂度为 
�
(
�
)
O(m)。整体复杂度为 
�
(
�
∗
�
2
)
O(m∗n 
2
 )
空间复杂度：同等空间大小。
�
(
�
∗
�
2
)
O(m∗n 
2
 )
总结
这本质其实是一个「所有边权均为 1」最短路问题：将 beginWord 和所有在 wordList 出现过的字符串看做是一个点。每一次转换操作看作产生边权为 1 的边。问题求以 beginWord 为源点，以 endWord 为汇点的最短路径。

借助这个题，我向你介绍了「双向 BFS」，「双向 BFS」可以有效解决「搜索空间爆炸」问题。

对于那些搜索节点随着层数增加呈倍数或指数增长的搜索问题，可以使用「双向 BFS」进行求解。

【补充】启发式搜索 AStar
可以直接根据本题规则来设计 A* 的「启发式函数」。

比如对于两个字符串 a b 直接使用它们不同字符的数量来充当估值距离，我觉得是合适的。

因为不同字符数量的差值可以确保不会超过真实距离（是一个理论最小替换次数）。

注意：本题数据比较弱，用 A* 过了，但通常我们需要「确保有解」，A* 的启发搜索才会发挥真正价值。而本题，除非 endWord 本身就不在 wordList 中，其余情况我们无法很好提前判断「是否有解」。这时候 A* 将不能带来「搜索空间的优化」，效果不如「双向 BFS」。

代码：

Java

class Solution {
    class Node {
        String str;
        int val;
        Node (String _str, int _val) {
            str = _str;
            val = _val;
        }
    }
    String s, e;
    int INF = 0x3f3f3f3f;
    Set<String> set = new HashSet<>();
    public int ladderLength(String _s, String _e, List<String> ws) {
        s = _s;
        e = _e;
        for (String w : ws) set.add(w);
        if (!set.contains(e)) return 0;
        int ans = astar();
        return ans == -1 ? 0 : ans + 1;
    }
    int astar() {
        PriorityQueue<Node> q = new PriorityQueue<>((a,b)->a.val-b.val);
        Map<String, Integer> dist = new HashMap<>();
        dist.put(s, 0);
        q.add(new Node(s, f(s)));
        
        while (!q.isEmpty()) {
            Node poll = q.poll();
            String str = poll.str;
            int distance = dist.get(str);
            if (str.equals(e)) {
                break;
            }
            int n = str.length();
            for (int i = 0; i < n; i++) {
                for (int j = 0; j < 26; j++) {
                    String sub = str.substring(0, i) + String.valueOf((char)('a' + j)) + str.substring(i + 1);
                    if (!set.contains(sub)) continue;
                    if (!dist.containsKey(sub) || dist.get(sub) > distance + 1) {
                        dist.put(sub, distance + 1);
                        q.add(new Node(sub, dist.get(sub) + f(sub)));
                    }
                }
            }
        }
        return dist.containsKey(e) ? dist.get(e) : -1;
    }
    int f(String str) {
        if (str.length() != e.length()) return INF;
        int n = str.length();
        int ans = 0;
        for (int i = 0; i < n; i++) {
            ans += str.charAt(i) == e.charAt(i) ? 0 : 1;
        }
        return ans;
    }
}
最后
如果有帮助到你，请给题解点个赞和收藏，让更多的人看到 ~ ("▔□▔)/

也欢迎你 关注我 和 加入我们的「组队打卡」小群 ，提供写「证明」&「思路」的高质量题解。

所有题解已经加入 刷题指南，欢迎 star 哦 ~

作者：AC_OIer
链接：https://leetcode.cn/problems/word-ladder/solution/gong-shui-san-xie-ru-he-shi-yong-shuang-magjd/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。











方法一：广度优先搜索 + 优化建图
思路

本题要求的是最短转换序列的长度，看到最短首先想到的就是广度优先搜索。想到广度优先搜索自然而然的就能想到图，但是本题并没有直截了当的给出图的模型，因此我们需要把它抽象成图的模型。

我们可以把每个单词都抽象为一个点，如果两个单词可以只改变一个字母进行转换，那么说明他们之间有一条双向边。因此我们只需要把满足转换条件的点相连，就形成了一张图。

基于该图，我们以 beginWord 为图的起点，以 endWord 为终点进行广度优先搜索，寻找 beginWord 到 endWord 的最短路径。



算法

基于上面的思路我们考虑如何编程实现。

首先为了方便表示，我们先给每一个单词标号，即给每个单词分配一个 id。创建一个由单词 word 到 id 对应的映射 wordId，并将 beginWord 与 wordList 中所有的单词都加入这个映射中。之后我们检查 endWord 是否在该映射内，若不存在，则输入无解。我们可以使用哈希表实现上面的映射关系。

然后我们需要建图，依据朴素的思路，我们可以枚举每一对单词的组合，判断它们是否恰好相差一个字符，以判断这两个单词对应的节点是否能够相连。但是这样效率太低，我们可以优化建图。

具体地，我们可以创建虚拟节点。对于单词 hit，我们创建三个虚拟节点 *it、h*t、hi*，并让 hit 向这三个虚拟节点分别连一条边即可。如果一个单词能够转化为 hit，那么该单词必然会连接到这三个虚拟节点之一。对于每一个单词，我们枚举它连接到的虚拟节点，把该单词对应的 id 与这些虚拟节点对应的 id 相连即可。

最后我们将起点加入队列开始广度优先搜索，当搜索到终点时，我们就找到了最短路径的长度。注意因为添加了虚拟节点，所以我们得到的距离为实际最短路径长度的两倍。同时我们并未计算起点对答案的贡献，所以我们应当返回距离的一半再加一的结果。

代码

C++JavaPython3GolangC

class Solution {
public:
    unordered_map<string, int> wordId;
    vector<vector<int>> edge;
    int nodeNum = 0;

    void addWord(string& word) {
        if (!wordId.count(word)) {
            wordId[word] = nodeNum++;
            edge.emplace_back();
        }
    }

    void addEdge(string& word) {
        addWord(word);
        int id1 = wordId[word];
        for (char& it : word) {
            char tmp = it;
            it = '*';
            addWord(word);
            int id2 = wordId[word];
            edge[id1].push_back(id2);
            edge[id2].push_back(id1);
            it = tmp;
        }
    }

    int ladderLength(string beginWord, string endWord, vector<string>& wordList) {
        for (string& word : wordList) {
            addEdge(word);
        }
        addEdge(beginWord);
        if (!wordId.count(endWord)) {
            return 0;
        }
        vector<int> dis(nodeNum, INT_MAX);
        int beginId = wordId[beginWord], endId = wordId[endWord];
        dis[beginId] = 0;

        queue<int> que;
        que.push(beginId);
        while (!que.empty()) {
            int x = que.front();
            que.pop();
            if (x == endId) {
                return dis[endId] / 2 + 1;
            }
            for (int& it : edge[x]) {
                if (dis[it] == INT_MAX) {
                    dis[it] = dis[x] + 1;
                    que.push(it);
                }
            }
        }
        return 0;
    }
};
复杂度分析

时间复杂度：
�
(
�
×
�
2
)
O(N×C 
2
 )。其中 
�
N 为 wordList 的长度，
�
C 为列表中单词的长度。

建图过程中，对于每一个单词，我们需要枚举它连接到的所有虚拟节点，时间复杂度为 
�
(
�
)
O(C)，将这些单词加入到哈希表中，时间复杂度为 
�
(
�
×
�
)
O(N×C)，因此总时间复杂度为 
�
(
�
×
�
)
O(N×C)。

广度优先搜索的时间复杂度最坏情况下是 
�
(
�
×
�
)
O(N×C)。每一个单词需要拓展出 
�
(
�
)
O(C) 个虚拟节点，因此节点数 
�
(
�
×
�
)
O(N×C)。

空间复杂度：
�
(
�
×
�
2
)
O(N×C 
2
 )。其中 
�
N 为 wordList 的长度，
�
C 为列表中单词的长度。哈希表中包含 
�
(
�
×
�
)
O(N×C) 个节点，每个节点占用空间 
�
(
�
)
O(C)，因此总的空间复杂度为 
�
(
�
×
�
2
)
O(N×C 
2
 )。

方法二：双向广度优先搜索
思路及解法

根据给定字典构造的图可能会很大，而广度优先搜索的搜索空间大小依赖于每层节点的分支数量。假如每个节点的分支数量相同，搜索空间会随着层数的增长指数级的增加。考虑一个简单的二叉树，每一层都是满二叉树的扩展，节点的数量会以 
2
2 为底数呈指数增长。

如果使用两个同时进行的广搜可以有效地减少搜索空间。一边从 beginWord 开始，另一边从 endWord 开始。我们每次从两边各扩展一层节点，当发现某一时刻两边都访问过同一顶点时就停止搜索。这就是双向广度优先搜索，它可以可观地减少搜索空间大小，从而提高代码运行效率。



代码

C++JavaPython3GolangC

class Solution {
public:
    unordered_map<string, int> wordId;
    vector<vector<int>> edge;
    int nodeNum = 0;

    void addWord(string& word) {
        if (!wordId.count(word)) {
            wordId[word] = nodeNum++;
            edge.emplace_back();
        }
    }

    void addEdge(string& word) {
        addWord(word);
        int id1 = wordId[word];
        for (char& it : word) {
            char tmp = it;
            it = '*';
            addWord(word);
            int id2 = wordId[word];
            edge[id1].push_back(id2);
            edge[id2].push_back(id1);
            it = tmp;
        }
    }

    int ladderLength(string beginWord, string endWord, vector<string>& wordList) {
        for (string& word : wordList) {
            addEdge(word);
        }
        addEdge(beginWord);
        if (!wordId.count(endWord)) {
            return 0;
        }

        vector<int> disBegin(nodeNum, INT_MAX);
        int beginId = wordId[beginWord];
        disBegin[beginId] = 0;
        queue<int> queBegin;
        queBegin.push(beginId);

        vector<int> disEnd(nodeNum, INT_MAX);
        int endId = wordId[endWord];
        disEnd[endId] = 0;
        queue<int> queEnd;
        queEnd.push(endId);

        while (!queBegin.empty() && !queEnd.empty()) {
            int queBeginSize = queBegin.size();
            for (int i = 0; i < queBeginSize; ++i) {
                int nodeBegin = queBegin.front();
                queBegin.pop();
                if (disEnd[nodeBegin] != INT_MAX) {
                    return (disBegin[nodeBegin] + disEnd[nodeBegin]) / 2 + 1;
                }
                for (int& it : edge[nodeBegin]) {
                    if (disBegin[it] == INT_MAX) {
                        disBegin[it] = disBegin[nodeBegin] + 1;
                        queBegin.push(it);
                    }
                }
            }

            int queEndSize = queEnd.size();
            for (int i = 0; i < queEndSize; ++i) {
                int nodeEnd = queEnd.front();
                queEnd.pop();
                if (disBegin[nodeEnd] != INT_MAX) {
                    return (disBegin[nodeEnd] + disEnd[nodeEnd]) / 2 + 1;
                }
                for (int& it : edge[nodeEnd]) {
                    if (disEnd[it] == INT_MAX) {
                        disEnd[it] = disEnd[nodeEnd] + 1;
                        queEnd.push(it);
                    }
                }
            }
        }
        return 0;
    }
};
复杂度分析

时间复杂度：
�
(
�
×
�
2
)
O(N×C 
2
 )。其中 
�
N 为 wordList 的长度，
�
C 为列表中单词的长度。

建图过程中，对于每一个单词，我们需要枚举它连接到的所有虚拟节点，时间复杂度为 
�
(
�
)
O(C)，将这些单词加入到哈希表中，时间复杂度为 
�
(
�
×
�
)
O(N×C)，因此总时间复杂度为 
�
(
�
×
�
)
O(N×C)。

双向广度优先搜索的时间复杂度最坏情况下是 
�
(
�
×
�
)
O(N×C)。每一个单词需要拓展出 
�
(
�
)
O(C) 个虚拟节点，因此节点数 
�
(
�
×
�
)
O(N×C)。

空间复杂度：
�
(
�
×
�
2
)
O(N×C 
2
 )。其中 
�
N 为 wordList 的长度，
�
C 为列表中单词的长度。哈希表中包含 
�
(
�
×
�
)
O(N×C) 个节点，每个节点占用空间 
�
(
�
)
O(C)，因此总的空间复杂度为 
�
(
�
×
�
2
)
O(N×C 
2
 )。

下一篇：广度优先遍历、双向广度优先遍历（Java）
© 著作权归作者所有
184
条评论

最热

编辑
预览







评论
精选评论(5)

alexSnap
来自广东
2020-11-05
“程乙己，你又没有思路了！”他不回答，只是将题解ctrl+c再ctrl+v到编辑器。他们又故意的高声嚷道，“你一定又写不出代码了了！”程乙己睁大眼睛说，“你怎么这样凭空污人清白……”“什么清白？我前天亲眼见你copy了别人的整段代码，吊着打。”程乙己便涨红了脸，额上的青筋条条绽出，争辩道，“copy代码不能算copy……抄！……程序员的事，能算抄么？”接连便是难懂的话，什么“拓展思路”，什么“借鉴”之类，引得众人都哄笑起来：评论区内外充满了快活的空气。


312

查看 17 条回复

回复

杜琦拓
L3
来自北京
2020-11-05
额，这中等不比昨天的困难恶心？


247

查看 15 条回复

回复

古今皆付笑谈中
L1
来自上海
2020-11-05
昨天的困难题把我抬到了不属于自己的高度


167

查看 4 条回复

回复

kat
L1
来自上海
2020-11-05
对于单词马冬梅，我们创建三个虚拟节点：马*梅、*冬梅、马冬*


141

查看 3 条回复

回复

Deepdoo
来自上海
2022-01-12
每次看官方题解都是直接看代码，中文写的比代码还难懂


75

回复
评论(184)

Christina
L3
来自重庆
2023-05-27
数据结构用对， 普通的 bfs 也能通过


class Solution {
private:
    const string str = "abcdefghijklmnopqrstuvwxyz";

public:
    int ladderLength(string beginWord, string endWord, vector<string>& wordList) {
        
        unordered_set<string> wordSet(wordList.begin(), wordList.end());
        if (wordSet.find(endWord) == wordSet.end())
            return 0;

        unordered_map<string, bool> m;

        for (auto e : wordList)
            m.insert(make_pair(e, false));

        queue<string> q;
        q.push(beginWord);
        int step = 0;
        


        while (!q.empty()){
            int size = q.size();
            step++;

            for (int i = 0; i < size; i++){
                string s1 = q.front();
                // if (s1 == endWord)
                //     return step + 1;
                q.pop();


                for (int i = 0; i < s1.size(); i++){
                    string s2 = s1;
                    for (auto e : str){
                        s2[i] = e;
                        if (s2 == endWord)
                            return step + 1;
                        if (wordSet.find(s2) != wordSet.end() && m[s2] == false){
                            m[s2] = true;
                            q.push(s2);
                        }
                    }
                }

            }
            
        }

        return 0;
    }
};


4

查看 1 条回复

回复

薛钉恶的锚
L1
来自陕西
2020-11-05
这叫中等题？是不是该跟昨天的困难题换换位置了兄贵？


73

查看 11 条回复

回复

BoredBean
L1
来自河北
2020-11-05
看完题目直接来找答案


70

回复

分享

举报


作者：LeetCode-Solution
链接：https://leetcode.cn/problems/word-ladder/solution/dan-ci-jie-long-by-leetcode-solution/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。