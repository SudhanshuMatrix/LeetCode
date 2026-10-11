func minDays(bloomDay []int, m int, k int) int {
    if m*k > len(bloomDay) {
        return -1
    }

    low, high := bloomDay[0], bloomDay[0]

    for _, day := range bloomDay {
        if day < low {
            low = day
        }
        if day > high {
            high = day
        }
    }

    for low <= high {
        mid := low + (high-low)/2

        bouquets := 0
        flowers := 0

        for _, day := range bloomDay {
            if day <= mid {
                flowers++

                if flowers == k {
                    bouquets++
                    flowers = 0
                }
            } else {
                flowers = 0
            }
        }

        if bouquets >= m {
            high = mid - 1
        } else {
            low = mid + 1
        }
    }

    return low
}