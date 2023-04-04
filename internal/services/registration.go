// Package services ...
package services

import (
	"context"
	"reflect"
	"sync"
)

// AddTrainsient ...
func AddTrainsient(factoryFunc interface{}) (err error) {
	return addService(Transient, factoryFunc)
}

// AddScoped ...
func AddScoped(factoryFunc interface{}) (err error) {
	return addService(Scoped, factoryFunc)
}

// AddSingletone ...
func AddSingletone(factoryFunc interface{}) (err error) {
	factoryFuncVal := reflect.ValueOf(factoryFunc)
	if factoryFuncVal.Kind() == reflect.Func && factoryFuncVal.Type().NumOut() == 1 {
		var result []reflect.Value
		once := sync.Once{}
		wrapper := reflect.MakeFunc(factoryFuncVal.Type(), func([]reflect.Value) []reflect.Value {
			once.Do(func() { result = invokeFunction(context.TODO(), factoryFuncVal) })
			return result
		})
		err = addService(Singleton, wrapper.Interface())
	}
	return err
}
