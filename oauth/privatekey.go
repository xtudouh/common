package oauth

import (
    "xtudouh/common/conf"
    "os"
    "bytes"
)

var pk []byte

func LookupPrivateKey() []byte {
    return pk
}

var signedMethod string

func SignedMethod() string {
    if signedMethod == "" {
        signedMethod = conf.String("jwt", "SIGNED_METHOD")
    }
    return signedMethod
}

func init() {
    registerInitFun(func() {
        pkFile := conf.String("jwt", "PRIVATE_KEY")
        in, err := os.Open(pkFile)
        if err != nil {
            panic(err)
        }
        defer in.Close()
        buf := new(bytes.Buffer)
        buf.ReadFrom(in)
        pk = buf.Bytes()
    })
}