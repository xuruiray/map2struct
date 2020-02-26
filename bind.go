import (
	"errors"
	"reflect"
	"strconv"
)

// bindConfig 仅能解析一层配置
func bindConfig(configMap map[string]string, result interface{}) error {

	if reflect.ValueOf(result).Kind() != reflect.Ptr {
		return errors.New("input not point")
	}

	if reflect.ValueOf(result).IsNil() {
		return errors.New("input is null")
	}

	v := reflect.ValueOf(result).Elem()
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag.Get("json")
		switch v.Field(i).Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			res, err := strconv.ParseInt(configMap[tag], 10, 64)
			if err != nil {
				return err
			}
			v.Field(i).SetInt(res)
		case reflect.String:
			v.Field(i).SetString(configMap[tag])
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			res, err := strconv.ParseUint(configMap[tag], 10, 64)
			if err != nil {
				return err
			}
			v.Field(i).SetUint(res)
		case reflect.Float32:
			res, err := strconv.ParseFloat(configMap[tag], 32)
			if err != nil {
				return err
			}
			v.Field(i).SetFloat(res)
		case reflect.Float64:
			res, err := strconv.ParseFloat(configMap[tag], 64)
			if err != nil {
				return err
			}
			v.Field(i).SetFloat(res)
		}
	}

	return nil
}
