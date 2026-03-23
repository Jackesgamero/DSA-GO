package strings

import "strings"

var dbs []map[string]string

const deleted = "__DELETED__"

func RunDBCommand(cmd string) string {
	parts := strings.Fields(cmd)
	if len(parts) == 0 {
		return ""
	}

	top := len(dbs) - 1

	switch parts[0] {
	case "SET":
		if len(parts) < 3 {
			return ""
		}
		key, value := parts[1], parts[2]
		dbs[top][key] = value
		return ""

	case "GET":
		key := parts[1]
		for i := len(dbs) - 1; i >= 0; i-- {
			if val, ok := dbs[i][key]; ok {
				if val == deleted {
					return "nil"
				}
				return val
			}
		}
		return "nil"

	case "EXISTS":
		key := parts[1]
		for i := len(dbs) - 1; i >= 0; i-- {
			if val, ok := dbs[i][key]; ok {
				if val == deleted {
					return "false"
				}
				return "true"
			}
		}
		return "false"

	case "UNSET":
		key := parts[1]
		dbs[top][key] = deleted
		return ""

	case "BEGIN":
		dbs = append(dbs, make(map[string]string))
		return ""

	case "ROLLBACK":
		if len(dbs) == 1 {
			return ""
		}
		dbs = dbs[:len(dbs)-1]
		return ""

	case "COMMIT":
		if len(dbs) == 1 {
			return ""
		}

		topMap := dbs[len(dbs)-1]
		for k, v := range topMap {
			if v == deleted {
				delete(dbs[len(dbs)-2], k)
			} else {
				dbs[len(dbs)-2][k] = v
			}
		}
		dbs = dbs[:len(dbs)-1]
		return ""

	default:
		return ""
	}
}
