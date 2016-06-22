package wechat

import (
    "xtudouh/common/conf"
    "xtudouh/common/log"
    "xtudouh/common/utils"
    "fmt"
    "time"
)

const (
    WECHAT_TOKEN = "prosnav:oauth:wechat:token:access:"
    WECHAT_REFRESH = "prosnav:oauth:wechat:token:refresh:"
)

var (
    Appid, appSecret, AccessToken, JsTicket string
    refreshTokenExpiration int
    logger = log.NewLogger()
)

type (
    WechatToken struct {
        AccessToken  string `form:"access_token" json:"access_token"`
        ExpiresIn    int    `from:"expires_in" json:"expires_in"`
        Openid       string `from:"openid" json:"openid"`
        RefreshToken string `from:"refresh_token" json:"refresh_token"`
        ErrCode      int    `form:"errcode" json:"errcode"`
        ErrMsg       string `form:"errmsg" json:"errmsg"`
    }

    Token struct {
        AccessToken string  `json:"access_token"`
        ExpiresIn   int     `json:"expires_in"`
    }

    JsApiTicket struct {
        ErrCode int `json:"errcode"`
        ErrMsg  string `json:"errmsg"`
        Ticket  string `json:"ticket"`
        ExpiresIn int  `json:"expires_in"`
    }

    UserInfo struct {
        OpenId     string   `form:"openid" json:"openid"`
        NickName   string   `form:"nickname" json:"nickname"`
        Sex        int      `form:"sex" json:"sex"`
        Province   string   `form:"province" json:"province"`
        City       string   `form:"city" json:"city"`
        Country    string   `form:"country" json:"country"`
        HeadImgurl string   `form:"headimgurl" json:"headimgurl"`
        Privilege  []string `form:"privilege" json:"privilege"`
        UnionId    string   `form:"unionid" json:"unionid"`
        ErrCode    int      `form:"errcode" json:"errcode"`
        ErrMsg     string   `form:"errmsg" json:"errmsg"`
    }
)

func refreshToken() {
    var token Token
    accessTokenUrl := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s", Appid, appSecret)
    if err := utils.FetchUrlWithJson(accessTokenUrl, &token); err != nil {
        logger.Error("%v", err)
        return
    }
    AccessToken = token.AccessToken
    logger.Debug("Wechat access token: %s", AccessToken)
}

func refreshJsTicket() {
    jsApiUrl := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/get_jsapi_ticket?access_token=%s", AccessToken)
    var ticket JsApiTicket
    if err := utils.FetchUrlWithJson(jsApiUrl, &ticket); err != nil {
        logger.Error("%v", err)
        return
    }
    if ticket.ErrCode != 0 {
        logger.Error("Request js api ticket failed, errmsg: %s", ticket.ErrMsg)
        return
    }
    JsTicket = ticket.Ticket
    logger.Debug("Wechat JS api ticket: %s", JsTicket)
}

func refresh() {
    refreshToken()
    refreshJsTicket()
}

func run() {
    refresh()
    tiker := time.NewTicker(7100 * time.Second)
    for {
        select {
        case <- tiker.C: refresh()
        }
    }
}

func Init() {
    fmt.Println("Wechat module start initializing.")
    Appid = conf.String("oauth2.wechat", "APP_ID")
    appSecret = conf.String("oauth2.wechat", "APP_SECRET")
    go run()
    fmt.Println("Wechat module initialize successfully.")
}
