# [207. Course Schedule](https://leetcode.com/problems/course-schedule/)

## 题目

There are a total of n courses you have to take, labeled from `0` to `n-1`.

Some courses may have prerequisites, for example to take course 0 you have to first take course 1, which is expressed as a pair: `[0,1]`

Given the total number of courses and a list of prerequisite **pairs**, is it possible for you to finish all courses?

**Example 1:**

    Input: 2, [[1,0]] 
    Output: true
    Explanation: There are a total of 2 courses to take. 
                 To take course 1 you should have finished course 0. So it is possible.

**Example 2:**

    Input: 2, [[1,0],[0,1]]
    Output: false
    Explanation: There are a total of 2 courses to take. 
                 To take course 1 you should have finished course 0, and to take course 0 you should
                 also have finished course 1. So it is impossible.

**Note:**

1. The input prerequisites is a graph represented by **a list of edges**, not adjacency matrices. Read more about [how a graph is represented](https://www.khanacademy.org/computing/computer-science/algorithms/graph-representation/a/representing-graphs).
2. You may assume that there are no duplicate edges in the input prerequisites.


## 题目大意

现在你总共有 n 门课需要选，记为 0 到 n-1。在选修某些课程之前需要一些先修课程。 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们: [0,1]。给定课程总量以及它们的先决条件，判断是否可能完成所有课程的学习？



## 解题思路

- 给出 n 个任务，每两个任务之间有相互依赖关系，比如 A 任务一定要在 B 任务之前完成才行。问是否可以完成所有任务。
- 这一题就是标准的 AOV 网的拓扑排序问题。拓扑排序问题的解决办法是主要是循环执行以下两步，直到不存在入度为0的顶点为止。
    - 1. 选择一个入度为0的顶点并输出之；
    - 2. 从网中删除此顶点及所有出边。

    循环结束后，若输出的顶点数小于网中的顶点数，则输出“有回路”信息，即无法完成所有任务；否则输出的顶点序列就是一种拓扑序列，即可以完成所有任务。









logo
学习
new
题库
竞赛
讨论
求职
商店

体验新版
Plus 会员
力扣面试
0

展示方式
题目描述
评论 (1.2k)
题解 (2.2k)
提交记录
搜索题解
314406
11:15
课程表
5665
41:08
LeetCode 每日一题 Daily Challenge 207. 课程表 （go）
11:15

课程表
📺 视频题解 📖 文字题解 前言 本题和 210. 课程表 II 是几乎一样的题目。如果在过去完成过该题，那么只要将代码中的返回值从「非空数组 / 空数组」修改成「\text{True} / \text{False}」就可以通过本题。 本题是一道经典的「拓扑排序」问题。 给定一个包含 n 个节点的有向图 G，我们

「图解」拓扑排序 | 课程表问题
题意解释 一共有 n 门课要上，编号为 0 ~ n-1。 先决条件 [1, 0]，意思是必须先上课 0，才能上课 1。 给你 n、和一个先决条件表，请你判断能否完成所有课程。 再举个生活的例子 先穿内裤再穿裤子，先穿打底再穿外套，先穿衣服再戴帽子，是约定俗成的。 内裤外穿、光着身子戴帽子等，都会有点奇怪。 我

课程表（拓扑排序：入度表BFS法 / DFS法，清晰图解）
解题思路： 本题可约化为： 课程安排图是否是 有向无环图(DAG)。即课程间规定了前置条件，但不能构成任何环路，否则课程前置条件将不成立。 思路是通过 拓扑排序 判断此课程安排图是否是 有向无环图(DAG) 。 拓扑排序原理： 对 DAG 的顶点进行排序，使得对每一条有向边 (u, v)，均有 u（在排序记录中）比

拓扑排序、深度优先遍历
这道题的做法同样适用于第 210 题。 方法一：拓扑排序（Kahn 算法，其实就是广度优先遍历的思路） 1 / 13 拓扑排序实际上应用的是贪心算法。贪心算法简而言之：每一步最优，全局就最优。 具体到拓扑排序，每一次都从图中删除没有前驱的顶点，这里并不需要真正的做删除操作，我们可以设置一个入度数组，每一轮都输出入度为

拓扑排序 时间 12 ms 击败 98.60% 内存 12.9 MB 击败 93.7%
[本题解无摘要]

多次遍历法
[本题解无摘要]

python3 bfs
dfs的思路有点不好理解，个人认为这样的题目用bfs的话可读性会比较高。

ts+拓扑排序+bfs+详细注释
[本题解无摘要]

【宫水三叶】拓扑排序运用题
拓扑排序 为了方便，我们记 numCourses 为 n，prerequisites 为 g。 若课程 a 存在前置课程 b 的话，我们添加一条从 b 到 a 的有向边，同时统计所有点的入度。 当处理完所有的 g[i] 后，将所有的入度为 0 的课程（含义为没有前置课程要求的科目）进行入队操作，跑一遍「拓扑排序」，若

双百，运气算法
[该用户太懒了，只写了 5 行代码]


1

2

3

4

5


218

207/3151

智能模式

模拟面试






123
func canFinish(numCourses int, prerequisites [][]int) bool {

}
控制台 
贡献
1 / 2183

课程表
力扣官方题解
314.4k
发布于 2020-08-03
未知归属地
📺 视频题解

📖 文字题解
前言
本题和 210. 课程表 II 是几乎一样的题目。如果在过去完成过该题，那么只要将代码中的返回值从「非空数组 / 空数组」修改成「
True
True / 
False
False」就可以通过本题。

本题是一道经典的「拓扑排序」问题。

给定一个包含 
�
n 个节点的有向图 
�
G，我们给出它的节点编号的一种排列，如果满足：

对于图 
�
G 中的任意一条有向边 
(
�
,
�
)
(u,v)，
�
u 在排列中都出现在 
�
v 的前面。

那么称该排列是图 
�
G 的「拓扑排序」。根据上述的定义，我们可以得出两个结论：

如果图 
�
G 中存在环（即图 
�
G 不是「有向无环图」），那么图 
�
G 不存在拓扑排序。这是因为假设图中存在环 
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
�
1
x 
1
​
 ,x 
2
​
 ,⋯,x 
n
​
 ,x 
1
​
 ，那么 
�
1
x 
1
​
  在排列中必须出现在 
�
�
x 
n
​
  的前面，但 
�
�
x 
n
​
  同时也必须出现在 
�
1
x 
1
​
  的前面，因此不存在一个满足要求的排列，也就不存在拓扑排序；

如果图 
�
G 是有向无环图，那么它的拓扑排序可能不止一种。举一个最极端的例子，如果图 
�
G 值包含 
�
n 个节点却没有任何边，那么任意一种编号的排列都可以作为拓扑排序。

有了上述的简单分析，我们就可以将本题建模成一个求拓扑排序的问题了：

我们将每一门课看成一个节点；

如果想要学习课程 
�
A 之前必须完成课程 
�
B，那么我们从 
�
B 到 
�
A 连接一条有向边。这样以来，在拓扑排序中，
�
B 一定出现在 
�
A 的前面。

求出该图是否存在拓扑排序，就可以判断是否有一种符合要求的课程学习顺序。事实上，由于求出一种拓扑排序方法的最优时间复杂度为 
�
(
�
+
�
)
O(n+m)，其中 
�
n 和 
�
m 分别是有向图 
�
G 的节点数和边数，方法见 210. 课程表 II 的官方题解。而判断图 
�
G 是否存在拓扑排序，至少也要对其进行一次完整的遍历，时间复杂度也为 
�
(
�
+
�
)
O(n+m)。因此不可能存在一种仅判断图是否存在拓扑排序的方法，它的时间复杂度在渐进意义上严格优于 
�
(
�
+
�
)
O(n+m)。这样一来，我们使用和 210. 课程表 II 完全相同的方法，但无需使用数据结构记录实际的拓扑排序。为了叙述的完整性，下面的两种方法与 210. 课程表 II 的官方题解 完全相同，但在「算法」部分后的「优化」部分说明了如何省去对应的数据结构。

方法一：深度优先搜索
思路

我们可以将深度优先搜索的流程与拓扑排序的求解联系起来，用一个栈来存储所有已经搜索完成的节点。

对于一个节点 
�
u，如果它的所有相邻节点都已经搜索完成，那么在搜索回溯到 
�
u 的时候，
�
u 本身也会变成一个已经搜索完成的节点。这里的「相邻节点」指的是从 
�
u 出发通过一条有向边可以到达的所有节点。

假设我们当前搜索到了节点 
�
u，如果它的所有相邻节点都已经搜索完成，那么这些节点都已经在栈中了，此时我们就可以把 
�
u 入栈。可以发现，如果我们从栈顶往栈底的顺序看，由于 
�
u 处于栈顶的位置，那么 
�
u 出现在所有 
�
u 的相邻节点的前面。因此对于 
�
u 这个节点而言，它是满足拓扑排序的要求的。

这样以来，我们对图进行一遍深度优先搜索。当每个节点进行回溯的时候，我们把该节点放入栈中。最终从栈顶到栈底的序列就是一种拓扑排序。

算法

对于图中的任意一个节点，它在搜索的过程中有三种状态，即：

「未搜索」：我们还没有搜索到这个节点；

「搜索中」：我们搜索过这个节点，但还没有回溯到该节点，即该节点还没有入栈，还有相邻的节点没有搜索完成）；

「已完成」：我们搜索过并且回溯过这个节点，即该节点已经入栈，并且所有该节点的相邻节点都出现在栈的更底部的位置，满足拓扑排序的要求。

通过上述的三种状态，我们就可以给出使用深度优先搜索得到拓扑排序的算法流程，在每一轮的搜索搜索开始时，我们任取一个「未搜索」的节点开始进行深度优先搜索。

我们将当前搜索的节点 
�
u 标记为「搜索中」，遍历该节点的每一个相邻节点 
�
v：

如果 
�
v 为「未搜索」，那么我们开始搜索 
�
v，待搜索完成回溯到 
�
u；

如果 
�
v 为「搜索中」，那么我们就找到了图中的一个环，因此是不存在拓扑排序的；

如果 
�
v 为「已完成」，那么说明 
�
v 已经在栈中了，而 
�
u 还不在栈中，因此 
�
u 无论何时入栈都不会影响到 
(
�
,
�
)
(u,v) 之前的拓扑关系，以及不用进行任何操作。

当 
�
u 的所有相邻节点都为「已完成」时，我们将 
�
u 放入栈中，并将其标记为「已完成」。

在整个深度优先搜索的过程结束后，如果我们没有找到图中的环，那么栈中存储这所有的 
�
n 个节点，从栈顶到栈底的顺序即为一种拓扑排序。

下面的幻灯片给出了深度优先搜索的可视化流程。图中的「白色」「黄色」「绿色」节点分别表示「未搜索」「搜索中」「已完成」的状态。



优化

由于我们只需要判断是否存在一种拓扑排序，而栈的作用仅仅是存放最终的拓扑排序结果，因此我们可以只记录每个节点的状态，而省去对应的栈。

代码


func canFinish(numCourses int, prerequisites [][]int) bool {
    var (
        edges = make([][]int, numCourses)
        visited = make([]int, numCourses)
        result []int
        valid = true
        dfs func(u int)
    )

    dfs = func(u int) {
        visited[u] = 1
        for _, v := range edges[u] {
            if visited[v] == 0 {
                dfs(v)
                if !valid {
                    return
                }
            } else if visited[v] == 1 {
                valid = false
                return
            }
        }
        visited[u] = 2
        result = append(result, u)
    }

    for _, info := range prerequisites {
        edges[info[1]] = append(edges[info[1]], info[0])
    }

    for i := 0; i < numCourses && valid; i++ {
        if visited[i] == 0 {
            dfs(i)
        }
    }
    return valid
}
复杂度分析

时间复杂度: 
�
(
�
+
�
)
O(n+m)，其中 
�
n 为课程数，
�
m 为先修课程的要求数。这其实就是对图进行深度优先搜索的时间复杂度。

空间复杂度: 
�
(
�
+
�
)
O(n+m)。题目中是以列表形式给出的先修课程关系，为了对图进行深度优先搜索，我们需要存储成邻接表的形式，空间复杂度为 
�
(
�
+
�
)
O(n+m)。在深度优先搜索的过程中，我们需要最多 
�
(
�
)
O(n) 的栈空间（递归）进行深度优先搜索，因此总空间复杂度为 
�
(
�
+
�
)
O(n+m)。

方法二: 广度优先搜索
思路

方法一的深度优先搜索是一种「逆向思维」：最先被放入栈中的节点是在拓扑排序中最后面的节点。我们也可以使用正向思维，顺序地生成拓扑排序，这种方法也更加直观。

我们考虑拓扑排序中最前面的节点，该节点一定不会有任何入边，也就是它没有任何的先修课程要求。当我们将一个节点加入答案中后，我们就可以移除它的所有出边，代表着它的相邻节点少了一门先修课程的要求。如果某个相邻节点变成了「没有任何入边的节点」，那么就代表着这门课可以开始学习了。按照这样的流程，我们不断地将没有入边的节点加入答案，直到答案中包含所有的节点（得到了一种拓扑排序）或者不存在没有入边的节点（图中包含环）。

上面的想法类似于广度优先搜索，因此我们可以将广度优先搜索的流程与拓扑排序的求解联系起来。

算法

我们使用一个队列来进行广度优先搜索。初始时，所有入度为 
0
0 的节点都被放入队列中，它们就是可以作为拓扑排序最前面的节点，并且它们之间的相对顺序是无关紧要的。

在广度优先搜索的每一步中，我们取出队首的节点 
�
u：

我们将 
�
u 放入答案中；

我们移除 
�
u 的所有出边，也就是将 
�
u 的所有相邻节点的入度减少 
1
1。如果某个相邻节点 
�
v 的入度变为 
0
0，那么我们就将 
�
v 放入队列中。

在广度优先搜索的过程结束后。如果答案中包含了这 
�
n 个节点，那么我们就找到了一种拓扑排序，否则说明图中存在环，也就不存在拓扑排序了。

下面的幻灯片给出了广度优先搜索的可视化流程。



优化

由于我们只需要判断是否存在一种拓扑排序，因此我们省去存放答案数组，而是只用一个变量记录被放入答案数组的节点个数。在广度优先搜索结束之后，我们判断该变量的值是否等于课程数，就能知道是否存在一种拓扑排序。

代码


func canFinish(numCourses int, prerequisites [][]int) bool {
    var (
        edges = make([][]int, numCourses)
        indeg = make([]int, numCourses)
        result []int
    )

    for _, info := range prerequisites {
        edges[info[1]] = append(edges[info[1]], info[0])
        indeg[info[0]]++
    }

    q := []int{}
    for i := 0; i < numCourses; i++ {
        if indeg[i] == 0 {
            q = append(q, i)
        }
    }

    for len(q) > 0 {
        u := q[0]
        q = q[1:]
        result = append(result, u)
        for _, v := range edges[u] {
            indeg[v]--
            if indeg[v] == 0 {
                q = append(q, v)
            }
        }
    }
    return len(result) == numCourses
}
复杂度分析

时间复杂度: 
�
(
�
+
�
)
O(n+m)，其中 
�
n 为课程数，
�
m 为先修课程的要求数。这其实就是对图进行广度优先搜索的时间复杂度。

空间复杂度: 
�
(
�
+
�
)
O(n+m)。题目中是以列表形式给出的先修课程关系，为了对图进行广度优先搜索，我们需要存储成邻接表的形式，空间复杂度为 
�
(
�
+
�
)
O(n+m)。在广度优先搜索的过程中，我们需要最多 
�
(
�
)
O(n) 的队列空间（迭代）进行广度优先搜索。因此总空间复杂度为 
�
(
�
+
�
)
O(n+m)。

下一篇：「图解」拓扑排序 | 课程表问题
© 著作权归作者所有
189
条评论
精选评论(5)
UBGaMgsayJ
DeepFeeling
来自江西
2020-08-04
本来想做完再睡的，看了两分钟，我好像不太行，打开题解看了看，算了，睡觉

LinusTorvalds
Rookie

来自上海
2020-08-04
求大家别再"预言"每日一题了，还评论区一片纯粹。

nrocky
Rocky
来自陕西
2022-02-28
拓扑排序不用区分什么广度优先深度优先把自己弄乱了，抓住节点入度和出度的本质特征。 方法一： 从入度思考(从前往后排序)， 入度为0的节点在拓扑排序中一定排在前面, 然后删除和该节点对应的边, 迭代寻找入度为0的节点。 方法二： 从出度思考(从后往前排序)， 出度为0的节点在拓扑排序中一定排在后面, 然后删除和该节点对应的边, 迭代寻找出度为0的节点。

KamenRider
小c能不能不做舔狗了

来自辽宁
2020-08-04
就是判断一个图中是否存在环吧

yu-niang-niang
钰娘娘丿-曱-乀

来自广东
2020-08-04
做了两个多小时，终于做出来了，但是只能击败百分之个位数，又喜又悲。思路很简单，做了双向哈希映射，如果反查有，正查没有就删点，删完就成功，如果一轮查找没有进行任何删除，则失败。

评论(189)
xiezanmei
xzm
来自重庆
2021-11-24
根据官解的DFS方法对代码进行了详细的注释，希望可以帮助大家理解

class Solution {
    List<List<Integer>>list=new ArrayList<>();
    int visit[];
    boolean isvaild=true;
    public boolean canFinish(int numCourses, int[][] prerequisites) {
        visit=new int[numCourses];
        for (int i=0;i<numCourses;i++){
            list.add(new ArrayList<Integer>());//初始化
        }
        for (int info[]:prerequisites){
            //学info[0]之前要先学info[1],所以info[1]指向info[0],
            //所以在info[1]的ArrayList中存储它指向哪个科目
            list.get(info[1]).add(info[0]);
        }
        for (int i=0;i<numCourses;i++){
            if (visit[i]==0){//如果是未搜索的科目，则深搜
                dfs(i);
            }
        }
        return isvaild;
    }
    public void dfs(int v){
        visit[v]=1;//标记该科目为搜索中
        for (int w:list.get(v)){//遍历该科目指向的学科
            if (visit[w]==0){//如果指向的某一学科未搜索过，则深搜
                dfs(w);
                if (!isvaild){
                    return;
                }
            }
            else if (visit[w]==1){//如果指向的某一学科在搜索中，则有环，标记isVaild
                isvaild=false;
                return;
            }
        }
        visit[v]=2;//因为该科目已经完成深搜，所以标记它的状态未搜索过
    }
}
hjfenghj-f
hjfenghj
来自广东
2022-03-17
这是中等题嘛，为什么看了以后没有一点思路

practical-paremtq
我开拖拉机1号
来自北京
2022-05-03
拖拉机一号表示不会做

cristianoRonaldo7
繁星
来自辽宁
2021-03-31
我是啥比

mitchell_huang
Mitchell_Huang
来自四川
2022-01-19
佛了 没用深度优先和广度优先 一直超时 最后俩小时搞出来个双5%的垃圾

isXiaoSong
小宋啊
来自北京
2022-02-22
public class Solution {
    /**
     * 课程表
     * 你这个学期必须选修 numCourses 门课程，记为0到numCourses - 1 。
     * 在选修某些课程之前需要一些先修课程。 先修课程按数组prerequisites给出，
     * 其中prerequisites[i] = [ai, bi],表示如果要学习课程ai，则必须先学习课程bi 。
     * 输入：numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
     * 输出：true,一个可行的修课序列0->1->2->3
     */
    public boolean canFinish(int numCourses, int[][] pres) {
        // 本题转换为pres的有向无环图是否存在
        // = 求pres的对应图的拓扑排序，如果最后图非空或者已经不存在入度为0的结点，图必有环=返回false，否则返回true
        // graph：图邻接表，key是先修课程，value是必须先修课程之后才能学习的所有课程
        Map<Integer, List<Integer>> graph = new HashMap<>(numCourses);
        for (int i = 0; i < numCourses; i++) {
            graph.put(i, new LinkedList<>());
        }

        int[] inDegrees = new int[numCourses];
        for (int[] node : pres) {
            int cur = node[0];
            int pre = node[1];

            inDegrees[cur]++;
            graph.get(pre).add(cur);
        }

        // 拓扑排序(BFS)
        Queue<Integer> queue = new LinkedList<>();
        for (int i = 0; i < numCourses; i++) {
            if (inDegrees[i] == 0) {
                queue.offer(i);
            }
        }

        while (!queue.isEmpty()) {
            // 队列出队=入度为0的出队
            int node = queue.remove();

            // 课程完成拓扑数-1
            numCourses--;

            for (int next : graph.get(node)) {
                inDegrees[next]--;
                if (inDegrees[next] == 0) {
                    queue.offer(next);
                }
            }
        }
        // 完成拓扑排序的课程数量减0，说明完成所有的先修课程
        return numCourses == 0;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        int nums = 4;
        int[][] pres = {{1, 0}, {2, 0}, {3, 1}, {3, 2}};
        System.out.println(solution.canFinish(nums, pres));
    }
}

kind-matsumoto
G_G
来自山东
2021-07-15
广度优先还蛮好理解的，深度优先有点难

shou-mu-zhe
梦想身高1米8
来自浙江
2020-08-04
上午看到这个，真的催眠

joey91
Joey 🙃

来自江苏
2023-02-08
这题我感觉是有问题的：若存在 prerequisites[i] = [$a_i$, $a_i$] ，则表示想要学习课程 $a_i$ ，你需要先完成课程 $a_i$ ，逻辑上根本说不通的！！

我是从 210. 课程表 II 做完再来做这题的，心想不会这么简单吧？题目没说不能出现「自环」呀，结果一发 WA 懵逼了。

class Solution {
  public boolean canFinish(int n, int[][] prerequisites) {
    ArrayList<Integer>[] g = new ArrayList[n];
    Arrays.setAll(g, i -> new ArrayList<>());
    int[] inDeg = new int[n];
    for (int[] edge : prerequisites) {
      // if (edge[0] != edge[1]) {
        g[edge[1]].add(edge[0]);
        inDeg[edge[0]]++;
      // }
    }
    
    Queue<Integer> q = new LinkedList<>();
    for (int i = 0; i < n; i++) {
      if (inDeg[i] == 0) {
        q.offer(i);
      }
    }
    
    int idx = 0;
    while (!q.isEmpty()) {
      idx++;
      for (int v : g[q.poll()]) {
        if (--inDeg[v] == 0) {
          q.offer(v);
        }
      }
    }
    
    return idx == n;
  }
}
rendermelon
Luu Bruce
来自浙江
2021-11-19
方法一 C++版本，if else 太多，visited = 1、2、3 看的头晕，代码优化如下，更清晰了。

TODOIT 表示还未访问过改节点.

WORKING 表示该节点深搜还未回溯回来，若是深搜时碰到该点，说明有环.

FINISHED 表示该节点已经遍历结束.

#define TODOIT   0
#define WORKING  1
#define FINISHED 2

class Solution {
public:
    bool canFinish(int n, vector<vector<int>>& pre) {
        vector<int> visited(n, TODOIT);
        vector<vector<int>> adj(n, vector<int>());
        for (auto &info : pre)
            adj[info[1]].push_back(info[0]);

        for (int i = 0; i < n; i++)
            if (TODOIT == visited[i] && !dfs(adj, visited, i))
                return false;
        return true;
    }
private:
    // 本轮搜索中没有环返回 true，有环返回 false
    bool dfs(vector<vector<int>> &adj, vector<int> &visited, int v) {
        if (visited[v] == WORKING) return false;
        if (visited[v] == FINISHED) return true;
        visited[v] = WORKING;     // 正在
        for (int u : adj[v]) {
            if (!dfs(adj, visited, u))
                return false;
        }
        visited[v] = FINISHED;     // 结束
        return true;
    }
};


1

2

3

4

5


13

有用 308

299 条评论

收藏

分享
207



#207 课程表
中等
#1207 独一无二的出现次数
简单
#2070 每一个查询的最大美丽值
中等
#2071 你可以安排的最多任务数目
困难
#2072 赢得比赛的大学
简单
#2073 买票需要的时间
简单
#2074 反转偶数长度组的节点
中等
#2075 解码斜向换位密码
中等
#2076 处理含限制条件的好友请求
困难
#2077 殊途同归
中等
#2078 两栋颜色不同且距离最远的房子
简单
#2079 给植物浇水
中等
#2207 字符串中最多数目的子字符串
中等
