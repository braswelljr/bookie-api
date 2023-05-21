package middleware

import (
	"fmt"
)

// this is the store that will be used to store the context values
var ContextKey CtxKey = "bookie-api-7kt2" // this is the key that will be used to store the context values

// String - returns the string representation of the context values.
func (v *CtxValues) String() string {
	return fmt.Sprintf("%v", v.m)
}

// GetValue - returns the value of the context key provided.
//
// @param key - string
// @return interface{}
func (v *CtxValues) GetCtxValue(key string) interface{} {
	// check if the key exists in the map
	if _, ok := v.m[key]; !ok {
		return nil
	}

	// return the value
	return v.m[key]
}

// SetValue - sets the value of the context key provided.
//
// @param key - string
// @param value - interface{}
func (v *CtxValues) SetCtxValue(key string, value interface{}) error {
	// check if the key exists in the map
	if _, ok := v.m[key]; ok {
		v.m[key] = value

		// return nil if the key exists
		return nil
	}

	// return an error if the key does not exist
	return fmt.Errorf("unable to set context value: key %s does not exist", key)
}
