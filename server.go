package main

import (
    "net/http"
    "log"
    "crypto/x509"
    "io/ioutil"
    "fmt"
    "crypto/tls"
)

type myhandler struct {
}

func (h *myhandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Content-Type", "text/plain")
    w.Write([]byte("This is an example server\n"))
}

func getHandler(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "hello, this is https get ")
}

func postHandler(w http.ResponseWriter, r *http.Request)  {
    body, _ := ioutil.ReadAll(r.Body)
    r.Body.Close()
    body_str := string(body)
    fmt.Println(body_str)
    //ret, _ := json.Marshal(user)
    ret := "{\"code\":200}"
    fmt.Fprint(w, string(ret))
}

func main() {
    pool := x509.NewCertPool()
    caCertPath := "ca.crt"
    
    caCrt, err := ioutil.ReadFile(caCertPath)
    if err != nil {
        fmt.Println("ReadFile err", err)
        return
    }   
    pool.AppendCertsFromPEM(caCrt)
    
    server := http.NewServeMux()
    server.HandleFunc("/get",getHandler)
    server.HandleFunc("/post",postHandler)
    
    s := &http.Server{
        Addr:    ":8088",
        //Handler: &myhandler{},
        Handler: server,
        TLSConfig: &tls.Config{
            ClientCAs:  pool,
            ClientAuth: tls.RequireAndVerifyClientCert,
        },  
    }
    
    log.Printf("About to listen on 10443.Go to https://127.0.0.1:8088")
    err = s.ListenAndServeTLS("server.crt", "server.key")
    if err != nil {
        log.Fatal(err)
    }   
}
