package leetcode

import "fmt"

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordMap, que, depth := getWordMap(wordList, beginWord), []string{beginWord}, 0
	for len(que) > 0 {
		depth++
		qlen := len(que)
		for i := 0; i < qlen; i++ {
			word := que[0]
			que = que[1:]
			candidates := getCandidates(word)
			for _, candidate := range candidates {
				if _, ok := wordMap[candidate]; ok {
					if candidate == endWord {
						return depth + 1
					}
					delete(wordMap, candidate)
					que = append(que, candidate)
				}
			}
		}
	}
	return 0
}

func getWordMap(wordList []string, beginWord string) map[string]int {
	wordMap := make(map[string]int)
	for i, word := range wordList {
		if _, ok := wordMap[word]; !ok {
			if word != beginWord {
				wordMap[word] = i
			}
		}
	}
	return wordMap
}

func getCandidates(word string) []string {
	var res []string
	for i := 0; i < 26; i++ {
		for j := 0; j < len(word); j++ {
			if word[j] != byte(int('a')+i) {
				res = append(res, word[:j]+string(rune(int('a')+i))+word[j+1:])
			}
		}
	}
	return res
}

func tj1ladderLength(beginWord string, endWord string, wordList []string) int {
	mWord := map[string]bool{}
	cnt := 0

	// 第 1 步：先将 wordList 放到哈希表里，便于判断某个单词是否在 wordList 里
	for _, v := range wordList {
		mWord[v] = true
		cnt++
	}

	if cnt == 0 || mWord[endWord] == false {
		return 0
	}
	var update func(*[]string, map[string]int, map[string]int) int
	var bfs func() int
	bfs = func() int {
		// d1 代表从起点 beginWord 开始搜索（正向）
		// d2 代表从结尾 endWord 开始搜索（反向）
		d1 := []string{}
		d2 := []string{}
		d1 = append(d1, beginWord)
		d2 = append(d2, endWord)

		/*
		 * m1 和 m2 分别记录两个方向出现的单词是经过多少次转换而来
		 * e.g.
		 * m1 = {"abc":1} 代表 abc 由 beginWord 替换 1 次字符而来
		 * m2 = {"xyz":3} 代表 xyz 由 endWord 替换 3 次字符而来
		 */
		m1 := map[string]int{}
		m2 := map[string]int{}
		m1[beginWord] = 0
		m2[endWord] = 0

		/*
		 * 只有两个队列都不空，才有必要继续往下搜索
		 * 如果其中一个队列空了，说明从某个方向搜到底都搜不到该方向的目标节点
		 * e.g.
		 * 例如，如果 d1 为空了，说明从 beginWord 搜索到底都搜索不到 endWord，反向搜索也没必要进行了
		 */
		for len(d1) > 0 && len(d2) > 0 {
			t := -1

			// 为了让两个方向的搜索尽可能平均，优先拓展队列内元素少的方向
			if len(d1) <= len(d2) {
				t = update(&d1, m1, m2)
			} else {
				t = update(&d2, m2, m1)
			}

			if t != -1 {
				return t
			}
		}

		return -1
	}

	// update 代表从 d 中取出一个单词进行扩展，
	// m1 为当前方向的距离字典；
	// m2 为另外一个方向的距离字典

	update = func(d1 *[]string, m1 map[string]int, m2 map[string]int) int {
		d := *d1
		n := len(d)
		// k:=0
		for n > 0 {
			x := d[0]
			d = d[1:]
			m := len(x)

			// 枚举替换原字符串的哪个字符 i
			for i := 0; i < m; i++ {
				// 枚举将 i 替换成哪个小写字母
				for j := 0; j < 26; j++ {
					sub := []byte{}
					sub = append(sub, x[:i]...)
					sub = append(sub, byte('a'+j))
					sub = append(sub, x[i+1:]...)
					//fmt.Println(string(sub))
					if mWord[string(sub)] {

						// 如果该字符串在「当前方向」被记录过（拓展过），跳过即可
						//备注：存在和赋值为0的区别

						if _, ok := m1[string(sub)]; ok && m1[string(sub)] <= m1[x]+1 {
							// fmt.Println(k)
							// k++
							continue
						}

						// 如果该字符串在「另一方向」出现过，说明找到了联通两个方向的最短路

						if _, ok := m2[string(sub)]; ok {
							// fmt.Println(1)
							return m1[x] + 1 + m2[string(sub)]
						} else {
							// fmt.Println(2)
							d = append(d, string(sub))
							m1[string(sub)] = m1[x] + 1
						}
					}
				}
			}

			n--
		}
		*d1 = d
		return -1
	}

	ans := bfs()
	if ans == -1 {
		return 0
	}
	return ans + 1
}

func myladderLength(start string, end string, bank []string) int {
	if start == end {
		return 0
	}
	//退化为Set
	bankSet := map[string]struct{}{}
	for _, v := range bank {
		bankSet[v] = struct{}{}
	}
	if _, ok := bankSet[end]; !ok {
		return 0
	}
	q := []string{start}
	for step := 1; q != nil; step++ {
		tmp := q
		q = nil
		for _, cur := range tmp {
			for i, v := range cur {
				for _, newv := range "abcdefghijklmnopqrstuvwxyz" {
					if v != newv {
						nxt := cur[:i] + string(newv) + cur[i+1:]
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
	return 0
}

func mydoubleladderLength(start string, end string, bank []string) int {
	if start == end {
		return 0
	}
	//退化为Set
	bankSet := map[string]struct{}{}
	for _, v := range bank {
		bankSet[v] = struct{}{}
	}
	if _, ok := bankSet[end]; !ok {
		return 0
	}
	q := []string{start}
	for step := 2; q != nil; step += 2 {
		tmp := q
		q = nil
		for _, cur := range tmp {
			for i, v := range cur {
				for _, newv := range "abcdefghijklmnopqrstuvwxyz" {
					if v != newv {
						nxt := cur[:i] + string(newv) + cur[i+1:]

						if nxt == end {
							return step + 1
						}
						if _, ok := bankSet[nxt]; ok {
							for j, v2 := range end {
								for _, newv2 := range "abcdefghijklmnopqrstuvwxyz" {
									if newv2 != v2 {
										endtemp := end[:j] + string(newv2) + end[j+1:]
										if _, ok := bankSet[endtemp]; ok {
											end = endtemp
											fmt.Println(nxt, end)
											if nxt == end {
												return step + 2
											}
											delete(bankSet, end)
											break
										}
									}
								}
							}

							delete(bankSet, nxt)
							q = append(q, nxt)
						}
					}
				}
			}
		}
	}
	return 0
}
