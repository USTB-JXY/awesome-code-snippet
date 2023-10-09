package leetcode

import "fmt"

type stringUnionFind struct {
	parents map[string]string
	vals    map[string]float64
}

func (suf stringUnionFind) add(x string) {
	if _, ok := suf.parents[x]; ok {
		return
	}
	suf.parents[x] = x
	suf.vals[x] = 1.0
}

func (suf stringUnionFind) find(x string) string {
	p := ""
	if v, ok := suf.parents[x]; ok {
		p = v
	} else {
		p = x
	}
	if x != p {
		pp := suf.find(p)
		suf.vals[x] *= suf.vals[p]
		suf.parents[x] = pp
	}
	if v, ok := suf.parents[x]; ok {
		return v
	}
	return x
}

func (suf stringUnionFind) union(x, y string, v float64) {
	suf.add(x)
	suf.add(y)
	px, py := suf.find(x), suf.find(y)
	suf.parents[px] = py
	// x / px = vals[x]
	// x / y = v
	// 由上面 2 个式子就可以得出 px = v * vals[y] / vals[x]
	suf.vals[px] = v * suf.vals[y] / suf.vals[x]
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	res, suf := make([]float64, len(queries)), stringUnionFind{parents: map[string]string{}, vals: map[string]float64{}}
	for i := 0; i < len(values); i++ {
		suf.union(equations[i][0], equations[i][1], values[i])
	}
	for i := 0; i < len(queries); i++ {
		x, y := queries[i][0], queries[i][1]
		if _, ok := suf.parents[x]; ok {
			if _, ok := suf.parents[y]; ok {
				if suf.find(x) == suf.find(y) {
					res[i] = suf.vals[x] / suf.vals[y]
				} else {
					res[i] = -1
				}
			} else {
				res[i] = -1
			}
		} else {
			res[i] = -1
		}
	}
	return res
}

func mycalcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	canreach := map[string][]string{}
	canreachvalue := map[string]float64{}
	res := []float64{}
	for i := 0; i < len(equations); i++ {
		canreach[equations[i][0]] = append(canreach[equations[i][0]], equations[i][1])
		canreachvalue[equations[i][0]+equations[i][1]] = values[i]
	}
	fmt.Println(canreach)
	fmt.Println(canreachvalue)
	for i := 0; i < len(queries); i++ {
		flag, value := false, 1.0
		if _, ok := canreach[queries[i][0]]; ok {
			bt(canreach, canreachvalue, queries[i][0], queries[i][1], &flag, &value)
		}
		if flag {
			res = append(res, value)
			continue
		}
		if _, ok := canreach[queries[i][1]]; ok {
			bt(canreach, canreachvalue, queries[i][1], queries[i][0], &flag, &value)
		}
		if flag {
			res = append(res, 1.0/value)
			continue
		}
		res = append(res, -1.0)
	}
	return res
}

func bt(canreach map[string][]string, canreachvalue map[string]float64, s1, s2 string, flag *bool, value *float64) {
	if s1 == s2 {
		*flag = true
		return
	}
	for _, v := range canreach[s1] {
		*value *= canreachvalue[s1+v]
		bt(canreach, canreachvalue, v, s2, flag, value)
		if *flag {
			return
		}
		*value /= canreachvalue[s1+v]
	}
}
