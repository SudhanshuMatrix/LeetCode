
func subarraysDivByK(nums []int, k int) int {
    freq := make([]int, k)
    freq[0] = 1

    sum := 0
    ans := 0

    for _, num := range nums {
        sum += num

        rem := sum % k
        if rem < 0 {
            rem += k
        }

        ans += freq[rem]
        freq[rem]++
    }

    return ans
}