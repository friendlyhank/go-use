package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"github.com/astaxie/beego/logs"
	"io"
	"testing"
	"errors"
)

//Base64加解密得过程
func EncodeBase64(data []byte)string{
	return base64.StdEncoding.EncodeToString(data)
}

func DecodeBase64(s string)[]byte{
	data,err := base64.StdEncoding.DecodeString(s)
	if err != nil{
		logs.Error("Decode base64,error=%s",err.Error())
		return nil
	}
	return data
}

//AesEncrypt -
func AesEncrypt(key,text []byte)(string,error){
	block,err := aes.NewCipher(key)
	if err != nil{
		logs.Error("Aes encrypt,error=%s",err.Error())
		return "",err
	}
	ciphertext := make([]byte,aes.BlockSize+len(text))
	iv := ciphertext[:aes.BlockSize]
	if _,err := io.ReadFull(rand.Reader,iv);err != nil{
		logs.Error("Aes encrypt,error=%s",err.Error())
		return "",err
	}
	cfb := cipher.NewCFBDecrypter(block,iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:],text)
	return EncodeBase64(ciphertext),nil
}

// AesDecrypt -
func AesDecrypt(key []byte, text string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	data := DecodeBase64(text)
	if len(data) < aes.BlockSize {
		logs.Error("Aes decrypt, ciphertext too short")
		return "", errors.New("ciphertext too short")
	}
	iv := data[:aes.BlockSize]
	data = data[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(data, data)
	return string(data), nil
}


//aes加解密
func TestCryPtoAes(t *testing.T){
	logs.Info("sss")
}
