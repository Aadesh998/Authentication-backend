package config

import (
	"crypto/rand"
	"encoding/hex"
	"log"
)

var JWTSecret []byte

func init() {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		log.Fatal("Failed to generate the key:", err)
	}
	JWTSecret = []byte(hex.EncodeToString(bytes))
}
