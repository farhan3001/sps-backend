package utils

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"hash"
	"time"
)

var (
	mac        hash.Hash
	lastSecret string
)

// ByteArrayToHexString converts a byte array to a hex string
func ByteArrayToHexString(bytes []byte) string {
	return hex.EncodeToString(bytes)
}

// GetSha calculates SHA hash of the input string with the specified algorithm
func GetSha(plainText, algorithm string) string {
	var hasher hash.Hash

	switch algorithm {
	case "SHA-1":
		hasher = sha1.New()
	case "SHA-256":
		hasher = sha256.New()
	case "SHA-512":
		hasher = sha512.New()
	default:
		// Default to SHA-256 if unknown algorithm
		hasher = sha256.New()
	}

	hasher.Write([]byte(plainText))
	hashBytes := hasher.Sum(nil)
	return ByteArrayToHexString(hashBytes)
}

// getObjMAC returns a HMAC hasher for the given secret key
func getObjMAC(secretKey string) (hash.Hash, error) {
	if mac == nil || lastSecret != secretKey {
		lastSecret = secretKey
		keyBytes := []byte(secretKey)
		mac = hmac.New(sha512.New, keyBytes)
	}

	// In Go, we can't clone the hash.Hash directly like in Java
	// Instead, we'll create a new one with the same key
	newMac := hmac.New(sha512.New, []byte(lastSecret))
	return newMac, nil
}

// HmacSha512 calculates HMAC-SHA512 of the input string with the given secret
func HmacSha512(clientSecret, stringToSign string) (string, error) {
	lmac, err := getObjMAC(clientSecret)
	if err != nil {
		return "", err
	}

	lmac.Write([]byte(stringToSign))
	hmacBytes := lmac.Sum(nil)

	encoded := base64.StdEncoding.EncodeToString(hmacBytes)
	return encoded, nil
}

// MinifyJSONString converts JSON to a compact one-line format
func MinifyJSONString(input string) (string, error) {
	var buf bytes.Buffer
	if err := json.Compact(&buf, []byte(input)); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func GenerateSignatureForGetToken(clientSecret, path, httpMethod, requestBody string, now time.Time) (string, string, error) {

	// Format date and time according to the pattern
	dateStr := now.Format("2006-01-02")
	timeStr := now.Format("15:04:05")
	ts := fmt.Sprintf("%sT%s+07:00", dateStr, timeStr)

	requestStringMinified, err := MinifyJSONString(requestBody)
	if err != nil {
		return "", "", nil
	}

	println()
	fmt.Println("1: ", requestStringMinified)
	println()

	sha256Hash := GetSha(requestStringMinified, "SHA-256")
	clearSig := fmt.Sprintf("%s:%s:%s:%s", httpMethod, path, sha256Hash, ts)
	signature, err := HmacSha512(clientSecret, clearSig)
	if err != nil {
		return "", "", nil
	}

	return signature, ts, nil

}

func GenerateSignatureForAPIReq(clientSecret, clientKey string, now time.Time) (string, string, error) {

	// Format date and time according to the pattern
	dateStr := now.Format("2006-01-02")
	timeStr := now.Format("15:04:05")
	ts := fmt.Sprintf("%sT%s+07:00", dateStr, timeStr)

	// sha256Hash := GetSha(requestStringMinified, "SHA-256")
	stringToSign := fmt.Sprintf("%s|%s", clientKey, ts)
	signature, err := HmacSha512(clientSecret, stringToSign)
	if err != nil {
		return "", "", nil
	}

	return signature, ts, nil

}

func GenerateSignature2(now time.Time) {
	// Get current time

	// Format date and time according to the pattern
	dateStr := now.Format("2006-01-02")
	timeStr := now.Format("15:04:05")
	ts := fmt.Sprintf("%sT%s+07:00", dateStr, timeStr)

	stringReq := `{"grantType":"client_credentials","additionalInfo":"{}"}`
	println()
	fmt.Println("2: ", stringReq)
	println()

	clinentKey := "02002"
	path := "/riskService/auth/v1/access_token"
	secretKey := "930561e9-ee9f-41ed-82b0-aa81249c1563"

	// Calculate SHA-256 hash of konten
	sha256Hash := GetSha(stringReq, "SHA-256")
	clearSig := fmt.Sprintf("POST:%s:%s:%s", path, sha256Hash, ts)

	// Calculate HMAC-SHA512 signature
	signature, err := HmacSha512(secretKey, clearSig)
	if err != nil {
		fmt.Println("ERROR MAKING SIGNATURE")
	}

	fmt.Println("Client Key:", clinentKey)
	fmt.Println("Timestamp:", ts)
	fmt.Println("Signature:", signature)

}
