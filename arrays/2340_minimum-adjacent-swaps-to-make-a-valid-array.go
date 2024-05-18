// https://leetcode.com/problems/minimum-adjacent-swaps-to-make-a-valid-array/

// 1. Determine position of high and low
// 2. Swap high to end
// 3. Swap low to start

func minimumSwaps(nums []int) int {
    var maxVal, maxPos int
    var minVal, minPos int

    for i, val := range nums {
        if i == 0 {
            maxVal, maxPos = val, i
            minVal, minPos = val, i
            continue
        }

        // val greater equal to get "last" highest number
        // when highest number is availible multiple time
        if val >= maxVal {
            maxVal, maxPos = val, i
        }
        // val lower than to get "first" min value
        if val < minVal {
            minVal, minPos = val, i
        }
    }

    maxSwaps := len(nums) - 1 - maxPos
    minSwaps := minPos

    if maxPos < minPos {
        minSwaps--
    }

    return maxSwaps + minSwaps
}
