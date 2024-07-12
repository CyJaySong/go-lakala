package common

import (
	"crypto/sha256"
	"encoding/json"
)

func Sha256(data []byte) (res []byte) {
	h := sha256.New()
	h.Write(data)
	return h.Sum(nil)
}

func JsonRow(val any) (jsonRaw json.RawMessage, jsonStr string, err error) {
	var jsonBytes []byte
	jsonRaw = json.RawMessage{}
	if jsonBytes, err = json.Marshal(val); err != nil {
		return nil, "", err
	}
	if err = jsonRaw.UnmarshalJSON(jsonBytes); err != nil {
		return nil, "", err
	}
	jsonStr = string(jsonRaw)
	return
}
