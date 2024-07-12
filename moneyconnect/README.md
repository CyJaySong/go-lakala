# 拉卡拉钱账通

### 证书编码转换后，解析失败问题

`openssl x509 -inform der -in certificate.cer -out certificate.pem`

当使用以上命令将的der编码公钥证书转换成pem编码公钥证书后，Go可能会解析失败

可以使用java语言脚本获取证书 base64，然后替换`certificate.pem` 的 base64 部分

```java
package com.lakala.moneyconnect;

import java.util.Base64;
import java.io.File;
import java.io.InputStream;
import java.io.FileInputStream;
import java.security.cert.CertificateFactory;
import java.security.cert.X509Certificate;
import java.security.PublicKey;

public class PublicKeyDer2PemConvert {
    public static void main(String[] args) throws Exception {
        File file = new File("Der编码公钥证书路径");
        InputStream inputStream = new FileInputStream(file);
        CertificateFactory certificateFactory = CertificateFactory.getInstance("X.509");
        X509Certificate cert = (X509Certificate) certificateFactory.generateCertificate(inputStream);
        PublicKey publicKey = cert.getPublicKey();
        Base64.Encoder encoder = Base64.getEncoder();
        System.out.println("public_base64:" + encoder.encodeToString(publicKey.getEncoded()));
    }
}

```