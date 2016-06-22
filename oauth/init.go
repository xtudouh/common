package oauth

import "xtudouh/common/log"

var l = log.NewLogger()

type initFun func()

var initFunList []initFun

func registerInitFun(f initFun) {
    initFunList = append(initFunList, f)
}

func Init() {
    for _, f := range initFunList {
        f()
    }
}
