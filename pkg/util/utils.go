package util

import "fmt"

func GetStringVal(key string, opt map[string]any) (string, error) {
	if _, ok := opt[key]; !ok {
		return "", fmt.Errorf("lacks %s information in map", key)
	}
	val, ok := opt[key].(string)
	if !ok {
		return "", fmt.Errorf("variable %v cannot be asserted as type string", opt[key])
	}
	return val, nil
}

func GetIntVal(key string, opt map[string]any) (int, error) {
	if _, ok := opt[key]; !ok {
		return 0, fmt.Errorf("lacks %s information in map", key)
	}
	val, ok := opt[key].(int)
	if !ok {
		return 0, fmt.Errorf("variable %v cannot be asserted as type int", opt[key])
	}
	return val, nil
}
