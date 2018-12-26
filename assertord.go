package assertord

import (
	"fmt"
	"reflect"

	"github.com/stretchr/testify/assert"
)

func compare(obj1, obj2 interface{}, kind reflect.Kind) int {
	switch kind {
	case reflect.Int:
		{
			intobj1 := obj1.(int)
			intobj2 := obj2.(int)
			if intobj1 > intobj2 {
				return -1
			}
			if intobj1 == intobj2 {
				return 0
			}
			if intobj1 < intobj2 {
				return 1
			}
		}
	case reflect.Int8:
		{
			int8obj1 := obj1.(int8)
			int8obj2 := obj2.(int8)
			if int8obj1 > int8obj2 {
				return -1
			}
			if int8obj1 == int8obj2 {
				return 0
			}
			if int8obj1 < int8obj2 {
				return 1
			}
		}
	case reflect.Int16:
		{
			int16obj1 := obj1.(int16)
			int16obj2 := obj2.(int16)
			if int16obj1 > int16obj2 {
				return -1
			}
			if int16obj1 == int16obj2 {
				return 0
			}
			if int16obj1 < int16obj2 {
				return 1
			}
		}
	case reflect.Int32:
		{
			int32obj1 := obj1.(int32)
			int32obj2 := obj2.(int32)
			if int32obj1 > int32obj2 {
				return -1
			}
			if int32obj1 == int32obj2 {
				return 0
			}
			if int32obj1 < int32obj2 {
				return 1
			}
		}
	case reflect.Int64:
		{
			int64obj1 := obj1.(int64)
			int64obj2 := obj2.(int64)
			if int64obj1 > int64obj2 {
				return -1
			}
			if int64obj1 == int64obj2 {
				return 0
			}
			if int64obj1 < int64obj2 {
				return 1
			}
		}
	case reflect.Float32:
		{
			float32obj1 := obj1.(float32)
			float32obj2 := obj2.(float32)
			if float32obj1 > float32obj2 {
				return -1
			}
			if float32obj1 == float32obj2 {
				return 0
			}
			if float32obj1 < float32obj2 {
				return 1
			}
		}
	case reflect.Float64:
		{
			float64obj1 := obj1.(float64)
			float64obj2 := obj2.(float64)
			if float64obj1 > float64obj2 {
				return -1
			}
			if float64obj1 == float64obj2 {
				return 0
			}
			if float64obj1 < float64obj2 {
				return 1
			}
		}
	case reflect.String:
		{
			stringobj1 := obj1.(string)
			stringobj2 := obj2.(string)
			if stringobj1 > stringobj2 {
				return -1
			}
			if stringobj1 == stringobj2 {
				return 0
			}
			if stringobj1 < stringobj2 {
				return 1
			}
		}
	}

	return 0
}

// isOrdered checks that collection contains only orderable elements and in the one order.
func isOrdered(object interface{}, direction int) (ok bool, isOrdered bool) {
	objKind := reflect.TypeOf(object).Kind()
	if objKind != reflect.Slice && objKind != reflect.Array {
		return false, false
	}

	objValue := reflect.ValueOf(object)
	objLen := objValue.Len()

	if objLen == 0 {
		return false, false
	}

	if objLen == 1 {
		return true, true
	}

	firstValueKind := objValue.Index(0).Kind()
	for i := 1; i < objLen; i++ {
		value := objValue.Index(i)
		valueKind := value.Kind()

		if firstValueKind != valueKind {
			return false, false
		}

		prevValue := objValue.Index(i - 1)
		compareResult := compare(prevValue.Interface(), value.Interface(), firstValueKind)

		if direction != 0 && compareResult != 0 && direction != compareResult {
			return true, false
		}

		if direction == 0 && compareResult != 0 {
			direction = compareResult
		}
	}

	return true, true
}

type tHelper interface {
	Helper()
}

// IsOrdered asserts that the specified object is ordered
func IsOrdered(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}

	ok, ordered := isOrdered(object, 0)
	if !ok {
		return assert.Fail(t, fmt.Sprintf("\"%s\" could not be applied builtin ordering", object), msgAndArgs...)
	}

	if !ordered {
		assert.Fail(t, fmt.Sprintf("Should be ordered, but was %v", object), msgAndArgs...)
	}

	return ordered
}

// IsNotOrdered asserts that the specified object is not ordered
func IsNotOrdered(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}

	ok, ordered := isOrdered(object, 0)
	if !ok {
		return assert.Fail(t, fmt.Sprintf("\"%s\" could not be applied builtin ordering", object), msgAndArgs...)
	}

	if ordered {
		assert.Fail(t, fmt.Sprintf("Should be not ordered, but was %v", object), msgAndArgs...)
	}

	return !ordered
}

// IsOrderedAsc asserts that the specified object in ascending order
func IsOrderedAsc(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}

	ok, ordered := isOrdered(object, 1)
	if !ok {
		return assert.Fail(t, fmt.Sprintf("\"%s\" could not be applied builtin ordering", object), msgAndArgs...)
	}

	if !ordered {
		assert.Fail(t, fmt.Sprintf("Should be ordered, but was %v", object), msgAndArgs...)
	}

	return ordered
}

// IsNotOrderedAsc asserts that the specified object in not ascending order
func IsNotOrderedAsc(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}

	ok, ordered := isOrdered(object, 1)
	if !ok {
		return assert.Fail(t, fmt.Sprintf("\"%s\" could not be applied builtin ordering", object), msgAndArgs...)
	}

	if !ordered {
		assert.Fail(t, fmt.Sprintf("Should be ordered, but was %v", object), msgAndArgs...)
	}

	return !ordered
}

// IsOrderedDesc asserts that the specified object in desceding order
func IsOrderedDesc(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}

	ok, ordered := isOrdered(object, -1)
	if !ok {
		return assert.Fail(t, fmt.Sprintf("\"%s\" could not be applied builtin ordering", object), msgAndArgs...)
	}

	if !ordered {
		assert.Fail(t, fmt.Sprintf("Should be ordered, but was %v", object), msgAndArgs...)
	}

	return ordered
}

// IsNotOrderedDesc asserts that the specified object in not desceding order
func IsNotOrderedDesc(t assert.TestingT, object interface{}, msgAndArgs ...interface{}) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}

	ok, ordered := isOrdered(object, -1)
	if !ok {
		return assert.Fail(t, fmt.Sprintf("\"%s\" could not be applied builtin ordering", object), msgAndArgs...)
	}

	if !ordered {
		assert.Fail(t, fmt.Sprintf("Should be ordered, but was %v", object), msgAndArgs...)
	}

	return !ordered
}
