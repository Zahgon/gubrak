package gubrak

import (
	"reflect"
)

func inspectFunc(err *error, data any) (reflect.Value, reflect.Type) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), *new(reflect.Type)
}

func inspectData(data any) (reflect.Value, reflect.Type, reflect.Kind, int) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), *new(reflect.Type), *new(reflect.Kind), 0
}

func makeSlice(valueType reflect.Type, args ...int) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

func validateFuncInputForSliceLoop(err *error, funcType reflect.Type, data reflect.Value) int {
	_ = "STUB: not implemented"
	return 0
}

func validateFuncInputForSliceLoopWithoutIndex(err *error, funcType reflect.Type, data reflect.Value) {
	_ = "STUB: not implemented"
	return
}

func validateFuncInputForCollectionLoop(err *error, funcType reflect.Type, data reflect.Value) int {
	_ = "STUB: not implemented"
	return 0
}

func validateFuncOutputNone(err *error, funcType reflect.Type) { _ = "STUB: not implemented"; return }

func validateFuncOutputOneVarDynamic(err *error, funcType reflect.Type) int {
	_ = "STUB: not implemented"
	return 0
}

func validateFuncOutputOneVarBool(err *error, callbackType reflect.Type, isMust bool) int {
	_ = "STUB: not implemented"
	return 0
}

func forEachSlice(slice reflect.Value, sliceLen int, eachCallback func(reflect.Value, int)) {
	_ = "STUB: not implemented"
	return
}

func forEachSliceStoppable(slice reflect.Value, sliceLen int, eachCallback func(reflect.Value, int) bool) {
	_ = "STUB: not implemented"
	return
}

func forEachCollection(collection reflect.Value, keys []reflect.Value, eachCallback func(reflect.Value, reflect.Value, int)) {
	_ = "STUB: not implemented"
	return
}

func forEachCollectionStoppable(collection reflect.Value, keys []reflect.Value, eachCallback func(reflect.Value, reflect.Value, int) bool) {
	_ = "STUB: not implemented"
	return
}

func callFuncSliceLoop(funcToCall, param reflect.Value, i int, numIn int) []reflect.Value {
	_ = "STUB: not implemented"
	return nil
}

func callFuncCollectionLoop(funcToCall, value, key reflect.Value, numIn int) []reflect.Value {
	_ = "STUB: not implemented"
	return nil
}

func isSlice(err *error, label string, dataValue ...reflect.Value) bool {
	_ = "STUB: not implemented"
	return false
}

func isNonNilData(err *error, label string, data any) bool { _ = "STUB: not implemented"; return false }

func isZeroOrPositiveNumber(err *error, label string, size int) bool {
	_ = "STUB: not implemented"
	return false
}

func isPositiveNumber(err *error, label string, size int) bool {
	_ = "STUB: not implemented"
	return false
}

func isLeftShouldBeGreaterOrEqualThanRight(err *error, labelLeft string, valueLeft int, labelRight string, valueRight int) bool {
	_ = "STUB: not implemented"
	return false
}

func isTypeEqual(err *error, labelLeft string, typeLeft reflect.Type, labelRight string, typeRight reflect.Type) bool {
	_ = "STUB: not implemented"
	return false
}

func catch(err *error) { _ = "STUB: not implemented"; return }

func catchWithCustomErrorMessage(err *error, callback func(string) string) {
	_ = "STUB: not implemented"
	return
}
