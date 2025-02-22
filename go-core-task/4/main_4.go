package main

func compareSlices(s1, s2 []string) []string {
	out := make([]string, 0)
	existsInS1 := make(map[string]bool)

	for _, val := range s1 {
		existsInS1[val] = true
	}

	for _, val := range s2 {
		if existsInS1[val] == false {
			out = append(out, val)
			existsInS1[val] = true
		}
	}

	return out
}