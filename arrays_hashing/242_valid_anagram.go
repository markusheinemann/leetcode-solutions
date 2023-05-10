// https://leetcode.com/problems/valid-anagram
// Time complexity: O(n)
// Space complexity: O(n)
// Runtime: 9 ms (beats 42.25%)
// Memory: 2.8 MB (beats 99.14%)

func isAnagram(s string, t string) bool {
    chars := make(map[rune]int)

    for _, char := range s {
        chars[char]++
    }

    for _, char := range t {
        if key, ok := chars[char]; !ok || key == 0 {
            return false
        }
        chars[char]--
    }

    return len(s) == len(t)
}
