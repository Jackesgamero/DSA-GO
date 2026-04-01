package main

import (
	"fmt"
	"strings"
)

func UpdatePreference(inputString string, userIndex int, prefKey, newValue string) map[string]interface{} {
	result := make(map[string]interface{})

	// 1. Split the string into individual users
	users := strings.Split(inputString, "User")
	for _, userData := range users {
		if userData == "" {
			continue
		}

		// Reconstruct the key (e.g., "User1") and split from attributes
		parts := strings.SplitN(userData, ":", 2)
		if len(parts) < 2 {
			continue
		}
		userKey := "User" + parts[0]
		attrString := parts[1]

		// Parse the attributes for this user
		result[userKey] = parseAttributes(attrString)
	}

	// 2. Modify the specific value
	targetUser := fmt.Sprintf("User%d", userIndex)
	if userMap, ok := result[targetUser].(map[string]interface{}); ok {
		// We look for nested maps (Preferences) to find the key
		for _, val := range userMap {
			if nestedMap, ok := val.(map[string]interface{}); ok {
				if _, exists := nestedMap[prefKey]; exists {
					nestedMap[prefKey] = newValue
				}
			}
		}
	}

	return result
}

// parseAttributes handles key=value pairs and nested {key=value} structures
func parseAttributes(input string) map[string]interface{} {
	m := make(map[string]interface{})

	// We need to split by ';' but ignore ';' inside curly braces
	start := 0
	depth := 0
	for i := 0; i <= len(input); i++ {
		if i == len(input) || (input[i] == ';' && depth == 0) {
			pair := input[start:i]
			if strings.Contains(pair, "=") {
				kv := strings.SplitN(pair, "=", 2)
				key := kv[0]
				val := kv[1]

				if strings.HasPrefix(val, "{") && strings.HasSuffix(val, "}") {
					// Handle nested structure
					inner := val[1 : len(val)-1]
					m[key] = parseAttributes(inner)
				} else {
					m[key] = val
				}
			}
			start = i + 1
		} else if i < len(input) && input[i] == '{' {
			depth++
		} else if i < len(input) && input[i] == '}' {
			depth--
		}
	}
	return m
}

func PrettyPrint(dict map[string]interface{}, indent string) string {
	result := "{\n"
	for key, value := range dict {
		switch value := value.(type) {
		case map[string]interface{}:
			result += fmt.Sprintf("%s  %s: %s", indent, key, PrettyPrint(value, indent+"  "))
		default:
			result += fmt.Sprintf("%s  %s: '%v',\n", indent, key, value)
		}
	}
	result += indent + "}\n"
	return result
}

func main() {
	input := "User1:Age1=21;Location1=USA;Preferences1={Food1=Italian;Sport1=Fencing};User2:Age2=30;Location2=Canada;Preferences2={Music2=Jazz;Color2=Blue}"
	userIndex := 1
	prefKey := "Sport1"
	newValue := "Hockey"
	updatedPreferences := UpdatePreference(input, userIndex, prefKey, newValue)
	fmt.Println(PrettyPrint(updatedPreferences, ""))
}
