package search

// normal search string
func NormalSearchString(target, search []byte) int {
	var i, j int
	for i < len(target) && j < len(search) {
		if target[i] == search[j] {
			i++
			j++
		} else {
			i = i - j + 1
			j = 0
		}

		if j == len(search) {
			return i - j
		}
	}
	return -1
}

func GetNextValueArray(sub []byte) (next []int) {
	var length = len(sub)
	var middle, compareLeft, compareRight, matchCount int

	next = make([]int, length)
	next[0] = 0
	next[1] = 0

	for i := 2; i < length; i++ {
		middle = i / 2
		matchCount = 0
		if i%2 == 0 {
			for j := 0; j < middle; j++ {
				compareLeft = 0
				compareRight = i - 1 - j
				for compareLeft <= j {
					if sub[compareLeft] != sub[compareRight] {
						break
					}
					compareLeft++
					compareRight++
				}
				if compareLeft == j+1 {
					matchCount++
				}
			}
			next[i] = matchCount

		} else {
			for j := 0; j <= middle; j++ {
				compareLeft = 0
				compareRight = i - 1 - j
				for compareLeft <= j {
					if sub[compareLeft] != sub[compareRight] {
						break
					}
					compareLeft++
					compareRight++
				}
				if compareLeft == j+1 {
					matchCount++
				}
			}
			next[i] = matchCount
		}
	}
	return next
}

func ReviseNextValueArray(next []int) []int {
	var length = len(next)
	for i := 2; i < length; i++ {
		if next[i] == next[next[i]] {
			next[i] = next[next[i]]
		}
	}
	return next
}

// kmp string search
func KMPStringSearch(content []byte, sub []byte) (index int) {
	var (
		next      = ReviseNextValueArray(GetNextValueArray(sub))
		subIndex  = 0
		subLength = len(sub)
	)
	for i := 0; i <= len(content); i++ {
		if content[i] == sub[subIndex] {
			matchStart := i
			for j := subIndex; j <= subLength; j++ {
				if j == subLength {
					return matchStart - subIndex
				}
				if i >= len(content) || content[i] != sub[j] {
					subIndex = next[j]
					break
				}
				i++
			}
		}
	}
	return -1
}
