func sortedSquares(nums []int) []int {
	left, right := 0,len(nums)-1
    result := make([]int,len(nums))
    for i:=len(nums)-1;i>=0;i--{
        if nums[left]*nums[left]>nums[right]*nums[right]{
            result[i] = nums[left]*nums[left]
            left++
        }else{
            result[i] = nums[right]*nums[right]
            right--
        }
    }
    return result
}