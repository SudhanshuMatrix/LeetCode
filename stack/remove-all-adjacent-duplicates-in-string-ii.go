func removeDuplicates(s string, k int) string {
	st := make([][2]int, 0, len(s))

	for i := 0; i < len(s); i++ {
		c := s[i]

		if len(st) == 0 {
			st = append(st, [2]int{int(c), 1})
			continue
		}

		if st[len(st)-1][0] == int(c) {
			st[len(st)-1][1]++

			if st[len(st)-1][1] == k {
				st = st[:len(st)-1]
			}
		} else {
			st = append(st, [2]int{int(c), 1})
		}
	}

	res := make([]byte, 0, len(s))

	for _, p := range st {
		for i := 0; i < p[1]; i++ {
			res = append(res, byte(p[0]))
		}
	}

	return string(res)
}