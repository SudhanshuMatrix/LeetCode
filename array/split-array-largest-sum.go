func splitArray(nums []int, k int) int {
    low, high := 0, 0

    for _, num := range nums {
        if num > low {
            low = num
        }
        high += num
    }

    for low <= high {
        mid := low + (high-low)/2

        count := 1
        sum := 0

        for _, num := range nums {
            if sum+num > mid {
                count++
                sum = num
            } else {
                sum += num
            }
        }

        if count > k {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }

    return low
}