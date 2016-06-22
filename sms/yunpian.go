package sms

import (
    "errors"
    "fmt"
    "github.com/cJrong/YunPianSMS"
    "xtudouh/common/conf"
    "xtudouh/common/log"
    "xtudouh/common/redis"
    "time"
)

const (
    SHORT_MSG_NUM_TAB = "sms:"
    TPL1 = "您的验证码是#code#【#company#】"
    TPL2 = "您的验证码是#code#。如非本人操作，请忽略本短信【#company#】"
    TPL3 = "亲爱的#name#，您的验证码是#code#。如非本人操作，请忽略本短信【#company#】"
    TPL4 = "亲爱的#name#，您的验证码是#code#。有效期为#hour#小时，请尽快验证【#company#】"
    TPL5 = "感谢您注册#app#，您的验证码是#code#【#company#】"
    TPL6 = "欢迎使用#app#，您的手机验证码是#code#。本条信息无需回复【#company#】"
    TPL7 = "正在找回密码，您的验证码是#code#【#company#】"
    TPL8 = "激活码是#code#。如非本人操作，请致电#tel#【#company#】"
    TPL9 = "#code#(Prosnav Investment 手机动态码，请完成验证)，如非本人操作，请忽略本短信【#company#】"
)

var (
    ApiKey string
    enable bool
    company string
    maxNum int64
    interval int64
    l = log.NewLogger()
)

func SendVerifyCode(code string, mobile string) (interface{}, error) {
    if !enable {
        return "", errors.New("Sms module is not enabled.")
    }
    mobileTab := smsNumTab(mobile)
    currNum, _ := redis.RedisClient.Get(mobileTab).Int64()
    if currNum >= maxNum {
        return "", errors.New("Too many short message have sent.")
    }
    patten := fmt.Sprintf("#code#=%s,#company#=%s", code, company)
    result, err := YunPianSMS.YunPianSMSTPLSend(ApiKey, TPL9, patten, mobile)
    if err != nil {
        l.Error("error:%v\n", err)
        return nil, err
    }
    if currNum == 0 {
        redis.RedisClient.Incr(mobileTab)
        redis.RedisClient.Expire(mobileTab, time.Duration(interval) * time.Second)
        return result, nil
    }
    redis.RedisClient.Incr(mobileTab)
    return result, nil
}

func smsNumTab(mobile string) string {
    return fmt.Sprintf("%s%s", SHORT_MSG_NUM_TAB, mobile)
}

func Init() {
    l.Debug("sms module start initializing.")
    ApiKey = conf.String("sms", "APIKEY")
    enable = conf.Bool("sms", "ENABLE", false)
    company = conf.String("sms", "COMPANY")
    maxNum = conf.Int64("sms", "MAX_NUM", 20)
    interval = conf.Int64("sms", "INTERVAL", 3600)
    if enable && ApiKey == "" {
        panic(errors.New("SMS module is enabled, but api is not configured."))
    }
    l.Debug("sms module initialize successfully.")
}
