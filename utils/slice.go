package utils

func StringUnique(strSlice []string) []string {
	keys := make(map[string]bool)
	var list []string
	for _, entry := range strSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func StringSliceContains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

func StringSliceFilter(slice []string, item string) []string {
	vsf := make([]string, 0)
	for _, v := range slice {
		if v != item {
			vsf = append(vsf, v)
		}
	}
	return vsf
}