package leetcode

import (
	"fmt"
)

// 解法一 位图
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var bitSet [256]bool
	result, left, right := 0, 0, 0
	for left < len(s) {
		// 右侧字符对应的 bitSet 被标记 true，说明此字符在 X 位置重复，需要左侧向前移动，直到将 X 标记为 false
		if bitSet[s[right]] {
			bitSet[s[left]] = false
			left++
		} else {
			bitSet[s[right]] = true
			right++
		}
		if result < right-left {
			result = right - left
		}
		if left+result >= len(s) || right >= len(s) {
			break
		}
	}
	return result
}

// 解法二 滑动窗口
func ls(s string) int {
	if len(s) == 0 {
		return 0
	}
	var freq [127]int
	result, left, right := 0, 0, -1
	for left < len(s) {
		if right+1 < len(s) && freq[s[right+1]] == 0 {
			freq[s[right+1]]++
			right++
		} else {
			freq[s[left]]--
			left++
		}
		fmt.Println(left, right)
		result = max(result, right-left+1)
	}
	return result
}

// 过程            窗口   left   right
// "abccef"        abc    0       2
// "abccef"        bc     1       2
// "abccef"        c      2       2
// "abccef"               3       2
// "abccef"        c      3       3
// 0 0
// 0 1
// 0 2
// 1 2
// 2 2
// 3 2
// 3 3
// 3 4
// 3 5
// 4 5
// 5 5
// 6 5
// 【input】:{abccef}     【ans】:{3}     【output】:3
func lengthOfLongestSubstring1(s string) int {
	if len(s) == 0 {
		return 0
	}
	var freq [127]int
	result, left, right := 0, 0, -1

	for left < len(s) {
		if right+1 < len(s) && freq[s[right+1]] == 0 {
			freq[s[right+1]]++
			right++

		} else {
			freq[s[left]]--
			left++
		}
		result = max(result, right-left+1)
	}
	return result
}

func ls2(s string) int {
	right, left, res := 0, 0, 0
	indexs := make(map[byte]int, len(s))
	for right < len(s) {
		fmt.Println(right, left)
		if idx, ok := indexs[s[right]]; ok && idx >= left {
			left = idx + 1
		}
		indexs[s[right]] = right
		right++

		res = max(res, right-left)
	}
	return res
}

// 0 0
// 1 0
// 2 0
// 3 0
// 4 3
// 5 3         0 1 2 3 4
// 【input】:{ a b c c e f }     【ans】:{3}     【output ls2】:3
// 解法三 滑动窗口-哈希桶
func lengthOfLongestSubstring2(s string) int {
	right, left, res := 0, 0, 0
	indexes := make(map[byte]int, len(s))
	for left < len(s) {
		if idx, ok := indexes[s[left]]; ok && idx >= right {
			right = idx + 1
		}
		indexes[s[left]] = left
		left++
		res = max(res, left-right)
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// ASCII值	控制字符	ASCII值	控制字符	ASCII值	控制字符	ASCII值	控制字符
// 0	NUT	32	(space)	64	@	96	、
// 1	SOH	33	!	65	A	97	a
// 2	STX	34	"	66	B	98	b
// 3	ETX	35	#	67	C	99	c
// 4	EOT	36	$	68	D	100	d
// 5	ENQ	37	%	69	E	101	e
// 6	ACK	38	&	70	F	102	f
// 7	BEL	39	,	71	G	103	g
// 8	BS	40	(	72	H	104	h
// 9	HT	41	)	73	I	105	i
// 10	LF	42	*	74	J	106	j
// 11	VT	43	+	75	K	107	k
// 12	FF	44	,	76	L	108	l
// 13	CR	45	-	77	M	109	m
// 14	SO	46	.	78	N	110	n
// 15	SI	47	/	79	O	111	o
// 16	DLE	48	0	80	P	112	p
// 17	DCI	49	1	81	Q	113	q
// 18	DC2	50	2	82	R	114	r
// 19	DC3	51	3	83	S	115	s
// 20	DC4	52	4	84	T	116	t
// 21	NAK	53	5	85	U	117	u
// 22	SYN	54	6	86	V	118	v
// 23	TB	55	7	87	W	119	w
// 24	CAN	56	8	88	X	120	x
// 25	EM	57	9	89	Y	121	y
// 26	SUB	58	:	90	Z	122	z
// 27	ESC	59	;	91	[	123	{
// 28	FS	60	<	92	\	124	|
// 29	GS	61	=	93	]	125	}
// 30	RS	62	>	94	^	126	`
// 31	US	63	?	95	_	127	DEL
