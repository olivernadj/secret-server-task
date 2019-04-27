/*
 * Secret Server
 *
 * This is an API of a secret service. You can save your secret by using the API. You can restrict the access of a secret after the certen number of views or after a certen period of time.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"time"
)

type Secret struct {

	// Unique hash to identify the secrets
	Hash string `json:"hash,omitempty"`

	// The secret itself
	SecretText string `json:"secretText,omitempty"`

	// The date and time of the creation
	CreatedAt time.Time `json:"createdAt,omitempty"`

	// The secret cannot be reached after this time
	ExpiresAt time.Time `json:"expiresAt,omitempty"`

	// How many times the secret can be viewed
	RemainingViews int32 `json:"remainingViews,omitempty"`

}



func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

// Encrypting Data with an AES Cipher
func (s * Secret) encrypt() error {
	if s.Hash != "" {
		return nil
	}
	j, _ := json.Marshal(s)
	s.Hash = GetMD5Hash(string(j))
	block, _ := aes.NewCipher([]byte(s.Hash))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}
	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return err
	}
	s.SecretText = string(gcm.Seal(nonce, nonce, []byte(s.SecretText), nil))
	s.SecretText = base64.StdEncoding.EncodeToString([]byte(s.SecretText))
	return nil
}

// Decrypting Data that uses an AES Cipher
func (s * Secret) getDecryptedSecret() (string, error) {
	if s.Hash == "" {
		return "", errors.New("secret has no hash")
	}
	data, err := base64.StdEncoding.DecodeString(s.SecretText)
	// data := []byte(s.SecretText)
	block, err := aes.NewCipher([]byte(s.Hash))
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonceSize := gcm.NonceSize()
	nonce, cipherText := data[:nonceSize], data[nonceSize:]
	plainText, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "", err
	}
	return string(plainText), nil
}

