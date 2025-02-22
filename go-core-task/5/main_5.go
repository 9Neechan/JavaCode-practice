package main

func compareSlices(s1, s2 []int) (bool, []int) {
	out := make([]int, 0)
	existsInS1 := make(map[int]bool)

	for _, val := range s1 {
		existsInS1[val] = true
	}

	for _, val := range s2 {
		if existsInS1[val] == true {
			out = append(out, val)
			existsInS1[val] = false
		}
	}

	if len(out) == 0 {
		return false, out
	}

	return true, out
}