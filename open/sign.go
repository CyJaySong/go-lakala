package open

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// Sign [timeStamp] 秒级时间戳, [nonceStr] 12个字符的随机字符串 [body] 请求体报文
func (t *Client) Sign(timeStamp int, nonceStr, body string) (signBase64 string, err error) {
	str := fmt.Sprintf("%s\n%s\n%d\n%s\n%s\n", t.appid, t.serialNo, timeStamp, nonceStr, body)
	h := sha256.New()
	h.Write([]byte(str))
	hashed := h.Sum(nil)
	signature, err := rsa.SignPKCS1v15(rand.Reader, t.appPrivateKey, crypto.SHA256, hashed)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString([]byte(signature)), nil
}
