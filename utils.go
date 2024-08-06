package urchain

import "encoding/json"

func FormatStruct(v any) string {
	objstr, _ := json.MarshalIndent(v, "", "\t")
	return string(objstr)
}
