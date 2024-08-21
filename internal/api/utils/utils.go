package utils

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"io"
	"strings"

	"github.com/jackc/pgx/v4"
	"golang.org/x/crypto/bcrypt"

	"github.com/wesleyfebarretos/ticket-sale/internal/api/config"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/exception"
	"github.com/wesleyfebarretos/ticket-sale/internal/infra/db"
)

func HashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashPassword), nil
}

func ComparePassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func WithTransaction[R any](c context.Context, fn func(pgx.Tx) R) R {
	tx, err := db.Conn.Begin(c)
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}
	defer tx.Rollback(c)

	response := fn(tx)

	if err := tx.Commit(c); err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	return response
}

func MaskCreditcardNumber(number string) string {
	firstFourDigits := number[:4]
	lastFourDigits := number[len(number)-4:]
	masked := strings.Repeat("*", len(number)-8)

	return firstFourDigits + masked + lastFourDigits
}

func Encrypt(plaintext string) string {
	// Hash the API token to create a 32-byte key
	hash := sha256.Sum256([]byte(config.Envs.ApiToken))
	key := hash[:]

	// Create a new AES cipher with the key
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	// Create a byte slice to hold the IV + ciphertext
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))

	// Generate a random IV and store it at the beginning of the ciphertext slice
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	// Create the CFB encrypter and XOR the plaintext with it to create the ciphertext
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(plaintext))

	// Encode the ciphertext to base64 to make it easily transferable
	return base64.StdEncoding.EncodeToString(ciphertext)
}

func Decrypt(ciphertextBase64 string) string {
	// Decode the base64-encoded ciphertext
	ciphertext, err := base64.StdEncoding.DecodeString(ciphertextBase64)
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	// Hash the API token to create a 32-byte key
	hash := sha256.Sum256([]byte(config.Envs.ApiToken))
	key := hash[:]

	// Create a new AES cipher with the key
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	// Ensure the ciphertext is long enough to contain the IV
	if len(ciphertext) < aes.BlockSize {
		panic(exception.InternalServerException("ciphertext too short"))
	}

	// Extract the IV from the beginning of the ciphertext
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	// Create the CFB decrypter and XOR the ciphertext with it to recover the plaintext
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	// Return the decrypted plaintext as a string
	return string(ciphertext)
}
