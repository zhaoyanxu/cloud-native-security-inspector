package data

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var _entityString *entityString

func init() {
	_entityString = new(entityString)
}

type entityString struct {
}

func String() *entityString {
	return _entityString
}

//直接fmt就行
//func (str *entity_string) String(i int) string {
//	val := strconv.Itoa(v)
//	return val
//}

func (str *entityString) Float(v string) float32 {
	val, _ := strconv.ParseFloat(v, 32)
	return float32(val)
}

func (str *entityString) Json(v interface{}) string {
	con, err := json.Marshal(v)

	if err != nil {
		return ""
	}

	return string(con)
}

func (str *entityString) AutoInt(v string) int {

	re := regexp.MustCompile("[^0-9]")
	v = re.ReplaceAllString(v, "")

	//fmt.Println("====", v)
	val, _ := strconv.Atoi(v)
	//fmt.Println("====||||", val)
	return val
}

func (str *entityString) Int(v string) int {
	val, _ := strconv.Atoi(v)
	return val
}

func (str *entityString) Int64(v string) int64 {
	val, _ := strconv.ParseInt(v, 10, 64)
	return val
}

func (str *entityString) Uint(v string) uint {
	val, _ := strconv.ParseUint(v, 10, 64)
	return uint(val)
}
func (str *entityString) Uint64(v string) uint64 {
	val, _ := strconv.ParseUint(v, 10, 64)
	return val
}

func (str *entityString) Int8(v string) int8 {
	val, _ := strconv.ParseInt(v, 10, 64)
	return int8(val)
}

func (str *entityString) Int32(v string) int32 {
	val, _ := strconv.ParseInt(v, 10, 64)
	return int32(val)
}

func (str *entityString) Bool(v string) bool {
	switch v {
	case "":
		fallthrough
	case "0":
		fallthrough
	case "false":
		fallthrough
	case "FALSE":
		return false

	case "true":
		fallthrough
	case "1":
		fallthrough
	case "TRUE":
		return true
	}
	return true
}

// 分割字符串
func Split(str string, sep ...string) (res []string) {
	if len(sep) > 0 {
		res = strings.Split(str, sep[0])
	} else {
		res = strings.Split(str, ",")
	}
	return res
}

func IsNumeric(str string) bool {

	// 0-9 = 48 - 57
	for _, v := range str {
		if v > 57 || v < 48 {
			return false
		}
	}
	return true
}

func Join(str []string, sep ...string) (res string) {
	if len(sep) > 0 {
		res = strings.Join(str, sep[0])
	} else {
		res = strings.Join(str, ",")
	}
	return res
}

func Lower(str string) string {
	return strings.ToLower(str)
}

func Upper(str string) string {
	return strings.ToUpper(str)
}

// 转成字符串
// 必须为 slice 类型
// 其他类型请使用其他函数
//
//	use : JoinSlice([]int{1, 2, 3, 4, 5}, ",")
func JoinSlice(v interface{}, sep string) string {
	return strings.Trim(strings.Join(strings.Split(fmt.Sprint(v), " "), sep), "[]")
}
