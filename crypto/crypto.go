// Package crypto provides utility functions for cryptographic operations
package crypto

import (
        "crypto/md5"
        "crypto/rand"
        "crypto/sha256"
        "encoding/base64"
        "encoding/hex"
        "math/big"

        "github.com/google/uuid"
)

// MD5Hash generates an MD5 hash of the input string
func MD5Hash(input string) string {
        hash := md5.Sum([]byte(input))
        return hex.EncodeToString(hash[:])
}

// SHA256Hash generates a SHA256 hash of the input string
func SHA256Hash(input string) string {
        hash := sha256.Sum256([]byte(input))
        return hex.EncodeToString(hash[:])
}

// GenerateUUID generates a new UUID
func GenerateUUID() string {
        return uuid.New().String()
}

// RandomString generates a random string of specified length
func RandomString(length int) string {
        const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
        result := make([]byte, length)
        
        for i := range result {
                num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
                result[i] = charset[num.Int64()]
        }
        
        return string(result)
}

// Base64Encode encodes a string to base64
func Base64Encode(input string) string {
        return base64.StdEncoding.EncodeToString([]byte(input))
}

// Base64Decode decodes a base64 string
func Base64Decode(input string) (string, error) {
        decoded, err := base64.StdEncoding.DecodeString(input)
        if err != nil {
                return "", err
        }
        return string(decoded), nil
}

// RandomBytes generates random bytes of specified length
func RandomBytes(length int) ([]byte, error) {
        bytes := make([]byte, length)
        _, err := rand.Read(bytes)
        if err != nil {
                return nil, err
        }
        return bytes, nil
}

// GenerateToken generates a secure random token
func GenerateToken(length int) (string, error) {
        bytes, err := RandomBytes(length)
        if err != nil {
                return "", err
        }
        return hex.EncodeToString(bytes), nil
}

// HashPassword creates a simple hash of a password (Note: Use bcrypt in production)
func HashPassword(password string) string {
        return SHA256Hash(password + "salt") // Simple example, use proper salt in production
}

// VerifyPassword verifies a password against its hash
func VerifyPassword(password, hash string) bool {
        return HashPassword(password) == hash
}
