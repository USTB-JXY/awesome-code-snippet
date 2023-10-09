package leetcode

import (
	"strconv"
	"strings"
)

func addBinary(a string, b string) string {
	if len(b) > len(a) {
		a, b = b, a
	}

	res := make([]string, len(a)+1)
	i, j, k, c := len(a)-1, len(b)-1, len(a), 0
	for i >= 0 && j >= 0 {
		ai, _ := strconv.Atoi(string(a[i]))
		bj, _ := strconv.Atoi(string(b[j]))
		res[k] = strconv.Itoa((ai + bj + c) % 2)
		c = (ai + bj + c) / 2
		i--
		j--
		k--
	}

	for i >= 0 {
		ai, _ := strconv.Atoi(string(a[i]))
		res[k] = strconv.Itoa((ai + c) % 2)
		c = (ai + c) / 2
		i--
		k--
	}

	if c > 0 {
		res[k] = strconv.Itoa(c)
	}

	return strings.Join(res, "")
}
func B2I(b byte) int {
	res, _ := strconv.Atoi(string(b))
	return res
}
func I2S(b int) string {
	res := strconv.Itoa(b)
	return res
}
func myaddBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}
	i, j, c, k := len(a)-1, len(b)-1, 0, len(a)
	res := make([]string, len(a)+1)
	for ; j >= 0; k, i, j = k-1, i-1, j-1 {
		res[k] = I2S((c + B2I(a[i]) + B2I(b[j])) % 2)
		c = (c + B2I(a[i]) + B2I(b[j])) / 2
	}
	for ; i >= 0; k, i = k-1, i-1 {
		res[k] = I2S((c + B2I(a[i])) % 2)
		c = (c + B2I(a[i])) / 2
	}
	if c == 1 {
		res[k] = I2S(c)
	}
	return strings.Join(res, "")
}
