package fplib

// 从map[string]interface{}中获取字符串
func MapGetString(m *map[string]interface{}, key string, defaults ...string) string {
	result := ""
	if len(defaults) > 0 {
		result = defaults[0]
	}
	if iv, ok := (*m)[key]; ok {
		if v, ok := iv.(string); ok {
			result = v
		}
	}
	return result
}
