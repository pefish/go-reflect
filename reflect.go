package go_reflect

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type ReflectClass struct {
}

var Reflect = ReflectClass{}


/**
通过反射获取指针中struct的所有tag值. 支持 []*Test{}、Test{}、*Test{}、[]Test{}、*[]Test{}
 */
func (reflectInstance *ReflectClass) GetValuesInTagFromStruct(interf interface{}, tag string) []string  {
	result := make([]string, 0)
	return reflectInstance.getValuesInTagFromStruct(result, reflect.TypeOf(interf), tag)
}

func (reflectInstance *ReflectClass) getValuesInTagFromStruct(result []string, type_ reflect.Type, tag string) []string {
	realValKind := type_.Kind()
	if realValKind == reflect.Ptr {
		type_ = type_.Elem()
		realValKind = type_.Kind()
		if realValKind == reflect.Slice {
			type_ = type_.Elem()
			realValKind = type_.Kind()
		}
	} else if realValKind == reflect.Slice {
		type_ = type_.Elem()
		realValKind = type_.Kind()
		if realValKind == reflect.Ptr {
			type_ = type_.Elem()
			realValKind = type_.Kind()
		}
	} else if realValKind == reflect.Struct {

	} else {
		return result
	}

	if realValKind == reflect.Struct {
		for i := 0; i < type_.NumField(); i++ {
			tagName := type_.Field(i).Tag.Get(tag)
			if tagName != `` {
				result = append(result, type_.Field(i).Tag.Get(tag))
			}
			result = reflectInstance.getValuesInTagFromStruct(result, type_.Field(i).Type, tag)
		}
	}

	return result
}

func (reflectInstance *ReflectClass) MustToInt(val interface{}) int {
	result, err := reflectInstance.ToInt(val)
	if err != nil {
		panic(err)
	}
	return result
}

func (reflectInstance *ReflectClass) ToInt(val interface{}) (int, error) {
	if val == nil {
		return 0, errors.New(`nil cannot convert to int`)
	}

	kind := reflect.TypeOf(val).Kind()
	if kind == reflect.String {
		int_, err := strconv.ParseUint(val.(string), 10, 64)
		if err != nil {
			return 0, err
		}
		return int(int_), nil
	} else if kind == reflect.Bool {
		if val.(bool) {
			return 1, nil
		} else {
			return 0, nil
		}
	} else if kind == reflect.Float32 {
		return int(val.(float32)), nil
	} else if kind == reflect.Float64 {
		return int(val.(float64)), nil
	} else if kind == reflect.Int {
		return val.(int), nil
	} else if kind == reflect.Int8 {
		return int(val.(int8)), nil
	} else if kind == reflect.Int16 {
		return int(val.(int16)), nil
	} else if kind == reflect.Int32 {
		return int(val.(int32)), nil
	} else if kind == reflect.Int64 {
		return int(val.(int64)), nil
	} else if kind == reflect.Uint {
		return int(val.(uint)), nil
	} else if kind == reflect.Uint8 {
		return int(val.(uint8)), nil
	} else if kind == reflect.Uint16 {
		return int(val.(uint16)), nil
	} else if kind == reflect.Uint32 {
		return int(val.(uint32)), nil
	} else if kind == reflect.Uint64 {
		return int(val.(uint64)), nil
	} else {
		return 0, errors.New(`convert not supported: ` + kind.String())
	}
}

func (reflectInstance *ReflectClass) MustToInt8(val interface{}) int8 {
	result, err := reflectInstance.ToInt8(val)
	if err != nil {
		panic(err)
	}
	return result
}

func (reflectInstance *ReflectClass) ToInt8(val interface{}) (int8, error) {
	if val == nil {
		return 0, errors.New(`nil cannot convert to int8`)
	}

	kind := reflect.TypeOf(val).Kind()
	if kind == reflect.String {
		int_, err := strconv.ParseUint(val.(string), 10, 64)
		if err != nil {
			return 0, err
		}
		return int8(int_), nil
	} else if kind == reflect.Bool {
		if val.(bool) {
			return 1, nil
		} else {
			return 0, nil
		}
	} else if kind == reflect.Float32 {
		return int8(val.(float32)), nil
	} else if kind == reflect.Float64 {
		return int8(val.(float64)), nil
	} else if kind == reflect.Int {
		return int8(val.(int)), nil
	} else if kind == reflect.Int8 {
		return val.(int8), nil
	} else if kind == reflect.Int16 {
		return int8(val.(int16)), nil
	} else if kind == reflect.Int32 {
		return int8(val.(int32)), nil
	} else if kind == reflect.Int64 {
		return int8(val.(int64)), nil
	} else if kind == reflect.Uint {
		return int8(val.(uint)), nil
	} else if kind == reflect.Uint8 {
		return int8(val.(uint8)), nil
	} else if kind == reflect.Uint16 {
		return int8(val.(uint16)), nil
	} else if kind == reflect.Uint32 {
		return int8(val.(uint32)), nil
	} else if kind == reflect.Uint64 {
		return int8(val.(uint64)), nil
	} else {
		return 0, errors.New(`convert not supported: ` + kind.String())
	}
}

