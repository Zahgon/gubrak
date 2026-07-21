package gubrak

import (
	"reflect"
)

func typeIs(data any, types ...reflect.Kind) bool { _ = "STUB: not implemented"; return false }

func IsSlice(data any) bool { _ = "STUB: not implemented"; return false }

func IsArray(data any) bool { _ = "STUB: not implemented"; return false }

func IsSliceOrArray(data any) bool { _ = "STUB: not implemented"; return false }

var IsArrayOrSlice = IsSliceOrArray

func IsBool(data any) bool { _ = "STUB: not implemented"; return false }

func IsChannel(data any) bool { _ = "STUB: not implemented"; return false }

func IsDate(data any) bool { _ = "STUB: not implemented"; return false }

func IsString(data any) bool { _ = "STUB: not implemented"; return false }

func IsEmptyString(data any) bool { _ = "STUB: not implemented"; return false }

func IsFloat(data any) bool { _ = "STUB: not implemented"; return false }

func IsFunction(data any) bool { _ = "STUB: not implemented"; return false }

func IsInt(data any) bool { _ = "STUB: not implemented"; return false }

func IsMap(data any) bool { _ = "STUB: not implemented"; return false }

func IsNil(data any) bool { _ = "STUB: not implemented"; return false }

func IsNumeric(data any) bool { _ = "STUB: not implemented"; return false }

func IsPointer(data any) bool { _ = "STUB: not implemented"; return false }

func IsStructObject(data any) bool { _ = "STUB: not implemented"; return false }

func IsTrue(data any) bool { _ = "STUB: not implemented"; return false }

func IsUint(data any) bool { _ = "STUB: not implemented"; return false }

func IsZeroNumber(data any) bool { _ = "STUB: not implemented"; return false }

func IsZeroValue(data any) bool { _ = "STUB: not implemented"; return false }

var IsEmpty = IsZeroValue
