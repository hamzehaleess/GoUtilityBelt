package crypto

import (
        "encoding/hex"
        "strings"
        "testing"
)

func TestMD5Hash(t *testing.T) {
        tests := []struct {
                input    string
                expected string
        }{
                {"hello", "5d41402abc4b2a76b9719d911017c592"},
                {"", "d41d8cd98f00b204e9800998ecf8427e"},
                {"test", "098f6bcd4621d373cade4e832627b4f6"},
        }

        for _, test := range tests {
                result := MD5Hash(test.input)
                if result != test.expected {
                        t.Errorf("MD5Hash(%q) = %q; expected %q", test.input, result, test.expected)
                }
        }
}

func TestSHA256Hash(t *testing.T) {
        tests := []struct {
                input    string
                expected string
        }{
                {"hello", "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"},
                {"", "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"},
                {"test", "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08"},
        }

        for _, test := range tests {
                result := SHA256Hash(test.input)
                if result != test.expected {
                        t.Errorf("SHA256Hash(%q) = %q; expected %q", test.input, result, test.expected)
                }
        }
}

func TestGenerateUUID(t *testing.T) {
        uuid1 := GenerateUUID()
        uuid2 := GenerateUUID()

        if uuid1 == uuid2 {
                t.Error("GenerateUUID() should generate unique UUIDs")
        }

        if len(uuid1) != 36 {
                t.Errorf("GenerateUUID() length = %d; expected 36", len(uuid1))
        }

        // Check format (8-4-4-4-12)
        parts := strings.Split(uuid1, "-")
        if len(parts) != 5 {
                t.Errorf("GenerateUUID() format invalid: %s", uuid1)
        }
}

func TestRandomString(t *testing.T) {
        lengths := []int{5, 10, 20}

        for _, length := range lengths {
                result := RandomString(length)
                if len(result) != length {
                        t.Errorf("RandomString(%d) length = %d; expected %d", length, len(result), length)
                }

                // Check that two calls generate different strings
                result2 := RandomString(length)
                if result == result2 {
                        t.Error("RandomString() should generate different strings")
                }
        }
}

func TestBase64EncodeAndDecode(t *testing.T) {
        tests := []string{
                "hello world",
                "test string",
                "",
                "special chars: !@#$%^&*()",
        }

        for _, test := range tests {
                encoded := Base64Encode(test)
                decoded, err := Base64Decode(encoded)
                
                if err != nil {
                        t.Errorf("Base64Decode failed for %q: %v", test, err)
                }

                if decoded != test {
                        t.Errorf("Base64 encode/decode failed for %q: got %q", test, decoded)
                }
        }
}

func TestRandomBytes(t *testing.T) {
        lengths := []int{8, 16, 32}

        for _, length := range lengths {
                bytes, err := RandomBytes(length)
                if err != nil {
                        t.Errorf("RandomBytes(%d) failed: %v", length, err)
                }

                if len(bytes) != length {
                        t.Errorf("RandomBytes(%d) length = %d; expected %d", length, len(bytes), length)
                }

                // Check that two calls generate different bytes
                bytes2, _ := RandomBytes(length)
                if string(bytes) == string(bytes2) {
                        t.Error("RandomBytes() should generate different byte arrays")
                }
        }
}

func TestGenerateToken(t *testing.T) {
        lengths := []int{8, 16, 32}

        for _, length := range lengths {
                token, err := GenerateToken(length)
                if err != nil {
                        t.Errorf("GenerateToken(%d) failed: %v", length, err)
                }

                // Token should be hex encoded, so length should be double
                expectedLength := length * 2
                if len(token) != expectedLength {
                        t.Errorf("GenerateToken(%d) length = %d; expected %d", length, len(token), expectedLength)
                }

                // Verify it's valid hex
                _, err = hex.DecodeString(token)
                if err != nil {
                        t.Errorf("GenerateToken(%d) generated invalid hex: %s", length, token)
                }
        }
}

func TestHashAndVerifyPassword(t *testing.T) {
        passwords := []string{
                "password123",
                "secure_password",
                "test",
        }

        for _, password := range passwords {
                hash := HashPassword(password)
                
                if !VerifyPassword(password, hash) {
                        t.Errorf("VerifyPassword failed for password %q", password)
                }

                if VerifyPassword("wrong_password", hash) {
                        t.Errorf("VerifyPassword incorrectly verified wrong password for %q", password)
                }
        }
}