func (reflectInstance *ReflectClass) MustToBool(val interface{}) bool {
	result, err := reflectInstance.ToBool(val)
	if err != nil {
		panic(err)
	}
	return result
}

func (reflectInstance *ReflectClass) ToBool(val interface{}) (bool, error) {
	if val == nil {
		return false, errors.New(`nil cannot convert to bool`)
	}

	kind := reflect.TypeOf(val).Kind()
	if kind == reflect.String {
		bool_, err := strconv.ParseBool(val.(string))
		if err != nil {
			return false, err
		}
		return bool_, nil
	} else if kind == reflect.Bool {
		return val.(bool), nil
	} else {
		return false, errors.New(`convert not supported: ` + kind.String())
	}
}

func (reflectInstance *ReflectClass) MustToInt32(val interface{}) int32 {
	result, err := reflectInstance.ToInt32(val)
	if err != nil {
		panic(err)
	}
	return result
}

func (reflectInstance *ReflectClass) ToInt32(val interface{}) (int32, error) {
	if val == nil {
		return 0, errors.New(`nil cannot convert to int32`)
	}

	kind := reflect.TypeOf(val).Kind()
	if kind == reflect.String {
		int_, err := strconv.ParseInt(val.(string), 10, 64)
		if err != nil {
			return 0, err
		}
		return int32(int_), nil
	} else if kind == reflect.Bool {
		if val.(bool) {
			return 1, nil
		} else {
			return 0, nil
		}
	} else if kind == reflect.Float32 {
		return int32(val.(float32)), nil
	} else if kind == reflect.Float64 {
		return int32(val.(float64)), nil
	} else if kind == reflect.Int {
		return int32(val.(int)), nil
	} else if kind == reflect.Int8 {
		return int32(val.(int8)), nil
	} else if kind == reflect.Int16 {
		return int32(val.(int16)), nil
	} else if kind == reflect.Int32 {
		return val.(int32), nil
	} else if kind == reflect.Int64 {
		return int32(val.(int64)), nil
	} else if kind == reflect.Uint {
		return int32(val.(uint)), nil
	} else if kind == reflect.Uint8 {
		return int32(val.(uint8)), nil
	} else if kind == reflect.Uint16 {
		return int32(val.(uint16)), nil
	} else if kind == reflect.Uint32 {
		return int32(val.(uint32)), nil
	} else if kind == reflect.Uint64 {
		return int32(val.(uint64)), nil
	} else {
		return 0, errors.New(`convert not supported: ` + kind.String())
	}
}

func (reflectInstance *ReflectClass) MustToInt64(val interface{}) int64 {
	result, err := reflectInstance.ToInt64(val)
	if err != nil {
		panic(err)
	}
	return result
}

func (reflectInstance *ReflectClass) ToInt64(val interface{}) (int64, error) {
	if val == nil {
		return 0, errors.New(`nil cannot convert to int64`)
	}

	kind := reflect.TypeOf(val).Kind()
	if kind == reflect.String {
		int_, err := strconv.ParseInt(val.(string), 10, 64)
		if err != nil {
			return 0, err
		}
		return int_, nil
	} else if kind == reflect.Bool {
		if val.(bool) {
			return 1, nil
		} else {
			return 0, nil
		}
	} else if kind == reflect.Float32 {
		return int64(val.(float32)), nil
	} else if kind == reflect.Float64 {
		return int64(val.(float64)), nil
	} else if kind == reflect.Int {
		return int64(val.(int)), nil
	} else if kind == reflect.Int8 {
		return int64(val.(int8)), nil
	} else if kind == reflect.Int16 {
		return int64(val.(int16)), nil
	} else if kind == reflect.Int32 {
		return int64(val.(int32)), nil
	} else if kind == reflect.Int64 {
		return val.(int64), nil
	} else if kind == reflect.Uint {
		return int64(val.(uint)), nil
	} else if kind == reflect.Uint8 {
		return int64(val.(uint8)), nil
	} else if kind == reflect.Uint16 {
		return int64(val.(uint16)), nil
	} else if kind == reflect.Uint32 {
		return int64(val.(uint32)), nil
	} else if kind == reflect.Uint64 {
		return int64(val.(uint64)), nil
	} else {
		return 0, errors.New(`convert not supported: ` + kind.String())
	}
}


