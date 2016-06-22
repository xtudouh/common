package conf

import (
    "gopkg.in/ini.v1"
    "time"
)

var (
    Cfg                                 *ini.File
    ENV, AppName, AppVer, ListeningPort string
)

func String(sec, key string, def ...string) string {
    if Cfg == nil {
        return ""
    }
    if len(def) > 0 {
        return Cfg.Section(sec).Key(key).MustString(def[0])
    }

    return Cfg.Section(sec).Key(key).String()
}

func Strings(sec, key, delim string) []string {
    if Cfg == nil {
        return nil
    }
    return Cfg.Section(sec).Key(key).Strings(delim)
}

func Int(sec, key string, def ...int) int {
    if Cfg == nil {
        return 0
    }
    return Cfg.Section(sec).Key(key).MustInt(def...)
}

func Ints(sec, key, delim string) []int {
    if Cfg == nil {
        return nil
    }
    return Cfg.Section(sec).Key(key).Ints(delim)
}

func Int64(sec, key string, def ...int64) int64 {
    if Cfg == nil {
        return 0
    }
    return Cfg.Section(sec).Key(key).MustInt64(def...)
}

func Int64s(sec, key, delim string) []int64 {
    if Cfg == nil {
        return nil
    }
    return Cfg.Section(sec).Key(key).Int64s(delim)
}

func Bool(sec, key string, def ...bool) bool {
    if Cfg == nil {
        return false
    }
    return Cfg.Section(sec).Key(key).MustBool(def...)
}

func Float64(sec, key string, def ...float64) float64 {
    if Cfg == nil {
        return 0.0
    }
    return Cfg.Section(sec).Key(key).MustFloat64(def...)
}

func Float64s(sec, key, delim string) []float64 {
    if Cfg == nil {
        return nil
    }
    return Cfg.Section(sec).Key(key).Float64s(delim)
}

// TimeFormat parses with given format and returns time.Time type value.
func TimeFormat(sec, key, format string) time.Time {
    if Cfg == nil {
        return time.Now()
    }
    return Cfg.Section(sec).Key(key).MustTimeFormat(format)
}

func TimesFormat(sec, key, format, delim string) []time.Time {
    if Cfg == nil {
        return nil
    }
    return Cfg.Section(sec).Key(key).TimesFormat(format, delim)
}

// Times parses with RFC3339 format and returns list of time.Time devide by given delimiter.
func Times(sec, key, delim string) []time.Time {
    if Cfg == nil {
        return nil
    }
    return Cfg.Section(sec).Key(key).Times(delim)
}

func Init(cfgPath string) {
    var err error
    Cfg, err = ini.Load(cfgPath)
    if err != nil {
        panic(err)
    }

    ENV = Cfg.Section("").Key("ENV").String()
    AppName = Cfg.Section("").Key("APP_NAME").String()
    AppVer = Cfg.Section("").Key("APP_VERSION").String()
    ListeningPort = Cfg.Section("server").Key("HTTP_PORT").String()

    return
}
