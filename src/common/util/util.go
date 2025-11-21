package util

import (
	"fmt"
	"strings"
)

func RemoveSpace(data string) string {

	// remove whitespace characters (spaces, newlines, tabs)
	data = strings.ReplaceAll(data, " ", "")
	data = strings.ReplaceAll(data, "\n", "")
	data = strings.ReplaceAll(data, "\t", "")

	return data
}

func Convert(i interface{}) interface{} {
	switch x := i.(type) {
	case map[interface{}]interface{}:
		m := make(map[string]interface{})
		for k, v := range x {
			m[fmt.Sprint(k)] = Convert(v)
		}
		return m
	case []interface{}:
		for i, v := range x {
			x[i] = Convert(v)
		}
	}
	return i
}
