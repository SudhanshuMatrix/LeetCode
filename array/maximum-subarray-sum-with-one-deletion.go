func maximumSum(arr []int) int {
    noDelete := arr[0]
    oneDelete := 0
    result := arr[0]

    for i := 1; i < len(arr); i++ {
        newOneDelete := max(
            oneDelete+arr[i], // deletion already used
            noDelete,         // delete arr[i]
        )

        newNoDelete := max(
            noDelete+arr[i],
            arr[i],
        )

        oneDelete = newOneDelete
        noDelete = newNoDelete

        result = max(result, max(noDelete, oneDelete))
    }

    return result
}