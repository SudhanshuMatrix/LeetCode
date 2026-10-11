import "math"

func minSubArrayLen(target int, nums []int) int {
    high := 0
    low := 0
    sum := 0
    res := math.MaxInt

    for high < len(nums) {
        sum += nums[high]

        for sum >= target {
            length := high - low + 1
            if length < res {
                res = length
            }
            sum -= nums[low]
            low++
        }

        high++
    }

    if res == math.MaxInt {
        return 0
    }

    return res
}