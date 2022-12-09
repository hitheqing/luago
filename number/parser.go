package number

import "strconv"

// ParseInteger 解析字符串为整数
func ParseInteger(s string) (int64, bool) {
	i, err := strconv.ParseInt(s, 10, 64)
	return i, err == nil
}

// ParseFloat 解析字符串为浮点数
func ParseFloat(s string) (float64, bool) {
	f, err := strconv.ParseFloat(s, 64)
	return f, err == nil
}
