package helpers

import (
	"encoding/json"
)

func Jason(in any) []byte {
	b, _ := json.Marshal(in)
	return b
}
