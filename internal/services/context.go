// Package services ...
package services

import (
	"context"
	"reflect"
)

// ServiceKey ...
const ServiceKey = "services"

type serviceMap map[reflect.Type]reflect.Value

// NewServiceContext ...
func NewServiceContext(c context.Context) context.Context {
	if c.Value(ServiceKey) == nil {
		return context.WithValue(c, ServiceKey, make(serviceMap))
	}

	return c
}
