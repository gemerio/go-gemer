package crypto

import (
	"crypto/sha256"
	"json"
)

func generateHash(moduleblocks []ModuleBlock) [64]byte {
	arrayBytes := []byte{}
	for _, codex := range moduleblocks {
		jsonReadBytes, _ := json.Marshal(codex)
		arrayBytes = append(arrayBytes, jsonReadBytes...)
	}
	return sha256.Sum(arrayBytes)
}
