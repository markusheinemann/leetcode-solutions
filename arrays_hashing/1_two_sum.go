// https://leetcode.com/problems/two-sum/
// Time complexity: O(n)
// Space complexity: O(n)
// Runtime: 4 ms (beats 89.94%)
// Memory: 4.3 MB (beats 36.13%)

func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    for i, val := range nums {
        if j, ok := m[target-val]; ok {
            return []int{i, j}
        }
        m[val] = i;
    }
    return []int{}
}
