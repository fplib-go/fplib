package fplib

type Conf struct {
	_ConfData map[string]interface{}
}

func (this *Conf) Set(key string, value interface{}) *Conf {
	this._ConfData[key] = value
	return this
}
func (this *Conf) Get(key string) interface{} {
	return this._ConfData[key]
}
func (this *Conf) GetString(key string, defaults ...string) string {
	v := this.Get(key)
	var result string
	if len(defaults) > 0 {
		result = defaults[0]
	}
	if v != nil {
		if val, ok := v.(string); ok {
			result = val
		}
	}
	return result
}

func (this *Conf) GetInt(key string, defaults ...int) int {
	v := this.Get(key)
	var result int
	if len(defaults) > 0 {
		result = defaults[0]
	}
	if v != nil {
		if val, ok := v.(int); ok {
			result = val
		}
	}
	return result
}
func (this *Conf) GetBool(key string, defaults ...bool) bool {
	v := this.Get(key)
	var result bool
	if len(defaults) > 0 {
		result = defaults[0]
	}
	if v != nil {
		if val, ok := v.(bool); ok {
			result = val
		} else if val, ok := v.(string); ok {
			result = Bool(val)
		}
	}
	return result
}
