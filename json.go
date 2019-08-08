package fplib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func init() {
	gjson.AddModifier("default", func(json, arg string) string {
		v := Trim(json)
		if v == "" {
			v = arg
		}
		return v
	})
	gjson.AddModifier("int", func(json, arg string) string {
		v := Trim(json)
		Debug("json", json)
		Debug("arg", arg)
		i := Int64(v, Int64(arg, 0))
		Debug("i", i)
		v = String(i)
		Debug("v", v)
		return v
	})
	gjson.AddModifier("bool", func(json, arg string) string {
		v := Trim(json)
		b := Bool(v)
		if b {
			v = "true"
		} else {
			v = "false"
		}
		return v
	})
}

type JSONClass struct {
}

// gjson的parse
func (this *JSONClass) Parse(json string) gjson.Result {
	if !gjson.Valid(json) {
		return gjson.Parse("{}")
	}
	return gjson.Parse(json)
}

// 生成sjson对象
func (this *JSONClass) Obj(obj ...interface{}) *SJSON {
	s := &SJSON{}
	if len(obj) > 0 {
		s.Parse(obj[0])
	}
	return s
}

// 生成sjson数组
func (this *JSONClass) Arr(objs ...interface{}) *SJSON {
	s := &SJSON{}
	s.Parse("[]")
	for _, v := range objs {
		s.SetArr(v)
	}
	return s
}

// 输出可阅读格式
func (this *JSONClass) Show(obj interface{}) string {
	g := this.Obj(obj)
	return g.Get("@pretty").String()
}

// 输出可阅读格式
func (this *JSONClass) Debug(objs ...interface{}) {
	title := "JSON"
	for k, v := range objs {
		if k == 0 {
			if str, ok := v.(string); ok {
				title = str
				continue
			}
		}
		Debug(title, k, this.Show(v))
	}
}

// 将gjson，sjson转换为标准interface
func (this *JSONClass) ToObj(obj interface{}) interface{} {
	switch v := obj.(type) {
	case gjson.Result:
		var o interface{}
		json.Unmarshal([]byte(v.Raw), &o)
		return o
	case *gjson.Result:
		var o interface{}
		json.Unmarshal([]byte((*v).Raw), &o)
		return o
	case SJSON:
		return v.Obj()
	case *SJSON:
		return (*v).Obj()
	default:
		return obj
	}
	// if v, ok := obj.(gjson.Result); ok {
	// 	var o interface{}
	// 	json.Unmarshal([]byte(v.Raw), &o)
	// 	return o
	// } else {
	// 	return obj
	// }
}

// 标准库的序列化
func (this *JSONClass) Stringify(obj interface{}) string {
	switch v := obj.(type) {
	case gjson.Result:
		return v.Raw
	case *gjson.Result:
		return (*v).Raw
	case SJSON:
		return v.Json
	case *SJSON:
		return (*v).Json
	default:
		jsonStr, err := json.Marshal(obj)
		if err != nil {
			return "{}"
		}
		return string(jsonStr)
	}

}

// 标准库的parse，自定义结构体
func (this *JSONClass) ParseObj(str string, obj interface{}) error {
	err := json.Unmarshal([]byte(str), &obj)
	return err
}

// 标准库的parse
func (this *JSONClass) ParseJSON(str string) (interface{}, error) {
	var obj interface{}
	err := json.Unmarshal([]byte(str), &obj)
	return obj, err
}
func (this *JSONClass) ParseFile(filename string) gjson.Result {
	file, err := os.Open(filename)
	if err != nil {
		return gjson.Parse("{}")
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return gjson.Parse("{}")
	}
	return this.Parse(string(content))
}

type SJSON struct {
	Json string
}

