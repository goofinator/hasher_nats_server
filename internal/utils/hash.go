package utils

import "crypto/sha256"

// MakeHashes make hashes from message
func MakeHashes(message []byte) [][]byte {
	result := make([][]byte, 0)
	for i := byte(0); i < 100; i++ {
		h := sha256.New()
		h.Write(message)

		// to make differen variants of hash
		h.Write([]byte{i})

		result = append(result, h.Sum(nil))
	}

	return result
}
