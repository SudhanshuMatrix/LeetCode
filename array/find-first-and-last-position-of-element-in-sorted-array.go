func searchRange(nums []int, target int) []int {
    first := -1
    last := -1

    low, high := 0, len(nums)-1
    for low <= high {
        mid := low + (high-low)/2

        if nums[mid] == target {
            first = mid
            high = mid - 1
        } else if nums[mid] < target {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }

    low, high = 0, len(nums)-1
    for low <= high {
        mid := low + (high-low)/2

        if nums[mid] == target {
            last = mid
            low = mid + 1
        } else if nums[mid] < target {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }

    return []int{first, last}
}