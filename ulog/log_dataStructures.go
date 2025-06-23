package ulog

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

// PrintMap prints the contents of a map to the standard output.
// It recursively handles nested maps, printing their contents with appropriate indentation.
//
// Parameters:
//   - m: The map to be printed
//
// Example:
//
//	data := map[string]any{"key1": "value1", "key2": map[string]any{"nested": "value"}}
//	PrintMap(data)
func PrintMap(m map[string]any) {
	for key, value := range m {

		switch v := value.(type) {
		case map[string]any:
			fmt.Printf("%s: {\n", key)
			PrintMap(v) // Recursive call for nested maps
			fmt.Println("}")
		default:
			fmt.Printf("%s: %v\n", key, value)
		}
	}
}

// PrintList prints a list of strings to the standard output with 1-based indexing.
// Each item is printed on a new line with its index number.
//
// Parameters:
//   - list: The string slice to be printed
//
// Example:
//
//	items := []string{"apple", "banana", "cherry"}
//	PrintList(items)
//	// Output:
//	// 1: apple
//	// 2: banana
//	// 3: cherry
func PrintList(list []string) {
	for i, item := range list {
		fmt.Printf("%d: %s\n", i+1, item)
	}
}

// MapAsPrettyString converts a map[string]interface{} to a pretty string representation.
// This is useful for logging or sending error messages in a readable format.
//
// Parameters:
//   - m: The map to be converted to a string
//   - beforeMessage: Optional parameter that can be used to prepend a message before the map.
//     Only the first message in the variadic parameter is used.
//
// Returns:
//   - A formatted string representation of the map
//
// Example:
//
//	data := map[string]interface{}{"name": "John", "age": 30}
//	str := MapAsPrettyString(data, "User info:")
//	// str = "User info: {age: 30, name: "John"}"
func MapAsPrettyString(m map[string]interface{}, beforeMessage ...string) string {
	result := "{"
	first := true
	for k, v := range m {
		if !first {
			result += ", "
		}
		first = false
		result += k + ": " + ValueAsString(v)
	}
	result += "}"

	if len(beforeMessage) > 0 {
		return beforeMessage[0] + " " + result
	}

	return result
}

// ValueAsString converts a value of any type to its string representation.
// It handles different types appropriately, including nested maps.
//
// Parameters:
//   - v: The value to be converted to a string
//
// Returns:
//   - A string representation of the value
//
// Example:
//
//	str := ValueAsString("hello")  // Returns: "hello"
//	str := ValueAsString(42)       // Returns: 42
//	str := ValueAsString(map[string]interface{}{"key": "value"})  // Returns: {key: "value"}
func ValueAsString(v interface{}) string {
	switch val := v.(type) {
	case string:
		return `"` + val + `"`
	case int, int64, float64, float32, bool:
		return fmt.Sprintf("%v", val)
	case map[string]interface{}:
		return MapAsPrettyString(val)
	default:
		return fmt.Sprintf("%v", val)
	}
}

// ListAsPrettyString converts a slice of strings to a pretty string representation.
// This is useful for logging or displaying string lists in a readable format.
//
// Parameters:
//   - list: The string slice to be converted to a string
//   - beforeMessage: Optional parameter that can be used to prepend a message before the list.
//     Only the first message in the variadic parameter is used.
//
// Returns:
//   - A formatted string representation of the list
//
// Example:
//
//	items := []string{"apple", "banana", "cherry"}
//	str := ListAsPrettyString(items, "Fruits:")
//	// str = "Fruits: ["apple", "banana", "cherry"]"
func ListAsPrettyString(list []string, beforeMessage ...string) string {
	result := "["
	first := true
	for _, v := range list {
		if !first {
			result += ", "
		}
		first = false
		result += `"` + v + `"`
	}
	result += "]"
	if len(beforeMessage) > 0 {
		return beforeMessage[0] + " " + result
	}

	return result
}

