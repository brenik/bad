package service

import (
	"reflect"
)

func ShallowCopyMap(src interface{}) interface{} {
	srcVal := reflect.ValueOf(src)
	dstVal := reflect.MakeMap(srcVal.Type())

	for _, key := range srcVal.MapKeys() {
		dstVal.SetMapIndex(key, srcVal.MapIndex(key))
	}

	return dstVal.Interface()
}

func CopyMap(m map[int]int) map[int]int {
	cp := make(map[int]int)
	for k, v := range m {
		cp[k] = v
	}
	return cp
}
