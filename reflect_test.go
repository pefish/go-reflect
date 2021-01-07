package go_reflect

import (
	"fmt"
	"github.com/pefish/go-test-assert"
	"testing"
)

type Test struct {
	UserId             uint64  `json:"user_id"`
	Type               uint64  `json:"type"`
	OrderNumber        string  `json:"order_number"`
	Price              float64 `json:"price"`
	Amount             float64 `json:"amount"`
	TransferMemo       string  `json:"tranfer_memo"`
	Status             uint64  `json:"status"`

	BaseModel
}

type BaseModel struct {
	Id        uint64 `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func TestReflectClass_MustToInt64(t *testing.T) {
	a := `1222222`
	test.Equal(t, int64(1222222), Reflect.MustToInt64(a))

	a1 := `0x16`
	test.Equal(t, int64(22), Reflect.MustToInt64(a1))
	a2 := `0o17`
	test.Equal(t, int64(15), Reflect.MustToInt64(a2))
	a3 := `0b101`
	test.Equal(t, int64(5), Reflect.MustToInt64(a3))

	var b int = 12
	test.Equal(t, int64(12), Reflect.MustToInt64(b))

	var c int8 = 12
	test.Equal(t, int64(12), Reflect.MustToInt64(c))

	var d int16 = 12
	test.Equal(t, int64(12), Reflect.MustToInt64(d))

	var f int32 = 12
	test.Equal(t, int64(12), Reflect.MustToInt64(f))

	var g uint8 = 12
	test.Equal(t, int64(12), Reflect.MustToInt64(g))

	var h uint16 = 12
	test.Equal(t, int64(12), Reflect.MustToInt64(h))

	var i uint32 = 12
	test.Equal(t, int64(12), Reflect.MustToInt64(i))

	var m bool = true
	test.Equal(t, int64(1), Reflect.MustToInt64(m))

	m_, err := Reflect.ToInt64(m)
	test.Equal(t, nil, err)
	test.Equal(t, int64(1), m_)

}

func TestReflectClass_MustToUint64(t *testing.T) {
	a := `1222222`
	test.Equal(t, uint64(1222222), Reflect.MustToUint64(a))

	var b int = 12
	test.Equal(t, uint64(12), Reflect.MustToUint64(b))

	var c int8 = 12
	test.Equal(t, uint64(12), Reflect.MustToUint64(c))

	var d int16 = 12
	test.Equal(t, uint64(12), Reflect.MustToUint64(d))

	var f int32 = 12
	test.Equal(t, uint64(12), Reflect.MustToUint64(f))

	var g uint8 = 12
	test.Equal(t, uint64(12), Reflect.MustToUint64(g))

	var h uint16 = 12
	test.Equal(t, uint64(12), Reflect.MustToUint64(h))

	var i uint32 = 12
	test.Equal(t, uint64(12), Reflect.MustToUint64(i))

	var j uint64 = 12
	test.Equal(t, uint64(12), Reflect.MustToUint64(j))

	var k float32 = 12
	test.Equal(t, uint64(12), Reflect.MustToUint64(k))

	var l float64 = 12
	test.Equal(t, uint64(12), Reflect.MustToUint64(l))

	var m bool = true
	test.Equal(t, uint64(1), Reflect.MustToUint64(m))

	var n string = "0xc00007a000"
	test.Equal(t, uint64(824634220544), Reflect.MustToUint64(n))

	m_, err := Reflect.ToUint64(n)
	test.Equal(t, nil, err)
	test.Equal(t, uint64(824634220544), m_)
}

func TestReflectClass_ToString(t *testing.T) {
	var a *float64
	b := 0.34
	a = &b
	test.Equal(t, "0.34", Reflect.ToString(a))

	type BType struct {
		B1 int
	}
	a1 := struct {
		A string
		B BType
	}{`1`, BType{2}}
	test.Equal(t, "{1 {2}}", Reflect.ToString(a1))

	type Test struct {
		Test1 *string `json:"test1"`
	}
	test_ := Test{}
	test.Equal(t, "*nil", Reflect.ToString(test_.Test1))

	a2 := 625462456
	test.Equal(t, "625462456", Reflect.ToString(a2))

	a3 := 0xf43f2
	test.Equal(t, "1000434", Reflect.ToString(a3))
}

func TestReflectClass_GetValuesInTagFromStruct(t *testing.T) {
	// []*Test{}
	test_ := []*Test{}
	fields := Reflect.GetValuesInTagFromStruct(test_, `json`)
	test.Equal(t, "[user_id type order_number price amount tranfer_memo status id created_at updated_at]", fmt.Sprint(fields))

	// Test{}
	test1 := Test{}
	fields = Reflect.GetValuesInTagFromStruct(test1, `json`)
	test.Equal(t, "[user_id type order_number price amount tranfer_memo status id created_at updated_at]", fmt.Sprint(fields))

	// *Test{}
	test2 := Test{}
	fields = Reflect.GetValuesInTagFromStruct(&test2, `json`)
	test.Equal(t, "[user_id type order_number price amount tranfer_memo status id created_at updated_at]", fmt.Sprint(fields))

	// []Test{}
	test3 := []Test{}
	fields = Reflect.GetValuesInTagFromStruct(test3, `json`)
	test.Equal(t, "[user_id type order_number price amount tranfer_memo status id created_at updated_at]", fmt.Sprint(fields))

	// *[]Test{}
	test4 := []Test{}
	fields = Reflect.GetValuesInTagFromStruct(&test4, `json`)
	test.Equal(t, "[user_id type order_number price amount tranfer_memo status id created_at updated_at]", fmt.Sprint(fields))
}

func TestReflectClass_MustToBool(t *testing.T) {
	a := `true`
	test.Equal(t, true, Reflect.MustToBool(a))
}

func TestReflectClass_MustToInt(t *testing.T) {
	a := "4546"
	test.Equal(t, 4546, Reflect.MustToInt(a))
}

func TestReflectClass_MustToInt8(t *testing.T) {
	a := "12"
	test.Equal(t, int8(12), Reflect.MustToInt8(a))
}

func TestReflectClass_MustToInt32(t *testing.T) {
	a := "4546"
	test.Equal(t, int32(4546), Reflect.MustToInt32(a))
}

func TestReflectClass_MustToUint32(t *testing.T) {
	a := "4546"
	test.Equal(t, uint32(4546), Reflect.MustToUint32(a))
}

func TestReflectClass_MustToFloat64(t *testing.T) {
	a := "4546.3526"
	test.Equal(t, 4546.3526, Reflect.MustToFloat64(a))
}

func TestReflectClass_MustToFloat32(t *testing.T) {
	a := "4546.3526"
	test.Equal(t, float32(4546.3526), Reflect.MustToFloat32(a))
}