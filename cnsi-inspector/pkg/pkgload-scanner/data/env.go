package data

import (
	"os"
)

type EntityEnv struct {
	name     string
	value    string
	_default string
}

func Env(envName string, _default string) *EntityEnv {
	e := new(EntityEnv)
	e.name = envName
	e._default = _default
	e.value = os.Getenv(envName)
	//if e.value == "" {
	//}
	return e
}

func (e *EntityEnv) Float() float32 {
	if e.value == "" {
		return String().Float(e._default)
	}
	return String().Float(e.value)
}

func (e *EntityEnv) Int() int {
	if e.value == "" {
		return String().Int(e._default)
	}
	return String().Int(e.value)
}

func (e *EntityEnv) Int64() int64 {
	if e.value == "" {
		return String().Int64(e._default)
	}
	return String().Int64(e.value)
}

func (e *EntityEnv) Uint64() uint64 {
	if e.value == "" {
		return String().Uint64(e._default)
	}
	return String().Uint64(e.value)
}

func (e *EntityEnv) String() string {

	if e.value == "" {
		return e._default
	}
	return e.value
}

func (e *EntityEnv) Bool() bool {
	if e.value == "" {
		return String().Bool(e._default)
	}
	return String().Bool(e.value)
}
