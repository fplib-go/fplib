package fplib

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	// "labix.org/v2/mgo/bson"
)

// 获取变量的类型名称
func Typeof(v ...interface{}) string {
	msg := make([]string, len(v))
	for index, value := range v {
		if value == nil {
			msg[index] = "nil"
		} else {
			msg[index] = reflect.TypeOf(value).String()

		}
	}
	return strings.Join(msg, ",")
}

// 深拷贝
func DeepCopy(value interface{}) interface{} {
	if valueMap, ok := value.(map[string]interface{}); ok {
		newMap := make(map[string]interface{})
		for k, v := range valueMap {
			newMap[k] = DeepCopy(v)
		}

		return newMap
	} else if valueSlice, ok := value.([]interface{}); ok {
		newSlice := make([]interface{}, len(valueSlice))
		for k, v := range valueSlice {
			newSlice[k] = DeepCopy(v)
		}

		return newSlice
	}
	//  else if valueMap, ok := value.(bson.M); ok {
	// 	newMap := make(bson.M)
	// 	for k, v := range valueMap {
	// 		newMap[k] = DeepCopy(v)
	// 	}
	// }
	return value
}

// 判断变量是否为空
func Empty(obj interface{}) bool {
	canUse := false
	switch v := obj.(type) {
	case nil:
		canUse = false
	case int:
		if v != 0 {
			canUse = true
		}
	case int8:
		if v != 0 {
			canUse = true
		}
	case int16:
		if v != 0 {
			canUse = true
		}
	case int32:
		if v != 0 {
			canUse = true
		}
	case int64:
		if v != 0 {
			canUse = true
		}
	case uint:
		if v != 0 {
			canUse = true
		}
	case uint8:
		if v != 0 {
			canUse = true
		}
	case uint16:
		if v != 0 {
			canUse = true
		}
	case uint32:
		if v != 0 {
			canUse = true
		}
	case uint64:
		if v != 0 {
			canUse = true
		}
	case string:
		val := Trim(v)
		if len(val) > 0 {
			canUse = true
		}
	case fmt.Stringer:
		vv := v.String()
		vv = Trim(vv)
		if len(vv) > 0 {
			canUse = true
		}
	default:
		canUse = reflect.ValueOf(v).IsValid()
		if canUse {
			canUse = !IsBlank(reflect.ValueOf(v))
		}
	}
	return !canUse
}

func IsBlank(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String:
		return value.Len() == 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return value.IsNil()
	}
	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}

// 判断是否是布尔真值，支持字符串的true和false等写法
func Bool(obj interface{}) bool {
	result := false
	switch v := obj.(type) {
	case nil:
		result = false
	case int:
		if v != 0 {
			result = true
		}
	case int8:
		if v != 0 {
			result = true
		}
	case int16:
		if v != 0 {
			result = true
		}
	case int32:
		if v != 0 {
			result = true
		}
	case int64:
		if v != 0 {
			result = true
		}
	case float32:
		if v != 0 {
			result = true
		}
	case float64:
		if v != 0 {
			result = true
		}
	case bool:
		result = v
	case string:
		v = Trim(v)
		if len(v) > 0 {
			vv := strings.ToLower(v)
			if vv == "on" || vv == "open" || vv == "pass" {
				result = true
			} else if r, err := strconv.ParseBool(v); err == nil {
				result = r
			}
		}

	default:
		result = reflect.ValueOf(obj).IsValid()
	}
	return result
}

// 转换为int型
func Int(obj interface{}, defaults ...int) int {
	result := 0
	if len(defaults) > 0 {
		result = defaults[0]
	}
	if !Empty(obj) {
		switch v := obj.(type) {
		case string:
			if i, err := strconv.Atoi(v); err == nil {
				result = i
			}
		default:
			if i, ok := v.(int); ok {
				result = i
			}
		}
	}
	return result
}

// 转换为int型
func Int64(obj interface{}, defaults ...int64) int64 {
	result := int64(0)
	if len(defaults) > 0 {
		result = defaults[0]
	}
	if !Empty(obj) {
		switch v := obj.(type) {
		case string:
			if i, err := strconv.ParseInt(v, 10, 64); err == nil {
				result = i
			}
		default:
			if i, ok := v.(int64); ok {
				result = i
			}
		}
	}
	return result
}

// 转换为string型
func String(obj interface{}, defaults ...string) string {
	result := ""
	if len(defaults) > 0 {
		result = defaults[0]
	}
	if !Empty(obj) {
		switch v := obj.(type) {
		case int:
			result = strconv.Itoa(v)
		case int64:
			result = strconv.FormatInt(v, 10)
		case float32:
			result = strconv.FormatFloat(float64(v), 'E', -1, 32)
		case float64:
			result = strconv.FormatFloat(v, 'E', -1, 64)
		case int8, int16, int32, uint, uint8, uint16, uint32, uint64:
			if i, ok := obj.(int64); ok {
				result = strconv.FormatInt(i, 10)
			}
		default:
			if i, ok := v.(string); ok {
				result = i
			}
		}
	}
	return result
}

func Float32(obj interface{}, defaults ...float32) float32 {
	result := float32(0)
	if len(defaults) > 0 {
		result = defaults[0]
	}
	if !Empty(obj) {
		switch v := obj.(type) {
		case string:
			if s, err := strconv.ParseFloat(v, 32); err == nil {
				result = float32(s)
			}
		default:
			if i, ok := v.(float32); ok {
				result = i
			}
		}
	}
	return result
}
func Float64(obj interface{}, defaults ...float64) float64 {
	result := float64(0)
	if len(defaults) > 0 {
		result = defaults[0]
	}
	if !Empty(obj) {
		switch v := obj.(type) {
		case string:
			if s, err := strconv.ParseFloat(v, 32); err == nil {
				result = s
			}
		default:
			if i, ok := v.(float64); ok {
				result = i
			}
		}
	}
	return result
}

// 支持多字符连续处理的trim
func Trim(strs ...string) string {
	result := ""
	l := len(strs)
	c := make([]string, 0)
	switch {
	case l == 1:
		result = strs[0]
		c = append(c, " ")
		c = append(c, "\t")
		c = append(c, "\n")
		c = append(c, "\r\n")
	case (l >= 1):
		result = strs[0]
		for index := 1; index < l; index++ {
			c = append(c, strs[index])
		}
	default:
		return result
	}
	oldl := 0
	for len(result) != oldl {
		oldl = len(result)
		for _, v := range c {
			result = strings.Trim(result, v)
		}
	}
	return result
}

// 如果变量为空，设置变量的默认值
func Default(obj interface{}, def interface{}) interface{} {
	return If(Empty(obj), def, obj)
}

// 三元运算符hack
func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

// 反射方式获取变量是否在数组中，如果存在返回索引，不存在返回-1
func In_Array_Index(array interface{}, val interface{}) int {
	index := -1
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		{
			s := reflect.ValueOf(array)
			for i := 0; i < s.Len(); i++ {
				if reflect.DeepEqual(val, s.Index(i).Interface()) {
					index = i
					return index
				}
			}
		}
	}
	return index
}

// 反射方式获取变量是否在数组中
func In_Array(array interface{}, val interface{}) bool {
	index := In_Array_Index(array, val)
	result := false
	if index > -1 {
		result = true
	}
	return result
}
