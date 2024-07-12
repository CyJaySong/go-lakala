package open

import (
	"crypto/rsa"
	"github.com/cyjaysong/lakala-go/common"
)

type Client struct {
	appid         string          // 机构或渠道接入申请的appid
	serialNo      string          // 接入方加签用的证书序列号
	appPrivateKey *rsa.PrivateKey // 机构私钥证书
	lklPublicKey  *rsa.PublicKey  // 拉卡拉公钥证书
	encryptKey    string          // 敏感信息加密秘钥
	isProd        bool            // 是否正式环境
}

func NewClient(appid, appPrivatePem string, lklPublicPem string, isProd bool) (client *Client, err error) {
	priKey, err := common.ParseRsaPrivatePem([]byte(appPrivatePem))
	if err != nil {
		return nil, err
	}
	pubKey, err := common.ParseRsaPublicPem([]byte(lklPublicPem))
	if err != nil {
		return nil, err
	}
	client = &Client{appid: appid, appPrivateKey: priKey, lklPublicKey: pubKey, isProd: isProd}
	return client, nil
}
