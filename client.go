package main

import (
    "net/http"
    "fmt"
    "io/ioutil"
    "crypto/tls"
    "crypto/x509"
    "bytes"
)

func main() {
    pool := x509.NewCertPool()
    caCertPath := "ca.crt"

    caCrt, err := ioutil.ReadFile(caCertPath)
    if err != nil {
        fmt.Println("ReadFile err", err)
        return
    }   
    pool.AppendCertsFromPEM(caCrt)

    cliCrt, err := tls.LoadX509KeyPair("client.crt", "client.key")
    if err != nil {
        fmt.Println("LoadX509KeyPair err:", err)
        return
    }   
    tr := &http.Transport{
        TLSClientConfig: &tls.Config{
            RootCAs:      pool,
            Certificates: []tls.Certificate{cliCrt},
            InsecureSkipVerify: true,//客户端关闭对服务端的验证
        },
    }
    client := &http.Client{Transport: tr}
    jsonStr := "{\"name\":\"wang\",\"age\":25}"
    req := bytes.NewBuffer([]byte(jsonStr))
    body_type := "application/json;charset=utf-8"
    resp, err := client.Post("https://test.com:8088/post",body_type,req)
    if err != nil {
        fmt.Println("Get error:", err)
        return
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))
}