// ListAsPrettyStringWithIndex converts a slice of strings to a pretty string representation
// with 1-based indexing for each element.
//
// Parameters:
//   - list: The string slice to be converted to a string
//   - beforeMessage: Optional parameter that can be used to prepend a message before the list.
//     Only the first message in the variadic parameter is used.
//
// Returns:
//   - A formatted string representation of the list with indexes
//
// Example:
//
//	items := []string{"apple", "banana", "cherry"}
//	str := ListAsPrettyStringWithIndex(items, "Fruits:")
//	// str = "Fruits: [1: "apple", 2: "banana", 3: "cherry"]"
func ListAsPrettyStringWithIndex(list []string, beforeMessage ...string) string {
	result := "["
	first := true
	for i, v := range list {
		if !first {
			result += ", "
		}
		first = false
		result += fmt.Sprintf("%d: \"%s\"", i+1, v)
	}
	result += "]"
	if len(beforeMessage) > 0 {
		return beforeMessage[0] + " " + result
	}

	return result
}

// PrintMapWithIndent prints a map with proper indentation for nested structures.
// It sorts keys alphabetically for consistent output and handles nested maps and arrays.
//
// Parameters:
//   - m: The map to be printed
//   - indent: The current indentation string to use
//
// Example:
//
//	data := map[string]interface{}{
//	    "user": map[string]interface{}{
//	        "name": "John",
//	        "scores": []interface{}{95, 87, 92},
//	    },
//	}
//	PrintMapWithIndent(data, "")
func PrintMapWithIndent(m map[string]interface{}, indent string) {
	// Get all keys and sort them for consistent output
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		v := m[k]
		switch value := v.(type) {
		case map[string]interface{}:
			fmt.Printf("%s%s: {\n", indent, k)
			PrintMapWithIndent(value, indent+"  ")
			fmt.Printf("%s}\n", indent)
		case []interface{}:
			fmt.Printf("%s%s: [\n", indent, k)
			for i, item := range value {
				if nestedMap, ok := item.(map[string]interface{}); ok {
					fmt.Printf("%s  [%d]: {\n", indent, i)
					PrintMapWithIndent(nestedMap, indent+"    ")
					fmt.Printf("%s  }\n", indent)
				} else {
					fmt.Printf("%s  [%d]: %v\n", indent, i, item)
				}
			}
			fmt.Printf("%s]\n", indent)
		default:
			fmt.Printf("%s%s: %v\n", indent, k, v)
		}
	}
}

// Print struct prints the contents of a struct to the standard output.
// It internally converts the struct to a map and then prints it using PrintMap.
func PrintStruct(data interface{}) {
	mapData, err := ConvertStructToMap(data)
	if err != nil {
		fmt.Println("Error converting struct to map:", err)
		return
	}
	PrintMap(mapData)
}

// ConvertToMap converts any struct to a map[string]interface{}
// ConvertStructToMap converts a Go struct to a map[string]interface{} using JSON marshaling and unmarshaling.
// This function takes any Go struct as input and serializes it to JSON, then deserializes the JSON back into a map.
// It is useful for dynamic manipulation, logging, or inspection of struct fields without knowing their types at compile time.
//
// Parameters:
//   - data: The input struct or value to be converted. It can be any value that is serializable by the encoding/json package.
//
// Returns:
//   - map[string]interface{}: A map representation of the struct, where keys are the struct's JSON field names and values are the corresponding field values.
//   - error: An error if the marshaling or unmarshaling process fails.
//
// Note:
//   - Fields in the struct must be exported (start with an uppercase letter) to be included in the resulting map.
//   - JSON struct tags are respected during the conversion.
//   - Nested structs and slices are recursively converted to maps and slices, respectively.
func ConvertStructToMap(data interface{}) (map[string]interface{}, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// FormatJSON returns a formatted JSON string with specified indentation
// FormatJSON takes any Go data structure as input and returns its JSON representation as a formatted string.
// The 'indent' parameter specifies the number of spaces to use for each indentation level in the output JSON.
// If the marshaling process encounters an error, an empty string and the error are returned.
// Parameters:
//   - data: The Go data structure to be marshaled into JSON.
//   - indent: The number of spaces to use for each indentation level in the output JSON.
//
// Returns:
//   - string: The formatted JSON string representation of the input data.
//   - error: An error object if marshaling fails, otherwise nil.
//
// Example usage:
//
//	jsonStr, err := FormatJSON(myStruct, 4)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(jsonStr)
func FormatJSON(data interface{}, indent int) (string, error) {
	indentStr := strings.Repeat(" ", indent)
	jsonData, err := json.MarshalIndent(data, "", indentStr)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}
