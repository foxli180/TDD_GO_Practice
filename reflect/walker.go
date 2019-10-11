package reflect

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i:= 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		//mapkeys 返回数组 所以 _ 是 index， key 是数组的value
		for _, key := range val.MapKeys(){
			walkValue(val.MapIndex(key))
		}
	case reflect.String:
		fn(val.String())

	}
}

//
//func walk(x interface{}, fn func(input string)) {
//	val := getValue(x)
//
//	numberOfValues := 0
//	var getField func(int) reflect.Value
//
//	switch val.Kind() {
//	case reflect.Struct:
//		numberOfValues = val.NumField()
//		getField = val.Field
//	case reflect.Slice, reflect.Array:
//		numberOfValues = val.Len()
//		getField = val.Index
//	case reflect.Map:
//		//mapkeys 返回数组 所以 _ 是 index， key 是数组的value
//		for _, key := range val.MapKeys(){
//			walk(val.MapIndex(key).Interface(), fn)
//		}
//	case reflect.String:
//		fn(val.String())
//	}
//
//	for i:= 0; i < numberOfValues; i++ {
//		walk(getField(i).Interface(), fn)
//	}
//}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return val
}
