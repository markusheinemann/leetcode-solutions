// https://leetcode.com/problems/contains-duplicate
// Time complexity: O(n)
// Space complexity: O(n)
// Runtime: 93 ms (beats 86.3%)
// Memory: 8.5 MB (beats 95.84%)

// The main idea of this solution is to create a map 
// which will have a specific value from the input
// array as keys. For each element of the input array
// we check if the map already has an key with the 
// current value. If that is the case a duplicate 
// exists and true will be returend. Otherwise the
// current value is added as key to the map. If the loop
// is run through completely, no duplicate exists and 
// true will be returned.

func containsDuplicate(nums []int) bool {
    m := make(map[int]struct{})
    for _, num := range nums {
        if _, ok := m[num]; ok {
            return true
        }
        m[num] = struct{}{}
    }
    return false
}
