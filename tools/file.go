package tools

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"time"
)

func GetFileModTime(fileName string) (time time.Time, err error){
	f, err := os.Open(fileName)
	defer f.Close()
	if err!=nil {
		return
	}

	fileInfo, err := f.Stat()
	if err!=nil{
		return
	}

	time  = fileInfo.ModTime()
	return
}

func GetFileSize(fileName string) (size int64){
	f, err := os.Open(fileName)
	defer f.Close()
	if err!=nil {
		return 0
	}

	fileInfo, err := f.Stat()
	if err!=nil{
		return 0
	}

	size  = fileInfo.Size()
	return
}


func FileExists(fileName string) bool {
	if _, err := os.Stat(fileName); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}



func CreateFileDir(fileDir string) error {
	if err := os.MkdirAll(fileDir, 0777); err != nil {
		return err
	}

	return nil
}


func FileMD5(fileName string) (string, error){
	f, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer f.Close()

	md5New := md5.New()
	if _, err := io.Copy(md5New, f); err != nil {
		return "", err
	}

	return hex.EncodeToString(md5New.Sum(nil)), nil
}