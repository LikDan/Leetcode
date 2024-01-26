package _3_longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {
	entries := make(map[int32]int)
	maxLength := 0
	substring := ""
	for _, c := range s {
		_, ok := entries[c]
		if ok {
			if len(entries) > maxLength {
				maxLength = len(entries)
			}

			for i, el := range substring {
				if el != c {
					delete(entries, el)
					continue
				}

				substring = substring[i+1:]
				break
			}
		}

		entries[c] = 1
		substring += string(c)
	}

	if len(entries) > maxLength {
		maxLength = len(entries)
	}
	return maxLength
}
