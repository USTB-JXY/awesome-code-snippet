package leetcode

import "sort"

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func groupAnagrams(strs []string) [][]string {
	record, res := map[string][]string{}, [][]string{}
	for _, str := range strs {
		sByte := []rune(str)
		sort.Sort(sortRunes(sByte))
		sstrs := record[string(sByte)]
		sstrs = append(sstrs, str)
		record[string(sByte)] = sstrs
	}
	for _, v := range record {
		res = append(res, v)
	}
	return res
}
func mygroupAnagrams(strs []string) [][]string {
	record, res := map[string][]string{}, [][]string{}
	for _, str := range strs {
		sByte := []rune(str)
		sort.Sort(sortRunes(sByte))
		sstr := record[string(sByte)]
		sstr = append(sstr, str)
		record[string(sByte)] = sstr
	}
	for _, v := range record {
		res = append(res, v)
	}
	return res
}

// // byte is an alias for uint8 and is equivalent to uint8 in all ways. It is
// // used, by convention, to distinguish byte values from 8-bit unsigned
// // integer values.
// type byte = uint8

// // rune is an alias for int32 and is equivalent to int32 in all ways. It is
// // used, by convention, to distinguish character values from integer values.
// type rune = int32
// Copy
// 原来是 byte 表示一个字节，rune 表示四个字节，那么就可以得出了结论了，来看一段代码，使用中文字符串

// first := "社区"
// fmt.Println([]rune(first))
// fmt.Println([]byte(first))
// [31038 21306] // 输出结果 [] rune
// [231 164 190 229 140 186]// 输出结果 [] byte

// 这里也可以很清晰的看出这里的中文字符串每个占三个字节， 区别也就一目了然了。
// 说道这里正好可以提一下 Go 语言切割中文字符串，Go 的字符串截取和切片是一样的 s [n:m] 左闭右开的原则，看一下例子

// s := "golangcaff"
// fmt.Println(s[:3])
// gol // 输出，看起来没问题， 顺利截取了三个字符

// 如果换成中文的呢？来看一下例子

// s := "截取中文"
// //试试这样能不能截取?
// fmt.Println(s[:2])
// ?? // 输出 在预料之中， 输出了常见的？？

// 那么该如何截取呢？这里就需要将中文利用 [] rune 转换成 unicode 码点， 再利用 string 转化回去， 来看一下例子。

// s := "截取中文"
// //试试这样能不能截取?
// res := []rune(s)
// fmt.Println(string(res[:2]))
// 截取 // 输出，顺利截取了
