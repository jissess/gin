package util

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

func RandomString(length int) string {
	var letters = []byte("asdfghjklqwertyuiopzxcvbnmASDFGHJKLQWERTYUIOPZXCVBNM")
	result := make([]byte, length)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func EncoderSha256(str string) string {
	h := sha256.New()
	h.Write([]byte(str))
	sum := h.Sum(nil)
	res := hex.EncodeToString(sum)
	return string(res)
}
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func Base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func Base64Decode(str string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(str)
}

func Mkdir(basePath string) string {
	folderName := time.Now().Format("2006/01/02")
	folderPath := filepath.Join(basePath, folderName)
	os.MkdirAll(folderPath, os.ModePerm)
	return folderPath
}
