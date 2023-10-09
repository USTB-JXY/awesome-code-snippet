package leetcode

func isIsomorphic(s string, t string) bool {
	strList := []byte(t)
	patternByte := []byte(s)
	if (s == "" && t != "") || (len(patternByte) != len(strList)) {
		return false
	}

	pMap := map[byte]byte{}
	sMap := map[byte]byte{}
	for index, b := range patternByte {
		if _, ok := pMap[b]; !ok {
			if _, ok = sMap[strList[index]]; !ok {
				pMap[b] = strList[index]
				sMap[strList[index]] = b
			} else {
				if sMap[strList[index]] != b {
					return false
				}
			}
		} else {
			if pMap[b] != strList[index] {
				return false
			}
		}
	}
	return true
}

func myisIsomorphic0(s string, t string) bool {
	strList := []byte(t)
	patternByte := []byte(s)
	if (s == "" && t != "") || (len(patternByte) != len(strList)) {
		return false
	}
	//pMap, sMap := map[byte]int{}, map[byte]int{}

	return true
}

// bool isIsomorphic(string s, string t) {
// 	if (s.size()!=t.size()) return false;

// 	const int MAXCHAR = 256;
// 	char maps[MAXCHAR]={0}, mapt[MAXCHAR]={0};
// 	//memset(maps, 0, sizeof(maps));
// 	//memset(mapt, 0, sizeof(mapt));

// 	for(int i=0; i<s.size(); i++){
// 		if(maps[s[i]] == 0 && mapt[t[i]] == 0){
// 			maps[s[i]] = t[i];
// 			mapt[t[i]] = s[i];
// 			continue;
// 		}
// 		if(maps[s[i]] == t[i] && mapt[t[i]] == s[i]) {
// 			continue;
// 		}
// 		return false;
// 	}
// 	return true;
// }

func myisIsomorphic(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}
	maps, mapt := map[byte]byte{}, map[byte]byte{}
	for i := 0; i < len(s); i++ {
		_, sok := maps[s[i]]
		_, tok := mapt[t[i]]
		if !sok && !tok {
			maps[s[i]] = s[i]
			mapt[t[i]] = t[i]
			continue
		}
		if (sok && tok) && (maps[s[i]] == s[i]) && (mapt[t[i]] == t[i]) {
			continue
		}
		return false
	}
	return true
}
