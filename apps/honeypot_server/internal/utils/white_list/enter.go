package white_list

// WhiteListCheck 检查是否在白名单内
func WhiteListCheck[T comparable](list []T, key T) bool {
	for _, t := range list {
		if t == key {
			return true
		}
	}
	return false
}
