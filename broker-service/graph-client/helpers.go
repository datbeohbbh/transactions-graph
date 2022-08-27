package graph

import (
	"encoding/json"
)

func readJson(data any) []byte {
	b, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return b
}

func writeJson(data []byte, v any) {
	err := json.Unmarshal(data, &v)
	if err != nil {
		panic(err)
	}
}
