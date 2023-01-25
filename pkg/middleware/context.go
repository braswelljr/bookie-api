package middleware

import (
	"context"
	"fmt"
)

var Store = CtxValues{map[string]interface{}{
	"token": "",
	"uid":   "",
	"role":  "",
}}                                                                   // this is the store that will be used to store the context values
var ContextKey CtxKey = "bookie-api-7kt2"                            // this is the key that will be used to store the context values
var CXT = context.WithValue(context.Background(), ContextKey, Store) // this is the context that will be used to store the context values

// String - returns the string representation of the context values.
func (v *CtxValues) String() string {
	return fmt.Sprintf("%v", v.m)
}

// GetValue - returns the value of the context key provided.
//
// @param key - string
// @return interface{}
func (v *CtxValues) GetCtxValue(key string) interface{} {
	return v.m[key]
}

// SetValue - sets the value of the context key provided.
//
// @param key - string
// @param value - interface{}
func (v *CtxValues) SetCtxValue(key string, value interface{}) interface{} {
	v.m[key] = value

	return v.GetCtxValue(key)
}