func (this *SJSON) Obj() interface{} {
	var obj interface{}
	json.Unmarshal([]byte(this.Json), &obj)
	return obj
}
func (this *SJSON) Set(k string, p interface{}) *SJSON {
	switch v := p.(type) {
	case *SJSON:
		this.Json, _ = sjson.Set(this.Json, k, (*v).Obj())

	case SJSON:
		this.Json, _ = sjson.Set(this.Json, k, v.Obj())
	default:
		this.Json, _ = sjson.Set(this.Json, k, p)
	}
	return this
}
func (this *SJSON) SetMap(maps map[string]interface{}) *SJSON {
	for k, v := range maps {
		this.Set(k, v)
	}
	return this
}
func (this *SJSON) SetArr(p ...interface{}) *SJSON {
	if len(this.Json) < 2 {
		this.Json = "[]"
	} else {
		if this.Json[0:1] != "[" {
			this.Json = "[]"
		}
	}
	for _, v := range p {
		this.Set("-1", v)
	}
	return this
}
func (this *SJSON) Delete(k string) *SJSON {
	this.Json, _ = sjson.Delete(this.Json, k)
	return this
}

// 获取gjson对象，方便查询
func (this *SJSON) GJSON() gjson.Result {
	if !gjson.Valid(this.Json) {
		return gjson.Parse("{}")
	}
	return gjson.Parse(this.Json)
}

// 使用gjson的get方法
func (this *SJSON) Get(key string) gjson.Result {
	g := this.GJSON()
	if key == "" {
		return g
	}
	return g.Get(key)
}

func (this *SJSON) String() string {
	return this.Json
}
func (this *SJSON) Show() string {
	g := this.GJSON()
	return g.Get("@pretty").String()
}

// 从字符串，gjson，sjson，结构体获得sjson
func (this *SJSON) Parse(obj interface{}) *SJSON {
	switch v := obj.(type) {
	case string:
		jsonStr := Trim(v)
		if gjson.Valid(jsonStr) {
			this.Json = jsonStr
		} else {
			this.Json = ""
		}
	case *SJSON:
		this.Json = (*v).Json
	case SJSON:
		this.Json = v.Json
	case gjson.Result:
		this.Json = v.Raw
	case *gjson.Result:
		this.Json = (*v).Raw
	case fmt.Stringer:
		this.Json = v.String()
	default:
		this.Json = ""
		jsonStr, err := json.Marshal(obj)
		if err == nil {
			this.Json = string(jsonStr)
		}
	}
	return this
}

// 是否为空
func (this *SJSON) IsEmpty() bool {
	if this.Json == "" {
		return true
	}
	if this.Json == "{}" {
		return true
	}
	if this.Json == "[]" {
		return true
	}
	return false
}

// 是否是数组
func (this *SJSON) IsArray() bool {
	if this.Json == "[]" {
		return true
	} else if strings.HasPrefix(this.Json, "[") && strings.HasSuffix(this.Json, "]") {
		return true
	}
	return false
}

// 根据key过滤，ps[0]包含的key，为空包含所有key，ps[1]排除的key
func (this *SJSON) Filter(ps ...string) *SJSON {
	isArr := this.IsArray()

	in := make([]string, 0)
	not := make([]string, 0)

	switch len(ps) {
	case 1:
		in = Str.ToArr(ps[0])
	case 2:
		in = Str.ToArr(ps[0])
		not = Str.ToArr(ps[1])
	}

	j := this.GJSON()
	this.Json = If(isArr, "[]", "{}").(string)
	j.ForEach(func(key, value gjson.Result) bool {
		k := key.String()
		v := value.Value()

		if Str.In_Arrayi(not, k) {
			return true
		}
		if len(in) > 0 {
			if !Str.In_Arrayi(in, k) {
				return true
			}
		}
		if isArr {
			this.SetArr(v)
		} else {
			this.Set(k, v)
		}
		return true // keep iterating
	})
	return this
}

func (this *SJSON) SnakeString() *SJSON {
	if this.IsArray() {
		return this
	}

	j := this.GJSON()
	this.Json = "{}"
	j.ForEach(func(key, value gjson.Result) bool {
		k := key.String()
		v := value.Value()
		ku := Str.SnakeString(k)
		this.Set(ku, v)
		return true // keep iterating
	})
	return this
}
func (this *SJSON) CamelString() *SJSON {
	if this.IsArray() {
		return this
	}

	j := this.GJSON()
	this.Json = "{}"
	j.ForEach(func(key, value gjson.Result) bool {
		k := key.String()
		v := value.Value()
		ku := Str.CamelString(k)
		this.Set(ku, v)
		return true // keep iterating
	})
	return this
}
