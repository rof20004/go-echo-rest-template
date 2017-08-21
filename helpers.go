package main

import (
	"encoding/json"
	"fmt"
)

// ModelToJSON -
func ModelToJSON(model interface{}) string {
	j, err := json.Marshal(model)
	if err != nil {
		panic(fmt.Sprintf("Error %v encoding JSON for %v", err, model))
	}

	jsonStr := string(j)
	// v := reflect.Indirect(reflect.ValueOf(model))
	// ot := v.Type()
	// t := ot
	// isArray := false
	// if t.Kind() == reflect.Array || t.Kind() == reflect.Slice {
	// 	t = t.Elem()
	// 	isArray = true
	// } else if t.Kind() == reflect.Interface {
	// 	t = v.Elem().Type()
	// }
	return jsonStr
}

// ModelToJSONMap -
func ModelToJSONMap(modl interface{}) map[string]interface{} {
	jsonStr := ModelToJSON(modl)
	m := JSONToMap(jsonStr)

	return m
}

// JSONToMap -
func JSONToMap(jsonStr string) map[string]interface{} {
	jsonMap := make(map[string]interface{})

	err := json.Unmarshal([]byte(jsonStr), &jsonMap)
	if err != nil {
		panic(fmt.Sprintf("Error %v unmarshaling JSON for %v", err, jsonStr))
	}

	return jsonMap
}

// JSONToMapArray -
func JSONToMapArray(jsonStr string) []map[string]interface{} {
	var arr []map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &arr)

	if err != nil {
		panic(fmt.Sprintf("Error %v unmarshaling JSON for %v", err, jsonStr))
	}

	return arr
}

// JSONToModel -
func JSONToModel(jsonStr string, item interface{}) error {
	err := json.Unmarshal([]byte(jsonStr), &item)

	// if err == nil {
	// 	v := reflect.Indirect(reflect.ValueOf(item))
	// 	ot := v.Type()
	// 	t := ot
	// 	isArray := false
	// 	if t.Kind() == reflect.Array || t.Kind() == reflect.Slice {
	// 		t = t.Elem()
	// 		isArray = true
	// 	} else if t.Kind() == reflect.Interface {
	// 		t = v.Elem().Type()
	// 	}
	// }
	return err
}
