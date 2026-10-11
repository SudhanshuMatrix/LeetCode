func pivotIndex(nums []int) int {
    total := 0
    for _, num := range nums {
        total += num
    }

    left := 0
    for i := 0; i < len(nums); i++ {
        right := total - left - nums[i]
        if left == right {
            return i
        }
        left += nums[i]
    }

    return -1
}