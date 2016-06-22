package oauth

import (
    "fmt"
    "xtudouh/common/conf"
    "xtudouh/common/utils"
)

var (
    tokenUrl, clientId, clientSecret string
    refreshTokenExpiration int
)

type Token struct {
    AccessToken  string `json:"access_token"`
    ExpiresIn    int    `json:"expires_in"`
    RefreshToken string `json:"refresh_token"`
    TokenType    string `json:"token_type"`
}

func RequestAccessToken(code string, accessToken *Token) error {
    tokenUrl := fmt.Sprintf("%s?grant_type=authorization_code&code=%s&client_id=%s&client_secret=%s", tokenUrl, code, clientId, clientSecret)
    if err := utils.FetchUrlWithJson(tokenUrl, accessToken); err != nil {
        l.Error("%v", err)
        return err
    }
    return nil
}

func RequestRefresh(token *Token) error {
    tokenUrl := fmt.Sprintf("%s?grant_type=refresh_token&refresh_token=%s&client_id=%s&client_secret=%s", tokenUrl, token.RefreshToken, clientId, clientSecret)
    if err := utils.FetchUrlWithJson(tokenUrl, token); err != nil {
        l.Error("%v", err)
        return err
    }
    return nil
}

func init() {
    registerInitFun(func() {
        tokenUrl = conf.String("sso", "TOKEN_URL")
        clientId = conf.String("sso", "CLIENT_ID")
        clientSecret = conf.String("sso", "CLIENT_SECRET")
        refreshTokenExpiration = conf.Int("sso", "REFRESH_TOKEN_EXPIRATION")
    })
}
