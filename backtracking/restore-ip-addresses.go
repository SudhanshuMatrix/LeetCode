func restoreIpAddresses(s string) []string {
    ans := []string{}

    for a := 1; a <= 3; a++ {
        for b := 1; b <= 3; b++ {
            for c := 1; c <= 3; c++ {
                for d := 1; d <= 3; d++ {
                    if a+b+c+d != len(s) {
                        continue
                    }

                    parts := []string{
                        s[:a],
                        s[a : a+b],
                        s[a+b : a+b+c],
                        s[a+b+c:],
                    }

                    valid := true

                    for _, p := range parts {
                        if len(p) > 1 && p[0] == '0' {
                            valid = false
                            break
                        }

                        num := 0
                        for i := 0; i < len(p); i++ {
                            num = num*10 + int(p[i]-'0')
                        }

                        if num > 255 {
                            valid = false
                            break
                        }
                    }

                    if valid {
                        ans = append(ans, parts[0]+"."+parts[1]+"."+parts[2]+"."+parts[3])
                    }
                }
            }
        }
    }

    return ans
}