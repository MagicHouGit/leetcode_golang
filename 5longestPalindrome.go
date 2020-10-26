func longestPalindrome(s string) string {
	if len(s) == 0 {
		return s
	}

	bytes := make([]byte, 2*len(s)+1)
	bytes[0] = '#'
	for i := 0; i < len(s); i++ {
		bytes[2*i+1] = s[i]
		bytes[2*i+2] = '#'
	}

	lengths := make([]int, len(bytes))
	lengths[0] = 0
	lengths[1] = 1
	maxLength := 1
	maxIndex := 1

	for i := 2; i < len(bytes); i++ {
		right := i + min(lengths[i-2], lengths[i-1])

		for right < len(bytes) && 2*i-right >= 0 {
			if bytes[right] != bytes[2*i-right] {
				break
			}
			right += 1
		}
		lengths[i] = right - i - 1

		if lengths[i] > maxLength {
			maxIndex = i
			maxLength = lengths[i]
		}
	}

	start := maxIndex - lengths[maxIndex]
	newBytes := make([]byte, lengths[maxIndex])
	if bytes[start] == '#' {
		start += 1
	}
	for i := 0; i < len(newBytes); i++ {
		newBytes[i] = bytes[start+2*i]
	}
	return string(newBytes)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

