package moneyconnect

import (
	"crypto/rsa"
	reqclient "github.com/imroc/req/v3"
	"time"
)

const testApiUrl = "http://47.105.146.241:9510/service/soa"
const prodApiUrl = "https://moneyconnect.lakala.com/service/soa"

type Client struct {
	appid                     string          // 机构或渠道接入申请的appid
	serialNo                  string          // 接入方加签用的证书序列号
	merchantSignPrivateKey    *rsa.PrivateKey // 商家平台签名Rsa私钥
	merchantEncryptPrivateKey *rsa.PrivateKey // 商家平台加密Rsa私钥
	cloudSignPublicKey        *rsa.PublicKey  // 钱账通签名Rsa公钥
	cloudEncryptPublicKey     *rsa.PublicKey  // 钱账通加密Rsa公钥
	isProd                    bool            // 是否正式环境
	reqClient                 *reqclient.Client
}

func NewClient(appid string, merchantSignPrivateKey, merchantEncryptPrivateKey *rsa.PrivateKey,
	cloudSignPublicKey, cloudEncryptPublicKey *rsa.PublicKey, isProd bool) (client *Client) {
	client = &Client{appid: appid, isProd: isProd,
		merchantSignPrivateKey: merchantSignPrivateKey, merchantEncryptPrivateKey: merchantEncryptPrivateKey,
		cloudSignPublicKey: cloudSignPublicKey, cloudEncryptPublicKey: cloudEncryptPublicKey}
	client.reqClient = reqclient.C().SetTimeout(time.Second*10).SetCommonRetryCount(1).
		SetCommonHeader("Authorization", "QZTAPI-SHA256withRSA").SetUserAgent("")
	//if !isProd {
	client.reqClient.DevMode()
	//}
	return
}
