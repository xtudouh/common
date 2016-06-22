package oauth

import (
    "fmt"
    "xtudouh/common/conf"
    "xtudouh/common/utils"
)

var userUrl string

func RequestUserInfo(token string, user interface{}) error {
    oneTimeUserUrl := fmt.Sprintf("%s?code=%s", userUrl, token)
    if err := utils.FetchUrlWithJson(oneTimeUserUrl, &user); err != nil {
        l.Error("%v", err)
        return err
    }
    return nil
}

func init() {
    registerInitFun(func() {
        userUrl = conf.String("sso", "USER_URL")
    })
}
