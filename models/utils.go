package models

func Contains(arr []int, u uint) bool {
	i := int(u)
	for _, n := range arr {
		if n == i {
			return true
		}
	}
	return false
}
