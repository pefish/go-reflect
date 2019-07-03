package p_reflect

import (
	"encoding/json"
	"errors"
	"reflect"
	"runtime"
	"strconv"
	"time"
)

type ReflectClass struct {
}

var Reflect = ReflectClass{}


/**
通过反射获取指针中struct的所有tag值. 支持 []*Test{}、Test{}、*Test{}、[]Test{}、*[]Test{}
 */
func (this *ReflectClass) GetValuesInTagFromStruct(interf interface{}, tag string) []string  {
	result := []string{}
	return this.getValuesInTagFromStruct(result, reflect.TypeOf(interf), tag)
}

func (this *ReflectClass) getValuesInTagFromStruct(result []string, type_ reflect.Type, tag string) []string {
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
			result = this.getValuesInTagFromStruct(result, type_.Field(i).Type, tag)
		}
	}

	return result
}

func (this *ReflectClass) ToInt(val interface{}) int {
	kind := reflect.TypeOf(val).Kind()
	if kind == reflect.String {
		int_, err := strconv.ParseUint(val.(string), 10, 64)
		if err != nil {
			panic(err)
		}
		return int(int_)
	} else if kind == reflect.Bool {
		if val.(bool) {
			return int(1)
		} else {
			return int(0)
		}
	} else if kind == reflect.Float32 {
		return int(val.(float32))
	} else if kind == reflect.Float64 {
		return int(val.(float64))
	} else if kind == reflect.Int {
		return val.(int)
	} else if kind == reflect.Int8 {
		return int(val.(int8))
	} else if kind == reflect.Int16 {
		return int(val.(int16))
	} else if kind == reflect.Int32 {
		return int(val.(int32))
	} else if kind == reflect.Int64 {
		return int(val.(int64))
	} else if kind == reflect.Uint {
		return int(val.(uint))
	} else if kind == reflect.Uint8 {
		return int(val.(uint8))
	} else if kind == reflect.Uint16 {
		return int(val.(uint16))
	} else if kind == reflect.Uint32 {
		return int(val.(uint32))
	} else if kind == reflect.Uint64 {
		return int(val.(uint64))
	} else {
		panic(errors.New(`convert not supported: ` + kind.String()))
	}
}

func (this *ReflectClass) ToInt32(val interface{}) int32 {
	kind := reflect.TypeOf(val).Kind()
	if kind == reflect.String {
		int_, err := strconv.ParseInt(val.(string), 10, 64)
		if err != nil {
			panic(err)
		}
		return int32(int_)
	} else if kind == reflect.Bool {
		if val.(bool) {
			return int32(1)
		} else {
			return int32(0)
		}
	} else if kind == reflect.Float32 {
		return int32(val.(float32))
	} else if kind == reflect.Float64 {
		return int32(val.(float64))
	} else if kind == reflect.Int {
		return int32(val.(int))
	} else if kind == reflect.Int8 {
		return int32(val.(int8))
	} else if kind == reflect.Int16 {
		return int32(val.(int16))
	} else if kind == reflect.Int32 {
		return val.(int32)
	} else if kind == reflect.Int64 {
		return int32(val.(int64))
	} else if kind == reflect.Uint {
		return int32(val.(uint))
	} else if kind == reflect.Uint8 {
		return int32(val.(uint8))
	} else if kind == reflect.Uint16 {
		return int32(val.(uint16))
	} else if kind == reflect.Uint32 {
		return int32(val.(uint32))
	} else if kind == reflect.Uint64 {
		return int32(val.(uint64))
	} else {
		panic(errors.New(`convert not supported: ` + kind.String()))
	}
}

func (this *ReflectClass) ToInt64(val interface{}) int64 {
	kind := reflect.TypeOf(val).Kind()
	if kind == reflect.String {
		int_, err := strconv.ParseInt(val.(string), 10, 64)
		if err != nil {
			panic(err)
		}
		return int_
	} else if kind == reflect.Bool {
		if val.(bool) {
			return int64(1)
		} else {
			return int64(0)
		}
	} else if kind == reflect.Float32 {
		return int64(val.(float32))
	} else if kind == reflect.Float64 {
		return int64(val.(float64))
	} else if kind == reflect.Int {
		return int64(val.(int))
	} else if kind == reflect.Int8 {
		return int64(val.(int8))
	} else if kind == reflect.Int16 {
		return int64(val.(int16))
	} else if kind == reflect.Int32 {
		return int64(val.(int32))
	} else if kind == reflect.Int64 {
		return val.(int64)
	} else if kind == reflect.Uint {
		return int64(val.(uint))
	} else if kind == reflect.Uint8 {
		return int64(val.(uint8))
	} else if kind == reflect.Uint16 {
		return int64(val.(uint16))
	} else if kind == reflect.Uint32 {
		return int64(val.(uint32))
	} else if kind == reflect.Uint64 {
		return int64(val.(uint64))
	} else {
		panic(errors.New(`convert not supported: ` + kind.String()))
	}
}

func (this *ReflectClass) ToUint64(val interface{}) uint64 {
	kind := reflect.TypeOf(val).Kind()
	if kind == reflect.String {
		int_, err := strconv.ParseUint(val.(string), 10, 64)
		if err != nil {
			panic(err)
		}
		return int_
	} else if kind == reflect.Bool {
		if val.(bool) {
			return uint64(1)
		} else {
			return uint64(0)
		}
	} else if kind == reflect.Float32 {
		return uint64(val.(float32))
	} else if kind == reflect.Float64 {
		return uint64(val.(float64))
	} else if kind == reflect.Int {
		return uint64(val.(int))
	} else if kind == reflect.Int8 {
		return uint64(val.(int8))
	} else if kind == reflect.Int16 {
		return uint64(val.(int16))
	} else if kind == reflect.Int32 {
		return uint64(val.(int32))
	} else if kind == reflect.Int64 {
		return uint64(val.(int64))
	} else if kind == reflect.Uint {
		return uint64(val.(uint))
	} else if kind == reflect.Uint8 {
		return uint64(val.(uint8))
	} else if kind == reflect.Uint16 {
		return uint64(val.(uint16))
	} else if kind == reflect.Uint32 {
		return uint64(val.(uint32))
	} else if kind == reflect.Uint64 {
		return val.(uint64)
	} else {
		panic(errors.New(`convert not supported: ` + kind.String()))
	}
}

func (this *ReflectClass) ToString(val interface{}) string {
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
		return this.ToString(reflectVal.Elem().Interface())
	} else {
		if typeStr_ == `runtime.TypeAssertionError` {
			err := val.(runtime.TypeAssertionError)
			return err.Error()
		}
		result, err := json.Marshal(val)
		if err != nil {
			panic(err)
		}
		return string(result)
	}
}