func (reflectInstance *ReflectClass) MustToUint64(val interface{}) uint64 {
	result, err := reflectInstance.ToUint64(val)
	if err != nil {
		panic(err)
	}
	return result
}

func (reflectInstance *ReflectClass) ToUint64(val interface{}) (uint64, error) {
	if val == nil {
		return 0, errors.New(`nil cannot convert to uint64`)
	}

	kind := reflect.TypeOf(val).Kind()
	if kind == reflect.String {
		base := 10
		str := val.(string)
		if strings.HasPrefix(str, "0x") {
			base = 16
			str = str[2:]
		}
		int_, err := strconv.ParseUint(str, base, 64)
		if err != nil {
			return 0, err
		}
		return int_, nil
	} else if kind == reflect.Bool {
		if val.(bool) {
			return 1, nil
		} else {
			return 0, nil
		}
	} else if kind == reflect.Float32 {
		return uint64(val.(float32)), nil
	} else if kind == reflect.Float64 {
		return uint64(val.(float64)), nil
	} else if kind == reflect.Int {
		return uint64(val.(int)), nil
	} else if kind == reflect.Int8 {
		return uint64(val.(int8)), nil
	} else if kind == reflect.Int16 {
		return uint64(val.(int16)), nil
	} else if kind == reflect.Int32 {
		return uint64(val.(int32)), nil
	} else if kind == reflect.Int64 {
		return uint64(val.(int64)), nil
	} else if kind == reflect.Uint {
		return uint64(val.(uint)), nil
	} else if kind == reflect.Uint8 {
		return uint64(val.(uint8)), nil
	} else if kind == reflect.Uint16 {
		return uint64(val.(uint16)), nil
	} else if kind == reflect.Uint32 {
		return uint64(val.(uint32)), nil
	} else if kind == reflect.Uint64 {
		return val.(uint64), nil
	} else {
		return 0, errors.New(`convert not supported: ` + kind.String())
	}
}

func (reflectInstance *ReflectClass) MustToUint32(val interface{}) uint32 {
	result, err := reflectInstance.ToUint32(val)
	if err != nil {
		panic(err)
	}
	return result
}

func (reflectInstance *ReflectClass) ToUint32(val interface{}) (uint32, error) {
	if val == nil {
		return 0, errors.New(`nil cannot convert to uint32`)
	}

	kind := reflect.TypeOf(val).Kind()
	if kind == reflect.String {
		int_, err := strconv.ParseUint(val.(string), 10, 64)
		if err != nil {
			return 0, err
		}
		return uint32(int_), nil
	} else if kind == reflect.Bool {
		if val.(bool) {
			return 1, nil
		} else {
			return 0, nil
		}
	} else if kind == reflect.Float32 {
		return uint32(val.(float32)), nil
	} else if kind == reflect.Float64 {
		return uint32(val.(float64)), nil
	} else if kind == reflect.Int {
		return uint32(val.(int)), nil
	} else if kind == reflect.Int8 {
		return uint32(val.(int8)), nil
	} else if kind == reflect.Int16 {
		return uint32(val.(int16)), nil
	} else if kind == reflect.Int32 {
		return uint32(val.(int32)), nil
	} else if kind == reflect.Int64 {
		return uint32(val.(int64)), nil
	} else if kind == reflect.Uint {
		return uint32(val.(uint)), nil
	} else if kind == reflect.Uint8 {
		return uint32(val.(uint8)), nil
	} else if kind == reflect.Uint16 {
		return uint32(val.(uint16)), nil
	} else if kind == reflect.Uint32 {
		return val.(uint32), nil
	} else if kind == reflect.Uint64 {
		return uint32(val.(uint64)), nil
	} else {
		return 0, errors.New(`convert not supported: ` + kind.String())
	}
}

func (reflectInstance *ReflectClass) MustToFloat64(val interface{}) float64 {
	result, err := reflectInstance.ToFloat64(val)
	if err != nil {
		panic(err)
	}
	return result
}

func (reflectInstance *ReflectClass) ToFloat64(val interface{}) (float64, error) {
	if val == nil {
		return 0, errors.New(`nil cannot convert to float64`)
	}

	kind := reflect.TypeOf(val).Kind()
	if kind == reflect.String {
		result, err := strconv.ParseFloat(val.(string), 64)
		if err != nil {
			return 0, err
		}
		return result, nil
	} else if kind == reflect.Bool {
		if val.(bool) {
			return 1, nil
		} else {
			return 0, nil
		}
	} else if kind == reflect.Float32 {
		return float64(val.(float32)), nil
	} else if kind == reflect.Float64 {
		return val.(float64), nil
	} else if kind == reflect.Int {
		return float64(val.(int)), nil
	} else if kind == reflect.Int8 {
		return float64(val.(int8)), nil
	} else if kind == reflect.Int16 {
		return float64(val.(int16)), nil
	} else if kind == reflect.Int32 {
		return float64(val.(int32)), nil
	} else if kind == reflect.Int64 {
		return float64(val.(int64)), nil
	} else if kind == reflect.Uint {
		return float64(val.(uint)), nil
	} else if kind == reflect.Uint8 {
		return float64(val.(uint8)), nil
	} else if kind == reflect.Uint16 {
		return float64(val.(uint16)), nil
	} else if kind == reflect.Uint32 {
		return float64(val.(uint32)), nil
	} else if kind == reflect.Uint64 {
		return float64(val.(uint64)), nil
	} else {
		return 0, errors.New(`convert not supported: ` + kind.String())
	}
}

