package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func JsonAssertion(src, dest interface{}) (err error) {
	var bts []byte
	if bts, err = json.Marshal(src); err != nil {
		return
	}

	if err = json.Unmarshal(bts, dest); err != nil {
		return
	}

	return
}

func PrintJson(input any) {
	fmt.Println(JsonIndent(input))
}

func JsonIndent(v any) string {
	b, err := json.MarshalIndent(v, "", "	")
	if err != nil {
		return ""
	}

	return string(b)
}

func Indent(v []byte) string {
	var out bytes.Buffer

	err := json.Indent(&out, v, "", "	")
	if err == nil {
		return string(out.Bytes())
	}

	return ""
}
