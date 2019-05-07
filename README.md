# go-https
go-https


# 生成密钥、证书

### 第一步，为服务器端和客户端准备公钥、私钥

生成服务器端私钥
```
openssl genrsa -out server.key 1024
```
生成服务器端公钥
```
openssl rsa -in server.key -pubout -out server.pem
```
生成客户端私钥
```
openssl genrsa -out client.key 1024
```
生成客户端公钥
```
openssl rsa -in client.key -pubout -out client.pem
```
### 第二步，生成 CA 证书

生成 CA 私钥
```
openssl genrsa -out ca.key 1024
```
X.509 Certificate Signing Request (CSR) Management.
```
openssl req -new -key ca.key -out ca.csr
```
X.509 Certificate Data Management.
```
openssl x509 -req -in ca.csr -signkey ca.key -out ca.crt
```
在执行第二步时
Common Name (e.g. server FQDN or YOUR name) []: 这一项，是最后可以访问的域名，为了方便测试，写成 test.com

### 第三步，生成服务器端证书和客户端证书

服务器端需要向 CA 机构申请签名证书，在申请签名证书之前依然是创建自己的 CSR 文件
```
openssl req -new -key server.key -out server.csr
```
向自己的 CA 机构申请证书，签名过程需要 CA 的证书和私钥参与，最终颁发一个带有 CA 签名的证书
```
openssl x509 -req -CA ca.crt -CAkey ca.key -CAcreateserial -in server.csr -out server.crt
```
client 端
```
openssl req -new -key client.key -out client.csr
```
client 端到 CA 签名
```
openssl x509 -req -CA ca.crt -CAkey ca.key -CAcreateserial -in client.csr -out client.crt
```
### 第四步，客户端机器添加test.com域名
```
sudo vim /etc/hosts
```
添加服务端的ip信息对应的域名：192.168.1.100  test.com
