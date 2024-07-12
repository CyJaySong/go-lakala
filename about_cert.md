# 关于证书


## Der、Pem、Pfx、Cer\Crt、X509 它们都是扩展名（文件名的后缀，代表格式）

这些扩展名都与数字证书和公钥/私钥对有关，但它们用于不同的场景和格式。下面是每个扩展名的简要说明：
1. **DER（Distinguished Encoding Rules）**:
    - DER是一种编码标准，用于将ASN.1（抽象语法标记）结构的数据编码成字节序列。
    - .der文件是二进制格式的证书或密钥文件，不包含私钥信息。
    - 它通常用于Java应用和其他需要二进制编码证书的场景。
2. **PEM（Privacy Enhanced Mail）**:
    - PEM是一种基于Base64编码的文本格式，可以包含证书、私钥、公钥和证书请求。
    - .pem文件以“-----BEGIN...-----”和“-----END...-----”标记开始和结束。
    - 它在OpenSSL和其他许多网络和加密工具中广泛使用。
3. **PFX/P12（Personal Information Exchange）**:
    - PFX是微软的一种证书文件格式，也称为PKCS#12。
    - .pfx或.p12文件包含私钥、公钥和证书，可以同时保护它们。
    - 它通常用于导入或导出证书和私钥，特别是在微软的产品中。
4. **CER/CRT（Certificate）**:
    - CER和CRT通常指的是证书文件，可以是二进制DER格式或Base64编码的PEM格式。
    - 在不同的系统和工具中，这两个扩展名可能互换使用。
    - 它们用于存储公钥证书，不包含私钥信息。
5. **X509**:
    - X.509是一个标准，定义了公钥证书的格式。
    - .x509文件通常指的是证书文件，但这个扩展名并不常见。
    - 它更多地用于描述证书的遵循的标准，而不是文件格式。
      总结：这些扩展名代表不同的文件格式和编码方式，用于存储证书、私钥和公钥。选择哪种格式取决于具体的应用场景和需求。


编码（也用作扩展）

    DER = DER扩展用于二进制DER编码证书。这些文件也可能带有CER或CRT扩展名。正确的英语用法是“我有DER编码证书”而不是“我有DER证书”。
    PEM = PEM扩展名用于不同类型的X.509v3文件，这些文件包含前缀为“-BEGIN ...”行的ASCII（Base64）装甲数据。
组合

    在某些情况下，将多个X.509基础结构组合成单个文件是有利的。一个常见的例子是将私钥和公钥组合到同一个证书中。

组合证书密钥和链的最简单方法是将每个密钥转换为PEM编码证书，然后将每个文件的内容简单复制到新文件中。这适用于组合文件以在Apache应用程序中使用。

萃取

    有些证书将以合并形式出现。其中一个文件可以包含以下任何一个：证书，私钥，公钥，签名证书，证书颁发机构（CA）和/或授权链。

PEM 格式

    PEM格式通常用于数字证书认证机构（Certificate Authorities，CA），扩展名为.pem, .crt, .cer, and .key。
    内容为Base64编码的ASCII码文件，有类似"-----BEGIN CERTIFICATE-----" 和 "-----END CERTIFICATE-----"的头尾标记。
    服务器认证证书，中级认证证书和私钥都可以储存为PEM格式（认证证书其实就是公钥）。Apache和类似的服务器使用PEM格式证书。

DER 格式

    DER格式与PEM不同之处在于其使用二进制而不是Base64编码的ASCII。扩展名为.der，
    但也经常使用.cer用作扩展名，所有类型的认证证书和私钥都可以存储为DER格式。Java是其典型使用平台。

转换方式

    可以使用OpenSSL命令行工具在不同证书格式之间的转换

PEM to DER

    openssl x509 -outform der -in certificate.pem -out certificate.der

PEM to PFX

    openssl pkcs12 -export -out certificate.pfx -inkey privateKey.key -in certificate.crt -certfile CACert.crt

DER to PEM

    openssl x509 -inform der -in certificate.cer -out certificate.pem

PFX to PEM

    openssl pkcs12 -in certificate.pfx -out certificate.cer -nodes
    (PFX转PEM后certificate.cer文件包含认证证书和私钥，需要把它们分开存储才能使用。)  

知识点：

    1、使用公钥操作数据属于加密
    
    2、使用私钥对原文的摘要操作属于签名
    
    3、公钥和私钥可以互相加解密
    
    4、不同格式的证书之间可以互相转换
    
    5、公钥可以对外公开，但是私钥千万不要泄露，要妥善保存

注意：在我们备份证书信息的时候，最好使用.jks或者.pfx文件进行保存，这样备份的证书文件可以被完整的导出。