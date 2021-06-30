package common

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"math/big"
	"math/rand"
	"time"
	cRand "crypto/rand"
)

func MD5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)

	return fmt.Sprintf("%x", has)
}

// EncryptPassword: Return encrypt password and random string
func EncryptPassword(password string) (string, string) {
	rand.Seed(time.Now().UnixNano())
	randStr := CreateRandomString(rand.Intn(5))
	return MD5(randStr + password), randStr
}

// CheckEncryptPassword: auth encrypt password
func CheckEncryptPassword(password, randomString , encryptPassword string) bool{
	if MD5(randomString + password) == encryptPassword{
		return true
	}else{
		return false
	}
}
// CreateRandomString: create random string
func CreateRandomString(len int) string {
	var container string
	var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < len; i++ {
		randomInt, _ := cRand.Int(cRand.Reader, bigInt)
		container += string(str[randomInt.Int64()])
	}
	return container
}
