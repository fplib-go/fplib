package fplib

import (
	"strings"
)

type StrClass struct{}

// 字符串转换为数组
func (this *StrClass) ToArr(strs ...string) []string {
	result := make([]string, 0)
	l := len(strs)
	c := ""
	str := ""
	switch l {
	case 1:
		str = strs[0]
		c = ","
	case 2:
		str = strs[0]
		c = strs[1]
	default:
		return nil
	}

	res := strings.Split(str, c)
	for _, v := range res {
		val := Trim(v)
		if val != "" {
			result = append(result, val)
		}
	}
	return result
}

// 数组转换为字符串
func (this *StrClass) FromArr(arr []string) string {
	return strings.Join(arr, ",")
}

// 数组转换为字符串，可定义分隔符
func (this *StrClass) FromArr2(arr []string, c string) string {
	return strings.Join(arr, c)
}

// 字符串数组去重
func (this *StrClass) RemoveRepeatedFromArr(arr []string) []string {
	newArr := make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return newArr
}

// 数组中是否包含指定字符串，如果存在返回索引，不存在返回-1
func (this *StrClass) In_Array_Index(array []string, val string) int {
	index := -1
	for i := 0; i < len(array); i++ {
		if array[i] == val {
			index = i
			return index
		}
	}
	return index
}

// 数组中是否包含
func (this *StrClass) In_Array(array []string, val string) bool {
	index := this.In_Array_Index(array, val)
	result := false
	if index > -1 {
		result = true
	}
	return result
}

// 忽略大小写版本
func (this *StrClass) In_Array_Indexi(array []string, val string) int {
	index := -1
	for i := 0; i < len(array); i++ {
		if strings.ToLower(array[i]) == strings.ToLower(val) {
			index = i
			return index
		}
	}
	return index
}

func (this *StrClass) In_Arrayi(array []string, val string) bool {
	index := this.In_Array_Indexi(array, val)
	result := false
	if index > -1 {
		result = true
	}
	return result
}

// snake string, XxYy to xx_yy , XxYY to xx_yy
func (this *StrClass) SnakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))
}

// camel string, xx_yy to XxYy
func (this *StrClass) CamelString(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}

func (this *StrClass) StartWith(str, find string) bool {
	return strings.HasPrefix(str, find)
}
func (this *StrClass) EndWith(str, find string) bool {
	return strings.HasSuffix(str, find)
}
func (this *StrClass) StartWithi(str, find string) bool {
	return strings.HasPrefix(strings.ToLower(str), strings.ToLower(find))
}
func (this *StrClass) EndWithi(str, find string) bool {
	return strings.HasSuffix(strings.ToLower(str), strings.ToLower(find))
}
func (this *StrClass) Has(str, find string) bool {
	return strings.Contains(str, find)
}
func (this *StrClass) Hasi(str, find string) bool {
	return strings.Contains(strings.ToLower(str), strings.ToLower(find))
}
