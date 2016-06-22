package fs

import (
    "errors"
    "net/http"
    "path/filepath"
)

var (
    mimetype = map[string]string{"image/gif": ".gif", "image/png": ".png", "image/jpeg": ".jpg", "image/pjpeg": ".jpg"}
)

func DownloadImageByUrl(url string) (string, error) {
    fileName := filepath.Base(url)

    response, err := http.Get(url)
    if err != nil {
        l.Error("%v", err)
        return "", err
    }

    mimeType := response.Header.Get("Content-Type")
    if _, ok := mimetype[mimeType]; !ok {
        err = errors.New("Unknown MIME type " + mimeType)
        l.Error("%v", err)
        return "", err
    }

    return SaveFile(fileName, mimeType, response.Body)
}
