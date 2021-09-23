package solution

func RemoveComments(source []string) []string {
	inBlock, res, newLine := false, make([]string, 0, len(source)), ""
	for _, line := range source {
		if !inBlock {
			newLine = ""
		}

		for i := 0; i < len(line); i++ {
			if i < len(line)-1 {
				if string(line[i:i+2]) == "/*" && !inBlock {
					inBlock = true
					i++
					continue
				} else if string(line[i:i+2]) == "*/" && inBlock {
					inBlock = false
					i++
					continue
				} else if !inBlock && string(line[i:i+2]) == "//" {
					break
				}
			}
			if !inBlock {
				newLine += string(line[i])
			}
		}

		if newLine != "" && !inBlock {
			res = append(res, newLine)
		}
	}

	return res
}
