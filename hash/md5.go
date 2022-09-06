package hash

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
)

func FileMd5(filename string) (string, error) {
	file, err := os.OpenFile(filename, os.O_RDWR, 0666)

	if err != nil {
		return "", errors.New(fmt.Sprintf("md5.go hash.FileMd5 os open error %v", err))
	}

	h := md5.New()
	_, err = io.Copy(h, file)
	if err != nil {
		return "", errors.New(fmt.Sprintf("md5.go hash.FileMd5 io copy error %v", err))
	}

	return hex.EncodeToString(h.Sum(nil)), nil

}

func StringMd5(str string) string {
	md := md5.New()
	md.Write([]byte(str))
	return hex.EncodeToString(md.Sum(nil))
}