func (reflectInstance *ReflectClass) MustToFloat32(val interface{}) float32 {
	result, err := reflectInstance.ToFloat32(val)
	if err != nil {
		panic(err)
	}
	return result
}

func (reflectInstance *ReflectClass) ToFloat32(val interface{}) (float32, error) {
	if val == nil {
		return 0, errors.New(`nil cannot convert to float32`)
	}

	kind := reflect.TypeOf(val).Kind()
	if kind == reflect.String {
		result, err := strconv.ParseFloat(val.(string), 64)
		if err != nil {
			return 0, err
		}
		return float32(result), nil
	} else if kind == reflect.Bool {
		if val.(bool) {
			return 1, nil
		} else {
			return 0, nil
		}
	} else if kind == reflect.Float32 {
		return val.(float32), nil
	} else if kind == reflect.Float64 {
		return float32(val.(float64)), nil
	} else if kind == reflect.Int {
		return float32(val.(int)), nil
	} else if kind == reflect.Int8 {
		return float32(val.(int8)), nil
	} else if kind == reflect.Int16 {
		return float32(val.(int16)), nil
	} else if kind == reflect.Int32 {
		return float32(val.(int32)), nil
	} else if kind == reflect.Int64 {
		return float32(val.(int64)), nil
	} else if kind == reflect.Uint {
		return float32(val.(uint)), nil
	} else if kind == reflect.Uint8 {
		return float32(val.(uint8)), nil
	} else if kind == reflect.Uint16 {
		return float32(val.(uint16)), nil
	} else if kind == reflect.Uint32 {
		return float32(val.(uint32)), nil
	} else if kind == reflect.Uint64 {
		return float32(val.(uint64)), nil
	} else {
		return 0, errors.New(`convert not supported: ` + kind.String())
	}
}

func (reflectInstance *ReflectClass) ToString(val interface{}) string {
	if val == nil {
		return `nil`
	}
	type_ := reflect.TypeOf(val)
	kind := type_.Kind()
	typeStr_ := type_.String()
	if kind == reflect.String {
		return val.(string)
	} else if kind == reflect.Bool {
		return strconv.FormatBool(val.(bool))
	} else if kind == reflect.Float32 {
		return strconv.FormatFloat(float64(val.(float32)), 'f', -1, 64)
	} else if kind == reflect.Float64 {
		return strconv.FormatFloat(val.(float64), 'f', -1, 64)
	} else if kind == reflect.Int {
		return strconv.FormatInt(int64(val.(int)), 10)
	} else if kind == reflect.Int8 {
		return strconv.FormatInt(int64(val.(int8)), 10)
	} else if kind == reflect.Int16 {
		return strconv.FormatInt(int64(val.(int16)), 10)
	} else if kind == reflect.Int32 {
		return strconv.FormatInt(int64(val.(int32)), 10)
	} else if kind == reflect.Int64 {
		if typeStr_ == `time.Duration` {
			return strconv.FormatInt(int64(val.(time.Duration)), 10)
		}
		return strconv.FormatInt(val.(int64), 10)
	} else if kind == reflect.Uint {
		return strconv.FormatUint(uint64(val.(uint)), 10)
	} else if kind == reflect.Uint8 {
		return strconv.FormatInt(int64(val.(uint8)), 10)
	} else if kind == reflect.Uint16 {
		return strconv.FormatUint(uint64(val.(uint16)), 10)
	} else if kind == reflect.Uint32 {
		return strconv.FormatUint(uint64(val.(uint32)), 10)
	} else if kind == reflect.Uint64 {
		return strconv.FormatUint(val.(uint64), 10)
	} else if kind == reflect.Ptr {
		reflectVal := reflect.ValueOf(val)
		if reflectVal.IsNil() { // IsNil 只接受 chan, func, interface, map, pointer, or slice value
			return `*nil`
		}
		return reflectInstance.ToString(reflectVal.Elem().Interface())
	} else {
		return fmt.Sprint(val)
	}
}
