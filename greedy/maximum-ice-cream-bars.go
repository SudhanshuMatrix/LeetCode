func maxIceCream(costs []int, coins int) int {

    const MAX_COST = 100000

    freq := make([]int, MAX_COST+1)

    for _, cost := range costs {
        freq[cost]++
    }

    answer := 0

    for cost := 1; cost <= MAX_COST; cost++ {

        if freq[cost] == 0 {
            continue
        }

        canBuy := freq[cost]
        affordable := coins / cost

        if canBuy > affordable {
            canBuy = affordable
        }

        answer += canBuy

        coins -= canBuy * cost
    }

    return answer
}