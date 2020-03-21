package dsafe

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"io"
	"mime/multipart"
)

func FileMD5Check(file multipart.File, sign string) (bool, error) {
	if m, err := FileMD5(file); err != nil {
		return false, err
	} else {
		if string(m) != sign {
			return false, nil
		}
		return true, nil
	}
}

func DataMD5Check(data []byte, sign string) (bool, error) {
	if m, err := DataMD5(data); err != nil {
		return false, err
	} else {
		if string(m) != sign {
			return false, nil
		}
		return true, nil
	}
}

func FileMD5(file multipart.File) ([]byte, error) {
	h := md5.New()
	if _, err := io.Copy(h, file); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

func DataMD5(data []byte) ([]byte, error) {
	h := md5.New()
	if _, err := h.Write(data); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

func HmacSha1(k []byte, v []byte) []byte {
	mac := hmac.New(sha1.New, k)
	mac.Write(v)
	return mac.Sum(nil)
}
