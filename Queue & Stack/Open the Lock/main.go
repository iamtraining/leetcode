package main

func main() {

}

func openLock(deadends []string, target string) int {
	deads := map[string]struct{}{}
	for _, dead := range deadends {
		deads[dead] = struct{}{}
	}
	visited := map[string]struct{}{}
	q := []string{"0000"}
	lvl := 0
	for len(q) > 0 {
		qs := len(q)
		for i := 0; i < qs; i++ {
			val := q[0]
			q = q[1:]
			if _, ok := deads[val]; ok {
				continue
			}
			if val == target {
				return lvl
			}
			for j := 0; j < 4; j++ {
				next := helper(val[:j], val[j+1:], val[j], false)
				if _, ok := visited[next]; !ok {
					q = append(q, next)
					visited[next] = struct{}{}
				}
				prev := helper(val[:j], val[j+1:], val[j], true)
				if _, ok := visited[prev]; !ok {
					q = append(q, prev)
					visited[prev] = struct{}{}
				}
			}
		}
		lvl++
	}
	return -1
}

func helper(prefix, postfix string, c byte, prev bool) string {
	if prev {
		if c == '0' {
			c = '9'
		} else {
			c--
		}
	} else {
		if c == '9' {
			c = '0'
		} else {
			c++
		}
	}
	return prefix + string(c) + postfix
}
