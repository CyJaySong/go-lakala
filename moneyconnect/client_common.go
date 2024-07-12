package moneyconnect

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cyjaysong/lakala-go/common"
	reqclient "github.com/imroc/req/v3"
	"strconv"
	"strings"
	"time"
)

func (t *Client) getApiUrl() string {
	if t.isProd {
		return prodApiUrl
	}
	return testApiUrl
}

func (t *Client) getBaseBody(service string) map[string]any {
	return map[string]any{"version": "3.0", "app_id": t.appid, "service": service,
		"timestamp": strconv.FormatInt(time.Now().Unix(), 10),
	}
}

func (t *Client) buildRequest(val map[string]any, others ...string) (req *reqclient.Request) {
	val = t.sign(val, others...)
	return t.reqClient.R().SetBodyJsonMarshal(val)
}

func (t *Client) buildUploadRequest(val map[string]any, others ...string) (req *reqclient.Request) {
	val = t.sign(val, others...)
	return t.reqClient.R().SetFormDataAnyType(val)
}

func (t *Client) reqPost(req *reqclient.Request, baseRes any) (err error) {
	response, err := req.Post(t.getApiUrl())
	if err != nil {
		return err
	}
	var responseBody struct {
		Sign     string `json:"sign"`
		Response string `json:"response"`
	}
	if err = response.UnmarshalJson(&responseBody); err != nil {
		return err
	}
	var bytes = []byte(responseBody.Response)
	if verifyPass := t.Verify(responseBody.Sign, bytes); !verifyPass {
		err = errors.New("response signature verification failed")
		return
	}
	return json.Unmarshal(bytes, baseRes)
}

func (t *Client) reqPostBytes(req *reqclient.Request) (bytes []byte, err error) {
	response, err := req.Post(t.getApiUrl())
	if err != nil {
		return nil, err
	}
	var responseBody struct {
		Sign     string `json:"sign"`
		Response string `json:"response"`
	}
	if err = response.UnmarshalJson(&responseBody); err != nil {
		return nil, err
	}
	bytes = []byte(responseBody.Response)
	if verifyPass := t.Verify(responseBody.Sign, bytes); !verifyPass {
		err = errors.New("response signature verification failed")
	}
	return
}

func (t *Client) reqPostDL(req *reqclient.Request) (res *reqclient.Response, err error) {
	res, err = req.Post(t.getApiUrl())
	if err != nil {
		return nil, err
	}
	return
}

// Sign [val] 基本参数, [others] 其他参与签名的参数
func (t *Client) sign(val map[string]any, others ...string) map[string]any {
	var timestamp, version, service = val["timestamp"], val["version"], val["service"]
	str := fmt.Sprintf("%s%s%s%s%s", t.appid, timestamp, version, service, strings.Join(others, ""))
	hashed := common.Sha256([]byte(str))
	signPKCS1v15, err := rsa.SignPKCS1v15(rand.Reader, t.merchantSignPrivateKey, crypto.SHA256, hashed)
	if err != nil {
		return nil
	}
	// 将签名转换为Base64字符串
	val["sign"] = base64.StdEncoding.EncodeToString(signPKCS1v15)
	return val
}

// Verify [sign] base64签名字符串, [body] 待验证签名内容
func (t *Client) Verify(sign string, body []byte) (verifyPass bool) {
	rsaSignBytes, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return false
	}
	hashed := common.Sha256(body)
	err = rsa.VerifyPKCS1v15(t.cloudSignPublicKey, crypto.SHA256, hashed, rsaSignBytes)
	return err == nil
}

// CloudRsaEncrypt 钱账通Rsa公钥加密
func (t *Client) CloudRsaEncrypt(body []byte) (encryptRes []byte, err error) {
	return rsa.EncryptPKCS1v15(rand.Reader, t.cloudEncryptPublicKey, body)
}

// MerchantRsaDecrypt 商家平台Rsa私钥解密
func (t *Client) MerchantRsaDecrypt(body []byte) (decryptRes []byte, err error) {
	return rsa.DecryptPKCS1v15(rand.Reader, t.merchantEncryptPrivateKey, body)
}
