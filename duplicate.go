package golangdulicatecount

func DuplicateCount(datas []int) int {
	tempNum := datas
	var skipped []int
	count := 0
	for i, v := range datas {
		var check = false
		var skip = false

		for x := 0; x < len(skipped); x++ {
			if v == skipped[x] {
				skip = true
				break
			}
		}
		if skip {
			break
		}

		for w := i + 1; w < len(tempNum); w++ {
			if v == tempNum[w] {
				check = true
				break
			}
		}
		if check {
			skipped = append(skipped, v)
			count++
		}
	}

	return count
}
