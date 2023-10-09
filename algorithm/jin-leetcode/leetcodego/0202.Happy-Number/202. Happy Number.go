package leetcode

func isHappy(n int) bool {
	record := map[int]int{}
	for n != 1 {
		record[n] = n
		n = getSquareOfDigits(n)
		for _, previous := range record {
			if n == previous {
				return false
			}
		}
	}
	return true
}

func getSquareOfDigits(n int) int {
	squareOfDigits := 0
	temporary := n
	for temporary != 0 {
		remainder := temporary % 10
		squareOfDigits += remainder * remainder
		temporary /= 10
	}
	return squareOfDigits
}
func mygetSquareOfDigits(n int) int {
	res := 0
	temp := n
	for temp != 0 {
		r := temp % 10
		res += r * r
		temp /= 10
	}
	return res
}

func myisHappy(n int) bool {
	record := map[int]int{}
	for n != 1 {
		record[n] = n
		n = mygetSquareOfDigits(n)
		for _, v := range record {
			if v == n {
				return false
			}
		}
	}
	return true
}
