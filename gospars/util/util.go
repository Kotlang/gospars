package util

func MergeMaps(map1 map[string]string, map2 map[string]string) map[string]string {
	for k, v := range map2 {
		map1[k] = v
	}
	return map1
}
