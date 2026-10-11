func grayCode(n int) []int {
    size := 1 << n
    ans := make([]int, size)

    for i := 0; i < size; i++ {
        ans[i] = i ^ (i >> 1)
    }

    return ans
}