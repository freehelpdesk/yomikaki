package deskjson

import (
	"fmt"
	"strings"
)

func DirectWrite(path string, jsonObj map[string]interface{}, newValue interface{}) map[string]interface{} {
	keys := strings.Split(path, "->")

	// Traverse the map and retrieve the sub-map at the given path
	subMap := jsonObj
	for _, key := range keys[:len(keys)-1] {
		if _, ok := subMap[key]; !ok {
			subMap[key] = make(map[string]interface{})
		}
		var ok bool
		subMap, ok = subMap[key].(map[string]interface{})
		if !ok {
			return jsonObj
		}
	}

	// Create a copy of the original sub-map and update the value at the given path
	newSubMap := make(map[string]interface{})
	for k, v := range subMap {
		newSubMap[k] = v
	}
	newSubMap[keys[len(keys)-1]] = newValue

	// Create a copy of the original map and replace the sub-map at the given path with the modified sub-map
	newMap := make(map[string]interface{})
	for k, v := range jsonObj {
		newMap[k] = v
	}
	newMap[keys[0]] = replaceSubMap(newMap[keys[0]].(map[string]interface{}), keys[1:], newSubMap)

	return newMap
}

func replaceSubMap(m map[string]interface{}, keys []string, newSubMap map[string]interface{}) map[string]interface{} {
	if len(keys) == 1 {
		return newSubMap
	}
	newM := make(map[string]interface{})
	for k, v := range m {
		newM[k] = v
	}
	newM[keys[0]] = replaceSubMap(m[keys[0]].(map[string]interface{}), keys[1:], newSubMap)
	return newM
}

func DirectRead(jsonPath string, jsonObj map[string]interface{}) (interface{}, error) {
	// Split the path into individual keys
	keys := strings.Split(jsonPath, "->")

	// Traverse the JSON object using the keys in the path
	for i, key := range keys {
		// Check if the key exists in the JSON object
		val, ok := jsonObj[key]
		if !ok {
			return nil, fmt.Errorf("key '%s' not found in JSON object", key)
		}

		// If this is the last key in the path, return the value
		if i == len(keys)-1 {
			return val, nil
		}

		// Otherwise, continue traversing the JSON object
		var ok2 bool
		jsonObj, ok2 = val.(map[string]interface{})
		if !ok2 {
			return nil, fmt.Errorf("key '%s' is not an object in JSON", key)
		}
	}

	return nil, fmt.Errorf("invalid JSON path '%s'", jsonPath)
}
