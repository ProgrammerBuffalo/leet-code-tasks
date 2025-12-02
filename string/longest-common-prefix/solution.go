// https://leetcode.com/problems/longest-common-prefix/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	arr := []byte{}

	for offset := 0; ; offset++ {
		var b byte
		for i := 0; i < len(strs); i++ {
			if len(strs[i]) != offset {
				if i == 0 {
					b = strs[i][offset]
				} else if b != strs[i][offset] {
					return string(arr)
				}
			} else {
				return string(arr)
			}
		}
		arr = append(arr, b)
	}
  
	return string(arr)
}
