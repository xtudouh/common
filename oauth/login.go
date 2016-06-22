package oauth

import (
    "bytes"
    "fmt"
    "net/http"
    "net/url"
    "xtudouh/common/conf"
    "regexp"
    "strings"
)

const (
    auth_tpl = "%s?response_type=code&client_id=%s&redirect_uri=%s"
)

var (
    authPath string
    redirectURI string
)

func Login(username, password string) (*http.Response, error) {
    authUrl := fmt.Sprintf(auth_tpl, authPath, url.QueryEscape(clientId), url.QueryEscape(redirectURI))
    l.Debug("auth url: %s", authUrl)
    resp, err := http.Get(authUrl)
    if err != nil {
        l.Error("%v\n", err)
        return nil, err
    }

    in := new(bytes.Buffer)
    in.ReadFrom(resp.Body)
    defer resp.Body.Close()
    var token string
    reg, err := regexp.Compile(`<input type="hidden" name="access_token".*>`)
    if err != nil {
        l.Error("%v\n", err)
        return nil, err
    }
    line := reg.FindString(in.String())
    token = strings.Split(line, "\"")[5]

    client := &http.Client{}
    req, err := http.NewRequest("POST", authUrl, strings.NewReader(fmt.Sprintf("token=%s&username=%s&password=%s", token, username, username)))
    if err != nil {
        l.Error("%v\n", err)
        return nil, err
    }
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    return client.Do(req)
}

func init() {
    registerInitFun(func() {
        authPath = conf.String("sso", "AUTH_URL")
        redirectURI = conf.String("sso", "REDIRECT_URI")
    })
}
