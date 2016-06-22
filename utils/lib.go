package utils

import (
	"bytes"
	"code.google.com/p/graphics-go/graphics"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path"
	"xtudouh/common/log"
	"golang.org/x/crypto/pbkdf2"
	"crypto/sha256"
	"encoding/hex"
)

var l = log.NewLogger()

func Clone(a, b interface{}) {
	buff := new(bytes.Buffer)
	enc := gob.NewEncoder(buff)
	dec := gob.NewDecoder(buff)
	enc.Encode(a)
	dec.Decode(b)
}

func Serialize(v interface{}) []byte {
	buff := new(bytes.Buffer)
	enc := gob.NewEncoder(buff)
	enc.Encode(v)
	return buff.Bytes()
}

func UnSerialize(data []byte, inst interface{}) {
	buf := new(bytes.Buffer)
	buf.Write(data)
	dec := gob.NewDecoder(buf)
	dec.Decode(inst)
	return
}

const (
	UPPER_CHARSET = iota
	LOWER_CHARSET
	INT_CHARSET
	UPPER_LOWER_CHARSET
	FULL_CHARSET
)

var charset = []byte{
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N',
	'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n',
	'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
}

func RandomSpecStr(num int, set uint) string {
	var subcharset []byte

	switch set {
	case UPPER_CHARSET:
		subcharset = charset[:25]
	case LOWER_CHARSET:
		subcharset = charset[26:51]
	case INT_CHARSET:
		subcharset = charset[52:]
	case UPPER_LOWER_CHARSET:
		subcharset = charset[:51]
	case FULL_CHARSET:
		subcharset = charset
	default:
		if l != nil {
			l.Error("%s", "Unkonwn charset identify.")
		}
		return ""
	}

	var buf = make([]byte, num)
	for i := 0; i < num; i++ {
		index := rand.Intn(len(subcharset))
		buf[i] = subcharset[index]
	}
	return string(buf)
}

func RandomStr(num int) string {
	return RandomSpecStr(num, FULL_CHARSET)
}

func EncryptPassword(password string) string {
	pk := pbkdf2.Key([]byte(password), []byte("salt"), 4096, 64, sha256.New)
	return hex.EncodeToString(pk)
}

func In(e string, list []string) bool {
	for _, v := range list {
		if e == v {
			return true
		}
	}

	return false
}

func SaveFile(src io.Reader, fileName string) (err error) {
	dst, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0766)
	if err != nil {
		return
	}

	defer dst.Close()

	_, err = io.Copy(dst, src)

	if err != nil {
		if l != nil {
			l.Error("%v", err)
		}
		return
	}

	return
}

func ImageSize(filePath string) (width, height int, err error) {
	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		return 0, 0, err
	}

	img, _, err := image.Decode(f)
	if err != nil {
		if l != nil {
			l.Error("%v", err)
		}
		return 0, 0, err

	}
	p := img.Bounds().Size()

	return p.X, p.Y, nil
}

func ScaleImage(filepath string, w, h int) (string, error) {
	src, err := loadImage(filepath)
	if err != nil {
		if l != nil {
			l.Error("%v", err)
		}
		return "", err
	}

	dst := image.NewRGBA(image.Rect(0, 0, w, h))
	if err := graphics.Scale(dst, src); err != nil {
		if l != nil {
			l.Error("%v", err)
		}
		return "", err
	}
	dstpath := path.Join(path.Dir(filepath), fmt.Sprintf("scale%dx%d.%s", w, h, path.Base(filepath)))
	if err := saveImage(dstpath, dst); err != nil {
		if l != nil {
			l.Error("%v", err)
		}
		return "", err
	}
	return dstpath, nil
}

func loadImage(filepath string) (image.Image, error) {
	f, err := os.Open(filepath)
	defer f.Close()
	if err != nil {
		if l != nil {
			l.Error("%v", err)
		}
		return nil, err
	}
	img, _, err := image.Decode(f)
	if err != nil {
		if l != nil {
			l.Error("%v", err)
		}
		return nil, err
	}

	return img, nil
}

func saveImage(filepath string, img image.Image) (err error) {
	imgfile, err := os.Create(filepath)
	defer imgfile.Close()
	switch path.Ext(filepath) {
	case ".jpg", ".jpeg":
		if err = jpeg.Encode(imgfile, img, nil); err != nil {
			return
		}
	case ".png":
		if err = png.Encode(imgfile, img); err != nil {
			return
		}
	case ".gif":
		if err = gif.Encode(imgfile, img, new(gif.Options)); err != nil {
			return
		}
	}
	return
}

func FetchUrlWithJson(urlStr string, result interface{}) error {
	resp, err := http.Get(urlStr)
	if err != nil {
		l.Error("%v", err)
		return err
	}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(result); err != nil {
		l.Error("%v", err)
		return err
	}

	return nil
}
