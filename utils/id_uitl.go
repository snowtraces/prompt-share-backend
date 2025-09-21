package utils

import (
	"crypto/rand"
	"fmt"
	mathrand "math/rand" // 重命名导入以避免冲突
	"time"
)

// UUID represents a universally unique identifier
type UUID [16]byte

// NewUUID generates a new UUID v4
func NewUUID() (UUID, error) {
	var uuid UUID

	// Generate random bytes using crypto/rand
	_, err := rand.Read(uuid[:])
	if err != nil {
		return uuid, err
	}

	// Set version (4) and variant bits
	uuid[6] = (uuid[6] & 0x0f) | 0x40 // Version 4
	uuid[8] = (uuid[8] & 0x3f) | 0x80 // Variant is 10

	return uuid, nil
}

// String returns the string representation of the UUID
func (u UUID) String() string {
	return fmt.Sprintf("%x-%x-%x-%x-%x",
		u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
}

// IdWorkerWithUUID generates an ID using UUID v4
func IdWorkerWithUUID() string {
	uuid, err := NewUUID()
	if err != nil {
		// Fallback to original method if UUID generation fails
		return idWorker(16)
	}
	return uuid.String()
}

var RandAlphaNumber = []byte("0123456789abcdefghijklmnopqrstuvwxyz")

func idWorker(len int) string {
	var id []byte

	var r = mathrand.New(mathrand.NewSource(time.Now().UnixNano()))
	for i := 0; i < len; i++ {
		id = append(id, RandAlphaNumber[(r.Intn(35)+1)])
	}

	return string(id)
}
