import (
	"errors"
	"reflect"
	"strconv"
)

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
		value, ok := configMap[tag]
		if !ok {
			continue
		}

		switch v.Field(i).Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			res, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return err
			}
			v.Field(i).SetInt(res)
		case reflect.String:
			v.Field(i).SetString(value)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			res, err := strconv.ParseUint(value, 10, 64)
			if err != nil {
				return err
			}
			v.Field(i).SetUint(res)
		case reflect.Float32:
			res, err := strconv.ParseFloat(value, 32)
			if err != nil {
				return err
			}
			v.Field(i).SetFloat(res)
		case reflect.Float64:
			res, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return err
			}
			v.Field(i).SetFloat(res)
		case reflect.Slice:
			var strArray []string
			var valArray []reflect.Value
			var valArr reflect.Value
			elemKind := v.Field(i).Type().Elem().Kind()
			elemType := t.Field(i).Type.Elem()

			value = strings.Trim(strings.Trim(value, "["), "]")
			strArray = strings.Split(value, ",")

			switch elemKind {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				for _, e := range strArray {
					ee, err := strconv.ParseInt(e, 10, 64)
					if err != nil {
						return err
					}
					valArray = append(valArray, reflect.ValueOf(ee).Convert(elemType))
				}
			case reflect.String:
				for _, e := range strArray {
					valArray = append(valArray, reflect.ValueOf(e).Convert(elemType))
				}
			}

			valArr = reflect.Append(v.Field(i), valArray...)
			v.Field(i).Set(valArr)
		}
	}

	return nil
}
