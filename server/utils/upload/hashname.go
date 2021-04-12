package upload

import (
	"github.com/satori/go.uuid"
	"path"
	"path/filepath"
	"time"
)

func GetUUID()  string{
	u := uuid.NewV4()
	return u.String()
}

func NewFilePath(filename string)  (newFilePath string){
	fid := GetUUID()
	fileSuffix := path.Ext(filename)
	newFilePath = filepath.Join("uploads", time.Now().Format("2006-01-02"))  + "/" + fid + fileSuffix
	return
}