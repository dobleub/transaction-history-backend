package helpers

func SortByDate(data [][]string, by string) [][]string {
	var tmp = data

	// sort transactions by date
	for i := 1; i < len(data); i++ {
		for j := 1; j < len(data)-i; j++ {
			if by == "desc" && StringToDate(data[j][3]).Before(StringToDate(data[j+1][3])) {
				// swap data in slice
				tmp[j], tmp[j+1] = tmp[j+1], tmp[j]
			}
			if by == "asc" && StringToDate(data[j][3]).After(StringToDate(data[j+1][3])) {
				// swap data in slice
				tmp[j], tmp[j+1] = tmp[j+1], tmp[j]
			}
		}
	}

	return tmp
}
