package leetcode

import "fmt"

func convert(s string, numRows int) string {
	matrix, down, up := make([][]byte, numRows, numRows), 0, numRows-2
	for i := 0; i != len(s); {
		if down != numRows {
			matrix[down] = append(matrix[down], byte(s[i]))
			down++
			i++
		} else if up > 0 {
			matrix[up] = append(matrix[up], byte(s[i]))
			up--
			i++
		} else {
			up = numRows - 2
			down = 0
		}
	}
	solution := make([]byte, 0, len(s))
	for _, row := range matrix {
		for _, item := range row {
			solution = append(solution, item)
		}
	}
	return string(solution)
}

// string convert(string s, int nRows) {
//     //The cases no need to do anything
//     if (nRows<=1 || nRows>=s.size()) return s;

//     vector<string> r(nRows);
//     int row = 0;
//     int step = 1;
//     for(int i=0; i<s.size(); i ++) {
//         if (row == nRows-1) step = -1;
//         if (row == 0) step = 1;
//         //cout << row <<endl;
//         r[row] += s[i];
//         row += step;
//     }

//	    string result;
//	    for (int i=0; i<nRows; i++){
//	        result += r[i];
//	    }
//	    return result;
//	}
func myconvert(s string, numRows int) string {
	fmt.Printf("---j---\n")
	if len(s) <= 1 || len(s) <= numRows {
		return s
	}
	res, step, index := "", 1, 0
	strvec := make([]string, numRows)
	for i := 0; i < len(s); i++ {
		if index == 0 {
			step = 1
		}
		if index == numRows-1 {
			step = -1
		}
		strvec[index] += s[i : i+1]
		index += step
	}
	for i := 0; i < numRows; i++ {
		res += strvec[i]
	}
	return res
}

// 【input】:{PAYPALISHIRING 3}   【ans】:{PAHNAPLSIIGYIR}     【output】:PAHNAPLSIIGYIR
// ---j---
// PAHNAPLSIIGYIR
// PAHNAPLSIIGYIR
// 【input】:{PAYPALISHIRING 4}   【ans】:{PINALSIGYAHRPI}     【output】:PINALSIGYAHRPI
// PINALSIGYAHRPI
// PINALSIGYAHRPI
