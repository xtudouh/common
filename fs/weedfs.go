package fs

import (
    "github.com/ginuerzh/weedo"
    "io"
    "xtudouh/common/conf"
    "xtudouh/common/log"
    "strings"
)

var (
    client *weedo.Client
    l = log.NewLogger()
)

func SaveFile(filename, mimetype string, file io.Reader) (string, error) {
    fid, _, err := client.AssignUpload(filename, mimetype, file)
    if err != nil {
        l.Error("error:%v\n", err)
        return "", err
    }
    return fid, nil
}

func GetUrlByFid(fid string) (string, string, error) {
    publicUrl, url, err := client.GetUrl(fid)
    if err != nil {
        l.Error("error:%v\n", err)
        return "", "", err
    }
    return publicUrl, url, nil
}

func Init() {
    l.Debug("fs module start initializing.")
    masterUrl := conf.String("weedfs", "MASTER_URL", "localhost:9333")
    filerUrls := conf.Strings("weedfs", "FILER_URL", ",")
    for i, u := range filerUrls {
        filerUrls[i] = strings.TrimSpace(u)
    }
    client = weedo.NewClient(masterUrl, filerUrls...)
    l.Debug("fs module initialize successfully.")
}
